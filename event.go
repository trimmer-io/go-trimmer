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

package trimmer

import (
	"time"
)

// EventKey is the list of allowed values for events.
type EventKey string

// EventType is the list of allowed values for groups of events.
// Allowed values are "account", "acl", "application", "asset",
// "action", "auth", "billing", "media", "invite", "job", "stash",
// "tag", "workspace", "team", "volume", "organization".
type EventType string

// EventListParams is the set of parameters that can be used when listing events.
type EventListParams struct {
	ListParams
	Key   EventKey      `json:"eventKey,omitempty"`
	Type  EventType     `json:"eventType,omitempty"`
	Embed ApiEmbedFlags `json:"embed,omitempty"`
}

// Event is the resource representing a Trimmer audit event.
type Event struct {
	ID          string            `json:"eventId"`
	AccountId   string            `json:"accountId"`
	AuthorId    string            `json:"authorId"`
	WorkspaceId string            `json:"workspaceId"`
	CreatedAt   time.Time         `json:"createdAt"`
	Key         EventKey          `json:"eventKey"`
	Type        EventType         `json:"eventType"`
	Data        map[string]string `json:"data"`
}

type EventList []*Event

func (l EventList) SearchId(id string) (int, *Event) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}
