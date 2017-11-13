// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package trimmer

import (
	"time"
)

// AssetState is the list of allowed values for the asset status.
// Allowed values are "nomedia", "uploading", "analyzing", "transcoding",
// "attention", "ready", "blocked", "banned", "deleting", "cleaning", "deleted"
type AssetState string

// AssetParams is the set of parameters that can be used to create an asset.
//
type AssetParams struct {
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Notes       string        `json:"notes,omitempty"`
	Copyright   string        `json:"copyright,omitempty"`
	License     string        `json:"license,omitempty"`
	Access      AccessClass   `json:"access,omitempty"`
	UUID        string        `json:"uuid,omitempty"`
	Actions     MetaValueList `json:"actions,omitempty"`
	Metadata    *MetaDocument `json:"meta,omitempty"`
	Embed       ApiEmbedFlags `json:"embed,omitempty"`
}

// AssetActivationParams is the set of parameters that can be used
// to activate a asset subscription.
//
type AssetForkParams struct {
	WorkspaceId    string        `json:"workspaceId,omitempty"`
	MediaIn        int64         `json:"mediaIn,omitempty"`
	MediaOut       int64         `json:"mediaOut,omitempty"`
	AccessClass    AccessClass   `json:"access,omitempty"`
	ExcludeRoles   MediaRoleList `json:"excludeRoles,omitempty"`
	ExcludeTags    TagLabelList  `json:"excludeTags,omitempty"`
	MetadataFilter string        `json:"metaFilter,omitempty"`
	Embed          ApiEmbedFlags `json:"embed,omitempty"`
}

// AssetSnapshotParams is the set of parameters that can be used
// to create a still-image snapshot from a video sequence.
//
type AssetSnapshotParams struct {
	MediaId  string        `json:"-"`
	VolumeId string        `json:"volumeId,omitempty"`
	Timecode string        `json:"timecode,omitempty"`
	Role     MediaRole     `json:"role,omitempty"`
	Embed    ApiEmbedFlags `json:"embed,omitempty"`
}

// AssetTrimParams is the set of parameters that can be used
// to change media in/out points and start timecode within an
// asset bundle.
//
type AssetTrimParams struct {
	MediaId  string        `json:"-"`
	Timecode string        `json:"timecode,omitempty"`
	MediaIn  time.Duration `json:"mediaIn,omitempty"`
	MediaOut time.Duration `json:"mediaOut,omitempty"`
	Embed    ApiEmbedFlags `json:"embed,omitempty"`
}

// AssetTranscodeParams is the set of parameters that can be used
// to transcode asset media.
//
type AssetTranscodeParams struct {
	Match   MediaMatch         `json:"match"`
	Options *TranscoderOptions `json:"options"`
	Embed   ApiEmbedFlags      `json:"embed,omitempty"`
}

// AssetAnalyzeParams is the set of parameters that can be used
// to analyze asset media.
//
type AssetAnalyzeParams struct {
	MediaId string           `json:"-"`
	Options *AnalyzerOptions `json:"options"`
	Embed   ApiEmbedFlags    `json:"embed,omitempty"`
}

// AssetEventType is the list of allowed values for the asset list event field.
// Allowed values are "created", "updated"
type AssetListEvent string

// AssetListParams is the set of parameters that can be used when listing assets.
type AssetListParams struct {
	ListParams
	IDs         []string       `json:"id,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	AccessClass AccessClass    `json:"access,omitempty"`
	State       AssetState     `json:"state,omitempty"`
	Event       AssetListEvent `json:"event,omitempty"`
	Embed       ApiEmbedFlags  `json:"embed,omitempty"`
}

// Asset is the resource representing a Trimmer asset.
type Asset struct {
	ID          string           `json:"assetId"`
	State       AssetState       `json:"state"`
	WorkspaceId string           `json:"workspaceId"`
	AuthorId    string           `json:"authorId"`
	OriginId    string           `json:"originId"`
	Uuid        string           `json:"uuid"`
	AccessClass AccessClass      `json:"access"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
	Statistics  *AssetStatistics `json:"stats"`
	Permissions *Permissions     `json:"perms"`
	Workspace   *Workspace       `json:"workspace"`
	Author      *User            `json:"author"`
	Origin      *AssetOrigin     `json:"origin"`
	Poster      []*MediaEmbed    `json:"poster"`
	Thumbnail   []*MediaEmbed    `json:"thumbnail"`
	Metadata    *MetaDocument    `json:"meta"`
	Version     string           `json:"version"`
}

// AssetList is representing a slice of Asset structs.
type AssetList []*Asset

func (l AssetList) SearchId(id string) (int, *Asset) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}

type AssetOrigin struct {
	WorkspaceId string `json:"workspaceId"`
	AssetId     string `json:"assetId"`
}

// AssetStatistics is the resource representing a Asset quota counters.
type AssetStatistics struct {
	Links int64 `json:"links"`
	Forks int64 `json:"forks"`
	Tags  int64 `json:"tags"`
	Media int64 `json:"media"`
	Time  int64 `json:"time"`
	Size  int64 `json:"size"`
	Files int64 `json:"files"`
}
