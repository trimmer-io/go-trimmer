// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package trimmer

import (
	"net/url"
	"strconv"
	"time"
)

const (
	minId  = "minId"
	maxId  = "maxId"
	before = "before"
	after  = "after"
)

const (
	LIST_GET_LIMIT = 5
	LIST_MAX_LIMIT = 200
)

// ListParams is the structure that contains the common properties
// of any *ListParams structure.
type ListParams struct {
	Count  int       `json:"count,omitempty"`
	Before time.Time `json:"before,omitempty"`
	After  time.Time `json:"after,omitempty"`
	MaxId  string    `json:"maxId,omitempty"`
	MinId  string    `json:"minId,omitempty"`
}

// ListMeta is the structure that contains the common properties
// of List iterators. The Count property is only populated if the
// total_count include option is passed in (see tests for example).
type ListMeta struct {
	Count int    `json:"count"`
	MaxId string `json:"maxId"`
	MinId string `json:"minId"`
	From  int64  `json:"from"` // timecode in ms (tags only)
	To    int64  `json:"to"`   // timecode in ms (tags only)
	Total int    `json:"-"`    // internal use, not sent by server
	More  bool   `json:"-"`    // internal use, not sent by server
}

// AppendTo adds the common parameters to the query string values.
func (p *ListParams) AppendTo(q *url.Values) {

	if p.MaxId != "" {
		q.Add(maxId, p.MaxId)
	}

	if p.MinId != "" {
		q.Add(minId, p.MinId)
	}

	if !p.Before.IsZero() {
		q.Add(before, p.Before.String())
	}

	if !p.After.IsZero() {
		q.Add(after, p.After.String())
	}

	if p.Count > 0 {
		if p.Count > LIST_MAX_LIMIT {
			p.Count = LIST_MAX_LIMIT
		}

		q.Add("count", strconv.Itoa(p.Count))
	}
}

func IsNilUUID(u string) bool {
	return u == "" || u == "00000000-0000-0000-0000-000000000000" || u == "00000000000000000000000000000000"
}
