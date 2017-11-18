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

package rfc

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	MIME_TYPE_DEFAULT = "application/octet-stream"
)

var (
	EMimeEmptyContent        = errors.New("cannot detect mime type: empty content")
	EMimeUnidentifiedContent = errors.New("cannot detect mime type: unidentified content")
	EMimeUnknownExtension    = errors.New("cannot detect mime type: unknown extension")
)

var (
	TRIMMER_EXTRA_MIMETYPES map[string]string = map[string]string{
		".3fr":  "image/x-hasselblad-3fr",
		".ale":  "application/vnd.avid.ale",
		".aml":  "application/vnd.arri.look",
		".ari":  "application/vnd.arri.arriraw",
		".arw":  "image/x-sony-arw",
		".cdl":  "application/cdl+xml",
		".cine": "video/x-cine",
		".cr2":  "image/x-canon-cr2",
		".csv":  "text/csv",
		".cube": "text/plain",
		".data": "text/plain",
		".dng":  "image/x-adobe-dng",
		".dpx":  "image/x-dpx",
		".edl":  "text/plain",
		".exr":  "image/x-exr",
		".flac": "audio/flac",
		".jp2":  "image/jp2",
		".jpx":  "image/jpx",
		".look": "application/look+xml",
		".lut":  "text/plain",
		".m2ts": "video/mp2t",
		".m3u":  "application/x-mpegurl",
		".m3u8": "application/x-mpegurl",
		".mkv":  "video/x-matroska",
		".mov":  "video/quicktime",
		".mpd":  "application/dash+xml",
		".mxf":  "application/mxf",
		".psd":  "image/vnd.adobe.photoshop",
		".r3d":  "application/vnd.red.redcode",
		".rmd":  "application/vnd.red.redmetadata",
		".srt":  "application/x-subrip",
		".svg":  "image/svg+xml",
		".vtt":  "text/vtt",
		".xmp":  "application/rdf+xml",
	}
)

func init() {
	for ext, typ := range TRIMMER_EXTRA_MIMETYPES {
		if err := mime.AddExtensionType(ext, typ); err != nil {
			log.Fatalln("mime type init failed:", err)
		}
	}
}

func MimeTypeFromPath(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	d, err := f.Stat()
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("'%s' not found", path)
		} else {
			return "", fmt.Errorf("cannot read '%s': %v", path, err)
		}
	}
	if !d.Mode().IsRegular() {
		return "", fmt.Errorf("'%s' is not a file", path)
	}
	if d.Size() == 0 {
		return "", fmt.Errorf("'%s' is empty", path)
	}
	return MimeTypeFromFile(f)
}

func MimeTypeFromFile(f *os.File) (string, error) {
	b := make([]byte, 512)
	f.Seek(0, io.SeekStart)
	n, _ := f.Read(b)
	f.Seek(0, io.SeekStart)
	return MimeTypeFromByte(b[:n])
}

func MimeTypeFromByte(b []byte) (string, error) {
	if len(b) == 0 {
		return "", EMimeEmptyContent
	}
	if m := http.DetectContentType(b); m != MIME_TYPE_DEFAULT {
		return m, nil
	}
	return MIME_TYPE_DEFAULT, EMimeUnidentifiedContent
}

func MimeTypeFromName(name string) string {
	ext := filepath.Ext(name)
	if mm := mime.TypeByExtension(ext); mm != "" {
		return mm
	}
	return MIME_TYPE_DEFAULT
}

// split anything after a ; or + first, then split again at /
func SplitMimeType(m string) (major, minor, codec string) {
	f := strings.Split(m, ";")
	if len(f) > 1 {
		codec = strings.ToLower(f[1])
	}
	f = strings.Split(f[0], "+")
	if len(f) > 1 {
		codec = strings.ToLower(f[1])
	}
	f = strings.Split(f[0], "/")
	if len(f) > 1 {
		minor = strings.ToLower(f[1])
	}
	major = strings.ToLower(f[0])
	return
}
