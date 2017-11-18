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

// Downloads media files from asset
package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/cheggaaa/pb"

	. "trimmer.io/go-trimmer"
	"trimmer.io/go-trimmer/asset"
	"trimmer.io/go-trimmer/media"
	"trimmer.io/go-trimmer/session"
)

var (
	Debug = flag.Bool("debug", false, "enable debugging")
)

func Download(ctx context.Context, aid string, m *Media) error {

	if m == nil {
		return nil
	}

	// build file name:
	path := filepath.Join(aid, string(m.Relation), m.Filename)

	// make sure directory exists
	var (
		w   *os.File
		err error
	)
	if err = os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}

	// try opening the output file at target key
	w, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	// cleanup (close on success, remove file on error)
	defer func() {
		n := w.Name()
		w.Close()
		if err != nil {
			os.Remove(n)
		}
	}()

	// progress bar
	bar := pb.New64(m.Size).SetUnits(pb.U_BYTES).SetRefreshRate(time.Second)
	bar.ShowBar = m.Size != 0
	bar.ShowTimeLeft = m.Size != 0
	bar.ShowSpeed = true
	bar.SetUnits(pb.U_BYTES)
	bar.Start()
	defer bar.Finish()

	fi, err := media.Download(ctx, m, io.MultiWriter(w, bar))
	if err != nil {
		return err
	}
	total += fi.Size
	return nil
}

var total int64

func main() {

	ctx := context.Background()

	start := time.Now()
	LogLevel = 2
	flag.Parse()
	if *Debug {
		LogLevel = 3
	}

	if len(flag.Args()) == 0 {
		log.Fatalln("Usage:", os.Args[0], "<assetId> [<role>]")
	}
	log.Println("Downloading asset", flag.Arg(0), "role", flag.Arg(1))

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

	// list media by role
	role := media.MediaRoleUndefined
	if len(flag.Args()) > 1 {
		role = media.ParseMediaRole(flag.Arg(1))
		if role == media.MediaRoleUndefined {
			log.Fatalln("Invalid media role", flag.Arg(1))
		}
	}

	// list media in asset
	aid := flag.Arg(0)
	mIter := asset.ListMedia(ctx, aid, &MediaListParams{
		Roles: MediaRoleList{role},
		Embed: API_EMBED_URLS,
	})
	for mIter.Next() {
		m := mIter.Media()
		if m.Url != "" {
			if err := Download(ctx, aid, m); err != nil {
				log.Println("Download failed:", err)
			}
		} else {
			log.Println("Empty media url!")
		}
	}

	log.Println("Downloaded", total, "bytes, Total Runtime", time.Since(start))
	log.Println("OK")

}
