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
	"encoding/json"
	"time"
)

// JobState is the list of allowed values for the job state of a status object.
// Allowed values are "created", "queued", "running", "complete", "failed", "aborted".
type JobState string

// JobType is the list of allowed values for the operation in progress.
type JobType string

// JobListParams is the set of parameters that can be used when listing workspaces.
type JobListParams struct {
	ListParams
	State    JobState      `json:"state,omitempty"`
	Type     JobType       `json:"type,omitempty"`
	Queue    string        `json:"queue,omitempty"`
	AuthorId string        `json:"authorId,omitempty"`
	MediaId  string        `json:"mediaId,omitempty"`
	VolumeId string        `json:"volumeId,omitempty"`
	Embed    ApiEmbedFlags `json:"embed,omitempty"`
}

// JobParams is the set of parameters that can be used to update jobs.
//
type JobParams struct {
	Progress int           `json:"progress,omitempty"`
	State    JobState      `json:"state,omitempty"`
	Priority int           `json:"priority,omitempty"`
	Embed    ApiEmbedFlags `json:"embed,omitempty"`
}

// Job is the resource representing an ongoing activity on Trimmer.
type Job struct {
	ID          string           `json:"jobId"`
	State       JobState         `json:"state"`
	Type        JobType          `json:"type"`
	Queue       string           `json:"queue"`
	AccountId   string           `json:"accountId"`
	WorkspaceId string           `json:"workspaceId"`
	AuthorId    string           `json:"authorId"`
	MediaCount  int              `json:"mediaCount"`
	AssetId     string           `json:"assetId"`
	MediaId     string           `json:"mediaId"`
	VolumeId    string           `json:"volumeId"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
	ExpiresAt   time.Time        `json:"expiresAt"`
	Progress    int              `json:"progress"`
	Asset       *Asset           `json:"asset"`
	Media       *Media           `json:"media"`
	Volume      *Volume          `json:"volume"`
	Account     *User            `json:"account"`
	Workspace   *Workspace       `json:"workspace"`
	Author      *User            `json:"author"`
	Error       *TrimmerError    `json:"error"`
	Statistics  *json.RawMessage `json:"stats"`
	Options     *json.RawMessage `json:"options"`
}

// JobList is representing a slice of Job structs.
type JobList []*Job

func (l JobList) SearchId(id string) (int, *Job) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}
