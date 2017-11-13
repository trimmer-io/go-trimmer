// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package trimmer

import (
	"encoding/json"
	"strings"
	"time"
)

// UserState is the list of allowed values for the user's account status.
// Allowed values are "parked", created", "invited", "active", "inactive",
// "expired", "blocked", "banned", "deleting", "cleaning", "deleted", "rejected".
type UserState string

// UserSearchFields is the list of allowed values for the user's search fields.
// Allowed values are "default", "all", "name", "displayName", "email",
// "location", "description"
type UserSearchFields string

// UserParams is the set of parameters that can be used to fetch a user.
//
type UserParams struct {
	DisplayName string        `json:"displayName,omitempty"`
	ImageId     string        `json:"imageId,omitempty"`
	Homepage    string        `json:"homepage,omitempty"`
	Location    string        `json:"location,omitempty"`
	Description string        `json:"description,omitempty"`
	Embed       ApiEmbedFlags `json:"embed,omitempty"`
}

// UserListParams is the set of parameters that can be used when looking for
// multiple users with a known name or id. Names and IDs can be mixed. The
// maximum total number is limited to 200 entries per call. When less than 5 items
// are requested a GET is performed, otherwise POST is used.
type UserLookupParams struct {
	Names []string      `json:"name,omitempty"`
	IDs   []string      `json:"id,omitempty"`
	Embed ApiEmbedFlags `json:"embed,omitempty"`
}

// UserSearchParams is the set of parameters that can be used when searching users.
type UserSearchParams struct {
	ListParams
	Query  string
	Fields UserSearchFields
	Embed  ApiEmbedFlags `json:"embed,omitempty"`
}

// User is the resource representing a Trimmer user.
type User struct {
	ID           string          `json:"userId"`
	Name         string          `json:"name"`
	DisplayName  string          `json:"displayName"`
	ImageId      string          `json:"imageId"`
	Homepage     string          `json:"homepage"`
	Location     string          `json:"location"`
	Description  string          `json:"description"`
	State        UserState       `json:"state"`
	Language     string          `json:"lang"`
	MaxClearance AccessClass     `json:"maxClearance"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
	Media        *MediaEmbed     `json:"media"`
	Permissions  *Permissions    `json:"perms"`
	Statistics   *UserStatistics `json:"stats"`
}

// UserList is representing a slice of User structs.
type UserList []*User

func (l UserList) SearchId(id string) (int, *User) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}

// UserStatistics is the resource representing a user account counters.
type UserStatistics struct {
	Workspaces int64 `json:"workspaces"`
	Volumes    int64 `json:"volumes"`
	Members    int64 `json:"members"`
	Teams      int64 `json:"teams"`
	Assets     int64 `json:"assets"`
	Time       int64 `json:"time"`
	Media      int64 `json:"media"`
	Size       int64 `json:"size"`
	Files      int64 `json:"files"`
}

// Custom Marshaller for UserLookupParams because the API expects a comma
// separated list
func (p *UserLookupParams) MarshalJSON() ([]byte, error) {

	if LogLevel > 1 {
		// warn when there are more than max
		if len(p.Names) > LIST_MAX_LIMIT {
			Logger.Println("WARN: too many names in user lookup, will be capped.")
		}
		// warn when there are more than max
		if len(p.IDs) > LIST_MAX_LIMIT {
			Logger.Println("WARN: too many ids in user lookup, will be capped.")
		}
	}

	return json.Marshal(struct {
		Type  string `json:"type"`
		Name  string `json:"name,omitempty"`
		Id    string `json:"id,omitempty"`
		Embed string `json:"embed,omitempty"`
	}{
		Type:  "user",
		Name:  strings.Join(p.Names, ","),
		Id:    strings.Join(p.IDs, ","),
		Embed: p.Embed.String(),
	})
}
