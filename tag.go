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

const (
	TAG_MAX_CONTENT_LENGTH = 140
	TAG_MAX_DATA_LENGTH    = 8192
)

// TagLabel is the list of allowed values for tag labels.
type TagLabel string
type TagLabelList []TagLabel

func (l TagLabelList) Contains(f TagLabel) bool {
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

func (l *TagLabelList) Add(f TagLabel) {
	for !l.Contains(f) {
		*l = append(*l, f)
	}
}

func (l *TagLabelList) Del(f TagLabel) {
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

func (l TagLabelList) String() string {
	if len(l) == 0 {
		return ""
	}
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = string(v)
	}
	return strings.Join(s, ",")
}

// TagParams is the set of parameters that can be used to create and
// update a tag.
//
type TagParams struct {
	Content     string           `json:"content,omitempty"`
	Start       time.Duration    `json:"start,omitempty"`
	Duration    time.Duration    `json:"duration,omitempty"`
	AccessClass AccessClass      `json:"access,omitempty"`
	Label       TagLabel         `json:"label,omitempty"`
	Data        *json.RawMessage `json:"data,omitempty"`
	Embed       ApiEmbedFlags    `json:"embed,omitempty"`
}

// TagListParams is the set of parameters that can be used when listing tages.
type TagListParams struct {
	ListParams
	IDs         []string      `json:"id,omitempty"`
	AuthorId    string        `json:"authorId,omitempty"`
	AccessClass AccessClass   `json:"access,omitempty"`
	Labels      TagLabelList  `json:"label,omitempty"`
	From        int64         `json:"from,omitempty"`
	To          int64         `json:"to,omitempty"`
	Embed       ApiEmbedFlags `json:"embed,omitempty"`
}

// Tag is the resource representing a Trimmer tag.
type Tag struct {
	ID          string          `json:"tagId"`
	WorkspaceId string          `json:"workspaceId"`
	AuthorId    string          `json:"authorId"`
	AssetId     string          `json:"assetId"`
	ParentId    string          `json:"parentId"`
	RootId      string          `json:"rootId"`
	Label       TagLabel        `json:"label"`
	AccessClass AccessClass     `json:"access"`
	Start       time.Duration   `json:"start"`
	Duration    time.Duration   `json:"duration"`
	Content     string          `json:"content"`
	Data        json.RawMessage `json:"data"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	Author      *User           `json:"author"`
	Workspace   *Workspace      `json:"workspace"`
	Asset       *Asset          `json:"asset"`
	ParentTag   *Tag            `json:"parent"`
	RootTag     *Tag            `json:"root"`
}

type TagList []*Tag

func (l TagList) SearchId(id string) (int, *Tag) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}
