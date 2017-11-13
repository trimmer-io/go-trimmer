// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package stash

import (
	trimmer "trimmer.io/go-trimmer"
)

const (
	StashTypeUndefined trimmer.StashType = ""
	StashTypeSystem    trimmer.StashType = "system"  // cannot be deleted
	StashTypeVirtual   trimmer.StashType = "virtual" // cannot delete, cannot link/unlink
	StashTypeUser      trimmer.StashType = "user"    // owned by user
)

func ParseStashType(s string) trimmer.StashType {
	switch s {
	case "system":
		return StashTypeSystem
	case "virtual":
		return StashTypeVirtual
	case "user":
		return StashTypeUser
	default:
		return StashTypeUndefined
	}
}
