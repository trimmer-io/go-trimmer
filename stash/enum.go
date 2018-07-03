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

package stash

import (
	trimmer "trimmer.io/go-trimmer"
)

const (
	StashTypeUndefined trimmer.StashType = ""
	StashTypeSystem    trimmer.StashType = "system"  // cannot delete, can have links
	StashTypeVirtual   trimmer.StashType = "virtual" // cannot delete, cannot link/unlink
	StashTypeUser      trimmer.StashType = "user"    // can delete, can have links
	StashTypeSmart     trimmer.StashType = "smart"   // can delete, cannot have links
)

func ParseStashType(s string) trimmer.StashType {
	switch s {
	case "system":
		return StashTypeSystem
	case "virtual":
		return StashTypeVirtual
	case "user":
		return StashTypeUser
	case "smart":
		return StashTypeSmart
	default:
		return StashTypeUndefined
	}
}
