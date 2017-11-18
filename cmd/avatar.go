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

// Uploads a new Avatar image for the user
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	. "trimmer.io/go-trimmer"
	"trimmer.io/go-trimmer/media"
	"trimmer.io/go-trimmer/rfc"
	"trimmer.io/go-trimmer/session"
	"trimmer.io/go-trimmer/user"
)

var (
	Debug = flag.Bool("debug", false, "enable debugging")
)

func main() {

	ctx := context.Background()

	flag.Parse()
	if *Debug {
		LogLevel = 3
	}

	if len(os.Args) < 2 {
		log.Fatalln("Usage:", os.Args[0], "<filename>")
	}

	// open the file to upload
	var f *os.File
	var err error
	if f, err = os.Open(os.Args[len(os.Args)-1]); err != nil {
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

	// upload media (implicitly creates it)
	m, err := user.UploadImage(ctx, &FileInfo{
		Filename: filename,
		Mimetype: rfc.MimeTypeFromName(filename),
		Role:     media.MediaRoleAvatar,
	}, f)
	if err != nil {
		log.Fatalln("Cannot upload media.")
	}

	// wait for media to become ready
	log.Printf("Waiting for media to become ready...")
	for {
		time.Sleep(time.Second)
		m, err = media.Get(ctx, m.ID, nil)
		if err != nil {
			log.Fatalln("waiting failed!")
		}
		if m.State == media.MediaStateReady {
			break
		}
		log.Print(".")
	}

	// set the new media as the user's avatar image
	var u *User
	u, err = user.UpdateMe(ctx, &UserParams{
		ImageId: m.ID,
		Embed:   API_EMBED_MEDIA,
	})
	if err != nil {
		log.Fatalln("Cannot udpate user profile.")
	}

	// check new media ID is used
	if u.ImageId != m.ID {
		log.Fatalln("Mismatched user profile image: expected", m.ID, "got", u.ImageId)
	}
	log.Println("OK")

}
