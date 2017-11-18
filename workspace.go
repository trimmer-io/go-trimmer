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
	"strings"
	"time"
)

// WorkspaceState is the list of allowed values for the workspace status.
// Allowed values are "inactive", "active", "blocked",
// "banned", "deleting", "cleaning", "deleted".
type WorkspaceState string
type WorkspaceStateList []WorkspaceState

func (l WorkspaceStateList) Contains(f WorkspaceState) bool {
	if len(l) == 0 {
		return false
	}
	for _, v := range l {
		if v == f {
			return true
		}
	}
	return false
}

func (l *WorkspaceStateList) Add(f WorkspaceState) {
	for !l.Contains(f) {
		*l = append(*l, f)
	}
}

func (l *WorkspaceStateList) Del(f WorkspaceState) {
	i := -1
	for j, v := range *l {
		if v == f {
			i = j
		}
	}
	if i > -1 {
		(*l)[i] = (*l)[len(*l)-1]
		*l = (*l)[:len(*l)-1]
	}
}

func (l WorkspaceStateList) String() string {
	if len(l) == 0 {
		return ""
	}
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = string(v)
	}
	return strings.Join(s, ",")
}

// WorkspaceParams is the set of parameters that can be used to create and
// update a workspace.
//
type WorkspaceParams struct {
	Name        string        `json:"name,omitempty"` // create only, required
	DisplayName string        `json:"displayName,omitempty"`
	Description string        `json:"description,omitempty"`
	Company     string        `json:"company,omitempty"`
	Location    string        `json:"location,omitempty"`
	Homepage    string        `json:"homepage,omitempty"`
	ImageId     string        `json:"imageId,omitempty"` // update only
	Embed       ApiEmbedFlags `json:"embed,omitempty"`
}

// WorkspaceListParams is the set of parameters that can be used when listing workspaces.
type WorkspaceListParams struct {
	ListParams
	Names  []string           `json:"name,omitempty"`
	IDs    []string           `json:"id,omitempty"`
	States WorkspaceStateList `json:"state,omitempty"`
	Embed  ApiEmbedFlags      `json:"embed,omitempty"`
}

// Workspace is the resource representing a Trimmer workspace.
type Workspace struct {
	ID          string               `json:"workspaceId"`
	State       WorkspaceState       `json:"state"`
	Name        string               `json:"name"`
	DisplayName string               `json:"displayName"`
	AccountId   string               `json:"accountId"`
	Company     string               `json:"company"`
	Description string               `json:"description"`
	Homepage    string               `json:"homepage"`
	ImageId     string               `json:"imageId"`
	Location    string               `json:"location"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
	Permissions *Permissions         `json:"perms"`
	Media       *MediaEmbed          `json:"media"`
	Statistics  *WorkspaceStatistics `json:"stats"`
	Membership  *Member              `json:"membership"`
	Org         *Org                 `json:"org"`
	User        *User                `json:"user"`
}

// WorkspaceList is representing a slice of Workspace structs.
type WorkspaceList []*Workspace

func (l WorkspaceList) SearchId(id string) (int, *Workspace) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}

// WorkspaceStatistics is the resource representing a Workspace counters.
type WorkspaceStatistics struct {
	Members int64 `json:"members"`
	Stashes int64 `json:"stashes"`
	Assets  int64 `json:"assets"`
	Time    int64 `json:"time"`
	Media   int64 `json:"media"`
	Files   int64 `json:"files"`
	Size    int64 `json:"size"`
}
