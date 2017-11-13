// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

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
