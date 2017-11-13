// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

// Package acl provides the /acl APIs
package acl

import (
	trimmer "trimmer.io/go-trimmer"
)

const (
	AclObjectTypeUndefined    trimmer.AclObjectType = ""
	AclObjectTypeApplication  trimmer.AclObjectType = "application"
	AclObjectTypeOrganization trimmer.AclObjectType = "organization"
	AclObjectTypeWorkspace    trimmer.AclObjectType = "workspace"
	AclObjectTypeVolume       trimmer.AclObjectType = "volume"
	AclObjectTypeStash        trimmer.AclObjectType = "stash"
	AclObjectTypeAsset        trimmer.AclObjectType = "asset"
	AclObjectTypeTeam         trimmer.AclObjectType = "team"
	AclObjectTypeUser         trimmer.AclObjectType = "user"
)

func ParseObjectType(s string) trimmer.AclObjectType {
	switch s {
	case "application":
		return AclObjectTypeApplication
	case "organization":
		return AclObjectTypeOrganization
	case "workspace":
		return AclObjectTypeWorkspace
	case "volume":
		return AclObjectTypeVolume
	case "stash":
		return AclObjectTypeStash
	case "asset":
		return AclObjectTypeAsset
	case "team":
		return AclObjectTypeTeam
	case "user":
		return AclObjectTypeUser
	default:
		return AclObjectTypeUndefined
	}
}

const (
	AclSubjectTypeUndefined trimmer.AclSubjectType = ""
	AclSubjectTypeAnonymous trimmer.AclSubjectType = "anon"
	AclSubjectTypeTeam      trimmer.AclSubjectType = "team"
	AclSubjectTypeUser      trimmer.AclSubjectType = "user"
)

func ParseAclSubjectType(s string) trimmer.AclSubjectType {
	switch s {
	case "anon":
		return AclSubjectTypeAnonymous
	case "user":
		return AclSubjectTypeUser
	case "team":
		return AclSubjectTypeTeam
	default:
		return AclSubjectTypeUndefined
	}
}
