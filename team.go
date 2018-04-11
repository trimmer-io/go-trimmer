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

// TeamListParams is the set of parameters that can be used when listing teams.
type TeamListParams struct {
	ListParams
	Embed ApiEmbedFlags `json:"embed,omitempty"`
}

// TeamParams is the set of parameters that can be used when reading, creating or
// updating teams.
type TeamParams struct {
	DisplayName string        `json:"displayName"`
	ImageId     string        `json:"imageId,omitempty"`
	Homepage    string        `json:"homepage,omitempty"`
	Location    string        `json:"location,omitempty"`
	Description string        `json:"description"`
	Role        MemberRole    `json:"role,omitempty"`
	Permissions Permissions   `json:"perm,omitempty"`
	Clearance   AccessClass   `json:"clearance,omitempty"`
	Embed       ApiEmbedFlags `json:"embed,omitempty"`
}

// Team is the resource representing a Trimmer team.
type Team struct {
	ID          string       `json:"teamId"`
	OrgID       string       `json:"orgId"`
	DisplayName string       `json:"displayName"`
	Location    string       `json:"location"`
	Description string       `json:"description"`
	Homepage    string       `json:"homepage"`
	ImageKey    string       `json:"imageKey"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	Media       *MediaEmbed  `json:"media"`
	Membership  *Member      `json:"membership"`
	Permissions *Permissions `json:"perms"`
	Org         *Org         `json:"org"`
}

type TeamList []*Team

func (l TeamList) SearchId(id string) (int, *Team) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}
