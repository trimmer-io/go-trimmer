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
