// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

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
