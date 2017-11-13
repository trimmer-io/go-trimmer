// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package trimmer

import (
	"time"
)

// MemberRole is the list of allowed values for a membership role.
// Allowed values are "owner", "admin", "member", "guest".
type MemberRole string

// MemberType is the list of allowed values for a membership type.
// Allowed values are "org", "team", "work".
type MemberType string

// MemberParams is the set of parameters that can be used to create and
// update memberships.
//
type MemberParams struct {
	Role       MemberRole    `json:"role,omitempty"`
	Permission Permission    `json:"perm,omitempty"`
	State      UserState     `json:"state,omitempty"`
	Clearance  AccessClass   `json:"clearance,omitempty"`
	Embed      ApiEmbedFlags `json:"embed,omitempty"`
}

// MemberListParams is the set of parameters that can be used when listing memberships.
type MemberListParams struct {
	ListParams
	State UserState     `json:"state,omitempty"`
	Role  MemberRole    `json:"role,omitempty"`
	Embed ApiEmbedFlags `json:"embed,omitempty"`
}

// Member is the resource representing a Trimmer user who is member of a workspace,
// organization or team.
type Member struct {
	ID          string      `json:"memberId"`
	AccountID   string      `json:"accountId"`
	AuthorId    string      `json:"authorId"`
	Type        MemberType  `json:"type"`
	State       UserState   `json:"state"`
	Role        MemberRole  `json:"role"`
	Clearance   AccessClass `json:"clearance"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	JoinedAt    time.Time   `json:"joinedAt"`
	WorkspaceId string      `json:"workspaceId"`
	TeamId      string      `json:"teamId"`
	OrgId       string      `json:"orgId"`
	UserId      string      `json:"userId"`
	Author      *User       `json:"author"`
	Org         *Org        `json:"organization"`
	Team        *Team       `json:"team"`
	Workspace   *Workspace  `json:"workspace"`
	User        *User       `json:"user"`
}

type MemberList []*Member

func (l MemberList) SearchId(id string) (int, *Member) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}
