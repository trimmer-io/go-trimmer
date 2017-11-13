// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package trimmer

import (
	"time"
)

// MountParams is the set of parameters that can be used to mount a volume
// into a workspace.
//
type MountParams struct {
	Readonly  VolumeReadonlyState `json:"readonly,omitempty"`
	RoleMatch MediaRoleMatch      `json:"roleMatch,omitempty"`
	Embed     ApiEmbedFlags       `json:"embed,omitempty"`
}

// MountListParams is the set of parameters that can be used when listing volumes.
type MountListParams struct {
	ListParams
	States   VolumeStateList     `json:"state,omitempty"`
	Types    VolumeTypeList      `json:"type,omitempty"`
	Provider string              `json:"provider,omitempty"`
	Region   string              `json:"region,omitempty"`
	Readonly VolumeReadonlyState `json:"readonly,omitempty"`
	Online   VolumeOnlineState   `json:"online,omitempty"`
	Embed    ApiEmbedFlags       `json:"embed,omitempty"`
}

// Mount is the secondary resource representing a Trimmer volume/workspace mount.
type Mount struct {
	ID          string        `json:"mountId"`
	WorkspaceId string        `json:"workspaceId"`
	AuthorId    string        `json:"authorId"`
	VolumeId    string        `json:"volumeId"`
	CreatedAt   time.Time     `json:"createdAt"`
	Volume      *Volume       `json:"volume"`
	Workspace   *Workspace    `json:"workspace"`
	Author      *User         `json:"author"`
	Statistics  *VolumeCounts `json:"stats"`
}

// MountList is representing a slice of Mount structs.
type MountList []*Mount

func (l MountList) SearchId(id string) (int, *Mount) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}
