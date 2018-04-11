// Trimmer SDK
//
// Copyright (c) 2017-2018 Alexander Eichhorn
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Upload a sequence of media files into a new asset using multipart/chunked upload
//
// Example:
//  ./sequence --debug --family=capture.arri 26887CcfeQK <dir>
//
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	. "trimmer.io/go-trimmer"
	"trimmer.io/go-trimmer/asset"
	"trimmer.io/go-trimmer/media"
	// "trimmer.io/go-trimmer/meta"
	"trimmer.io/go-trimmer/rfc"
	"trimmer.io/go-trimmer/session"
	"trimmer.io/go-trimmer/workspace"
)

var (
	Debug  = flag.Bool("debug", false, "enable debugging")
	Family = flag.String("family", "capture", "default media family")
)

func fail(v interface{}) {
	fmt.Printf("Error: %v\n", v)
	os.Exit(1)
}

func usage() {
	log.Println("Usage:", os.Args[0], "[options] <workspaceId> <filename>")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {

	start := time.Now()
	ctx := context.Background()

	flag.Parse()
	if *Debug {
		LogLevel = 3
	}

	if len(os.Args) < 3 {
		usage()
	}
	log.Println("Uploading directory", flag.Arg(1), "into workspace", flag.Arg(0))

	wid := flag.Arg(0)
	directory, err := filepath.Abs(flag.Arg(1))
	if err != nil {
		fail(fmt.Errorf("no absolute path: %v", err))
	}

	family := media.ParseMediaFamily(*Family)
	if family == media.MediaFamilyUndefined {
		fail("invalid media family")
		usage()
	}

	// open the directory to upload
	f, err := os.Open(directory)
	if err != nil {
		log.Fatalln("Directory open failed:", err)
	}
	d, err := f.Stat()
	if err != nil {
		log.Fatalln("Stat failed:", err)
	}
	if !d.IsDir() {
		fail("Not a directory")
	}
	dirname := filepath.Base(f.Name())
	f.Close()

	// prepare list of files to upload
	seq := make(SequenceFileList, 0)
	var (
		format       MediaFormat
		mime         string
		commonPrefix string
		assetName    string
		size         int64
	)
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		// ignore invalid files and files that are not sequence media
		if !info.Mode().IsRegular() {
			return nil
		}
		f := media.ParseMediaFormat(filepath.Ext(path))
		if f == media.MediaFormatUndefined {
			return nil
		}
		if !media.IsSequenceMediaFormat(f) {
			return nil
		}

		// find a common path prefix relative to the walked directory
		relPath, err := filepath.Rel(directory, path)
		if err != nil {
			return nil
		}
		baseDir := filepath.Dir(relPath)

		// first iteration init
		if format == media.MediaFormatUndefined {
			format = f
			mime = rfc.MimeTypeFromName(path)
			commonPrefix = baseDir
		}
		commonPrefix = CommonPrefix(os.PathSeparator, commonPrefix, baseDir)

		// try reading name and frame number from filename xxx.<num>.format
		parts := strings.Split(filepath.Base(relPath), ".")
		if len(parts) < 2 {
			return nil
		}
		if assetName == "" {
			assetName = parts[0]
		}
		num, err := strconv.ParseInt(parts[len(parts)-2], 10, 64)
		if err != nil {
			return nil
		}
		size += info.Size()
		seq = append(seq, &SequenceFile{
			Frame:    num,
			Size:     info.Size(),
			Filename: relPath,
		})
		return nil
	})

	// remove the common file prefix from file names and add it to dirname
	if commonPrefix != "" {
		dirname = filepath.Join(dirname, commonPrefix)
		log.Println("Final Dir", dirname)
		for _, v := range seq {
			v.Filename, _ = filepath.Rel(commonPrefix, v.Filename)
			log.Println("Final Filename", v.Filename)
		}
	}

	// try fetching client token or user credentials from env
	if _, err := NewClientSession(""); err != nil {
		if err := session.Login(ctx, session.ParseEnv()); err != nil {
			log.Fatalln("Login failed.")
		}
		defer session.Logout(ctx)
	}
	if err := session.Check(ctx); err != nil {
		log.Fatalln("Check failed.")
	}

	// create new asset
	ap := &AssetParams{
		Title:  assetName,
		Access: ACCESS_PUBLIC,
		Embed:  API_EMBED_META,
	}

	a, err := workspace.NewAsset(ctx, wid, ap)
	if err != nil {
		log.Fatalln("Cannot create asset.")
	}

	// a new media
	mp := &MediaParams{
		Filename: dirname,
		Size:     size,
		Mimetype: mime,
		Role:     media.MediaRoleVideo,
		Type:     media.MediaTypeSequence,
		Family:   family,
		Format:   format,
		Attr:     NewMediaAttr(),
		Embed:    API_EMBED_META | API_EMBED_DETAILS,
	}

	mp.Attr.Sequence = append(mp.Attr.Sequence, &SequenceAttr{
		FrameCount: int64(len(seq)),
		MediaList:  seq,
	})

	// create media for upload
	m, err := asset.NewUpload(ctx, a.ID, mp)
	if err != nil {
		log.Fatalln("Cannot create asset media:", err)
	}

	// upload files into media
	var files FileInfoList
	files, err = media.UploadMulti(ctx, m, func(fi *FileInfo) (io.ReadSeeker, error) {

		// close the previous file, update metrics
		if f != nil {
			f.Close()
		}

		path := filepath.Join(directory, commonPrefix, fi.Filename)
		log.Println("UPLOAD", path, "->", fi.Url)

		// open file for upload
		f, err = os.Open(path)
		if err != nil {
			return nil, err
		}

		// get file type and size
		s, err := f.Stat()
		if err != nil {
			return nil, err
		}

		// be resilient to concurrent writes to files
		sz := s.Size()
		if fi.Size == 0 {
			fi.Size = sz
		} else if sz < fi.Size {
			return nil, fmt.Errorf("file size mismatch: expected %d, got %d", fi.Size, sz)
		}

		return f, nil
	})

	// close the last file opened by the callback, update metrics
	if f != nil {
		f.Close()
	}

	// if anything went wrong during multi-file upload, this will be the first error
	if err != nil {
		fail(err)
	}

	// complete the file upload
	cp := &MediaUploadCompletionParams{
		Files: files,
	}
	if _, err := media.CompleteUpload(ctx, m.ID, cp); err != nil {
		fail(err)
	}

	log.Println("Created new asset ", a.ID, ", media", m.ID, "and uploaded", m.Size, "bytes in", time.Since(start))
}

func CommonPrefix(sep byte, paths ...string) string {
	switch len(paths) {
	case 0:
		return ""
	case 1:
		return path.Clean(paths[0])
	}

	c := []byte(path.Clean(paths[0]))
	c = append(c, sep)
	for _, v := range paths[1:] {
		v = path.Clean(v) + string(sep)
		if len(v) < len(c) {
			c = c[:len(v)]
		}
		for i := 0; i < len(c); i++ {
			if v[i] != c[i] {
				c = c[:i]
				break
			}
		}
	}

	for i := len(c) - 1; i >= 0; i-- {
		if c[i] == sep {
			c = c[:i]
			break
		}
	}

	return string(c)
}
