// Trimmer SDK
//
// Copyright (c) 2017 Alexander Eichhorn
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

// Upload a media file into a new asset using multipart/chunked upload
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	. "trimmer.io/go-trimmer"
	"trimmer.io/go-trimmer/asset"
	"trimmer.io/go-trimmer/media"
	"trimmer.io/go-trimmer/meta"
	"trimmer.io/go-trimmer/rfc"
	"trimmer.io/go-trimmer/session"
	"trimmer.io/go-trimmer/workspace"
)

var (
	Timecode = flag.String("tc", "00:00:00:00", "media start timecode")
	Access   = flag.String("access", "", "access class (public, private, personal)")
	Role     = flag.String("role", "", "media role (e.g. video)")
	Reel     = flag.String("reel", "", "media reel name")
	Debug    = flag.Bool("debug", false, "enable debugging")
	Family   = flag.String("family", "capture", "default media family")
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
	log.Println("Uploading file", flag.Arg(1), "into workspace", flag.Arg(0))

	wid := flag.Arg(0)

	family := media.ParseMediaFamily(*Family)
	if family == media.MediaFamilyUndefined {
		fail("invalid media family")
		usage()
	}

	// open the file to upload
	var f *os.File
	var err error
	if f, err = os.Open(flag.Arg(1)); err != nil {
		log.Fatalln("File open failed:", err)
	}
	defer f.Close()

	// get file type and size
	var d os.FileInfo
	if d, err = f.Stat(); err != nil {
		log.Fatalln("File stat failed:", err)
	}

	if !d.Mode().IsRegular() {
		log.Fatalln("File is not a regular file")
	}

	filename := d.Name()

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
		Title: strings.TrimSuffix(filename, filepath.Ext(filename)),
		Embed: API_EMBED_META,
	}

	if *Access != "" {
		ap.Access = AccessClass(*Access)
	}

	if *Reel != "" {
		ap.Actions = append(ap.Actions, MetaValue{
			Path:  meta.ShotReelName,
			Value: *Reel,
		})
	}

	a, err := workspace.NewAsset(ctx, wid, ap)
	if err != nil {
		log.Fatalln("Cannot create asset:", err)
	}

	// a new media
	mp := &MediaParams{
		Filename:   filename,
		Size:       d.Size(),
		Mimetype:   rfc.MimeTypeFromName(filename),
		RecordedAt: d.ModTime(),
		Timecode:   *Timecode,
		Type:       media.MediaTypeAudioVideo,
		Family:     family,
		Format:     media.ParseMediaFormat(filepath.Ext(filename)),
	}

	if r := media.ParseMediaRole(*Role); media.IsValidRole(r) {
		mp.Role = r
	}

	// upload media (implicitly creates it)
	m, err := asset.UploadMedia(ctx, a.ID, mp, f)
	if err != nil {
		log.Fatalln("Cannot upload media into asset.")
	}

	log.Println("Created new media", m.ID, "and uploaded", m.Size, "bytes in", time.Since(start))
}
