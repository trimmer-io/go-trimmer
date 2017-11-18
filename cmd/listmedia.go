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

// Lists all assets and contained media across a user's workspaces
package main

import (
	"context"
	"flag"
	"log"

	. "trimmer.io/go-trimmer"
	"trimmer.io/go-trimmer/asset"
	"trimmer.io/go-trimmer/meta"
	"trimmer.io/go-trimmer/session"
	"trimmer.io/go-trimmer/user"
	"trimmer.io/go-trimmer/workspace"
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

	w := user.ListWorkspaces(ctx, nil)
	if err := w.Err(); err != nil {
		log.Printf("%v\n", err)
	}

	// auto-page through all workspaces
	for w.Next() {
		v := w.Workspace()
		if v == nil {
			log.Fatalln("No nil values expected")
		}

		// print workspace entry to console
		log.Printf("Workspace %-11s %-11s\n",
			v.ID,
			v.Name,
		)

		a := workspace.ListAssets(ctx, v.ID, &AssetListParams{
			Embed: API_EMBED_META | API_EMBED_STATS,
		})
		if err := a.Err(); err != nil {
			log.Printf("%v\n", err)
		}
		for a.Next() {
			v := a.Asset()
			if v == nil {
				log.Fatalln("No nil values expected")
			}

			// print workspace entry to console
			t, err := v.Metadata.GetPath(meta.AssetTitle)
			if err != nil {
				log.Println("ERROR", err)
				return
			}
			log.Printf("Asset %-11s %-25s %-20s %+v\n",
				v.ID,
				"",
				t,
				v.Statistics,
			)

			m := asset.ListMedia(ctx, v.ID, &MediaListParams{
				Embed: API_EMBED_URLS,
			})
			if err := m.Err(); err != nil {
				log.Printf("%v\n", err)
			}
			for m.Next() {
				v := m.Media()
				if v == nil {
					log.Fatalln("No nil values expected")
				}

				// print media entry to console
				log.Printf("%-11s %-11s %-11s %-5s %-20s %-6s %s\n",
					v.ID,
					v.State,
					v.Type,
					v.Format,
					v.Role,
					v.Relation,
					v.Url,
				)
			}
		}
	}
}
