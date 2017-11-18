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

package member

import (
	trimmer "trimmer.io/go-trimmer"
)

const (
	MemberRoleUndefined trimmer.MemberRole = ""
	MemberRoleOwner     trimmer.MemberRole = "owner"
	MemberRoleAdmin     trimmer.MemberRole = "admin"
	MemberRoleMember    trimmer.MemberRole = "member"
	MemberRoleGuest     trimmer.MemberRole = "guest"
)

func ParseMemberRole(s string) trimmer.MemberRole {
	switch s {
	case "owner":
		return MemberRoleOwner
	case "admin":
		return MemberRoleAdmin
	case "member":
		return MemberRoleMember
	case "guest":
		return MemberRoleGuest
	default:
		return MemberRoleUndefined
	}
}

const (
	MemberTypeUndefined    trimmer.MemberType = ""
	MemberTypeOrganization trimmer.MemberType = "org"
	MemberTypeTeam         trimmer.MemberType = "team"
	MemberTypeWorkspace    trimmer.MemberType = "work"
)

func ParseMemberType(s string) trimmer.MemberType {
	switch s {
	case "org":
		return MemberTypeOrganization
	case "team":
		return MemberTypeTeam
	case "work":
		return MemberTypeWorkspace
	default:
		return MemberTypeUndefined
	}
}
