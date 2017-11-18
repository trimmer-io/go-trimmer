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
