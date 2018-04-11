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

// ReplicaParams is the set of parameters that can be used to manage media
// replicas on volumes.
//
type ReplicaParams struct {
	Embed ApiEmbedFlags `json:"embed,omitempty"`
}

// ReplicaListParams is the set of parameters that can be used to list media
// replicas on volumes.
//
type ReplicaListParams struct {
	ListParams
	Type     VolumeType          `json:"type,omitempty"`
	Provider string              `json:"provider,omitempty"`
	Region   string              `json:"region,omitempty"`
	Readonly VolumeReadonlyState `json:"readonly,omitempty"`
	Online   VolumeOnlineState   `json:"online,omitempty"`
	Embed    ApiEmbedFlags       `json:"embed,omitempty"`
}

// ReplicaDeleteParams is the set of parameters that can be used to delete media
// replicas from volumes.
//
type ReplicaDeleteParams struct {
	Wipe bool `json:"wipeMedia,omitempty"`
}

// Replica is the secondary resource representing a Trimmer volume/media relation.
type Replica struct {
	MediaEmbed
	State    MediaState `json:"state"`
	VolumeId string     `json:"volumeId"`
	JobId    string     `json:"jobId"`
	Volume   *Volume    `json:"volume"`
	Job      *Job       `json:"job"`
}

// ReplicaList is representing a slice of Replica structs.
type ReplicaList []*Replica

func (l ReplicaList) SearchId(id string) (int, *Replica) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}
