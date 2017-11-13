// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package trimmer

import (
	"time"
)

// LinkParams is the set of parameters that can be used to create and
// update a invitations.
//
type LinkParams struct {
	AssetId  string        `json:"assetId"`
	MediaIn  time.Duration `json:"mediaIn,omitempty"`
	MediaOut time.Duration `json:"mediaOut,omitempty"`
	Embed    ApiEmbedFlags `json:"embed,omitempty"`
}

// LinkListParams is the set of parameters that can be used when listing stashes.
type LinkListParams struct {
	ListParams
	AssetId  string        `json:"assetId,omitempty"`
	StashId  string        `json:"stashId,omitempty"`
	AuthorId string        `json:"authorId,omitempty"`
	Embed    ApiEmbedFlags `json:"embed,omitempty"`
}

// Link is the resource representing a Trimmer stash.
type Link struct {
	ID        string        `json:"linkId"`
	OwnerId   string        `json:"ownerId"`
	OwnerName string        `json:"ownerName"`
	AssetId   string        `json:"assetId"`
	StashId   string        `json:"stashId"`
	AuthorId  string        `json:"authorId"`
	MediaIn   time.Duration `json:"mediaIn"`
	MediaOut  time.Duration `json:"mediaOut"`
	Duration  time.Duration `json:"duration"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	Owner     *Workspace    `json:"owner"`
	Stash     *Stash        `json:"stash"`
	Asset     *Asset        `json:"asset"`
	Author    *User         `json:"author"`
}

type LinkList []*Link

func (l LinkList) SearchId(id string) (int, *Link) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}
