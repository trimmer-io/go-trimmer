// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

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
