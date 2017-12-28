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

// Outputs hash values for a given file
package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"trimmer.io/go-trimmer/hash"
)

var (
	sha1   bool
	sha256 bool
	md5    bool
	xxh    bool
	t2     bool
	all    bool
	none   bool
)

func init() {
	flag.BoolVar(&md5, "md5", false, "enable MD5 hash")
	flag.BoolVar(&sha1, "sha1", false, "enable SHA1 hash")
	flag.BoolVar(&sha256, "sha256", false, "enable SHA256 hash")
	flag.BoolVar(&xxh, "xxh", false, "enable XXhash")
	flag.BoolVar(&t2, "t2", false, "enable Tiger2 hash")
	flag.BoolVar(&all, "all", false, "enable all hashes")
	flag.BoolVar(&none, "none", false, "test file read performance only")
}

func fail(v interface{}) {
	fmt.Printf("Error: %s\n", v)
	os.Exit(1)
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Error: missing filename")
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		fail(err)
	}
	defer f.Close()

	if none {
		io.Copy(ioutil.Discard, f)
		return
	}

	ht := make(hash.HashTypeList, 0)
	if sha1 || all {
		ht.Add(hash.HashTypeSha1)
	}
	if sha256 || all {
		ht.Add(hash.HashTypeSha256)
	}
	if md5 || all {
		ht.Add(hash.HashTypeMd5)
	}
	if xxh || all {
		ht.Add(hash.HashTypeXxhash)
	}
	if t2 || all {
		ht.Add(hash.HashTypeTiger)
	}

	var block hash.HashBlock
	if _, err := io.Copy(ioutil.Discard, block.NewReader(f, ht.Flags())); err != nil {
		fail(err)
	}
	block.Sum()
	fmt.Println(block.JsonString())
}
