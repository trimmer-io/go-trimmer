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
