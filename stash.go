// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package trimmer

import (
	"time"
)

// StashKind is the list of allowed values for stash kinds.
// Allowed values are "system", "virtual", "user".
type StashType string

// StashParams is the set of parameters that can be used to create and
// update a invitations.
//
type StashParams struct {
	Name        string        `json:"name"`
	DisplayName string        `json:"displayName,omitempty"`
	AccessClass AccessClass   `json:"access,omitempty"`
	ImageId     string        `json:"imageId,omitempty"`
	Embed       ApiEmbedFlags `json:"embed,omitempty"`
}

// StashListParams is the set of parameters that can be used when listing stashes.
type StashListParams struct {
	ListParams
	Name        string        `json:"name,omitempty"`
	AccessClass AccessClass   `json:"access,omitempty"`
	Type        StashType     `json:"type,omitempty"`
	Embed       ApiEmbedFlags `json:"embed,omitempty"`
}

// Stash is the resource representing a Trimmer stash.
type Stash struct {
	ID          string           `json:"stashId"`
	WorkspaceId string           `json:"workspaceId"`
	AuthorId    string           `json:"authorId"`
	Name        string           `json:"name"`
	DisplayName string           `json:"displayName"`
	Type        StashType        `json:"type"`
	AccessClass AccessClass      `json:"access"`
	ImageId     string           `json:"imageId"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
	Watching    bool             `json:"watching"`
	Permissions *Permissions     `json:"perms"`
	Workspace   *Workspace       `json:"workspace"`
	Author      *User            `json:"author"`
	Media       *MediaEmbed      `json:"media"`
	Statistics  *StashStatistics `json:"stats"`
}

type StashList []*Stash

func (l StashList) SearchId(id string) (int, *Stash) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}

// StashStatistics is the sub-resource representing stash counters.
type StashStatistics struct {
	Links    int64 `json:"clips"`
	Assets   int64 `json:"assets"`
	Time     int64 `json:"time"`
	Watchers int64 `json:"watchers"`
	Media    int64 `json:"media"`
	Size     int64 `json:"size"`
	Files    int64 `json:"files"`
}
