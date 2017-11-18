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

// Package tag provides the /tags APIs
package tag

import (
	"context"
	"fmt"
	"net/url"

	trimmer "trimmer.io/go-trimmer"
)

// Client is used to invoke /users APIs.
type Client struct {
	B    trimmer.Backend
	Key  trimmer.ApiKey
	Sess *trimmer.Session
}

func getC() Client {
	return Client{trimmer.GetBackend(trimmer.APIBackend), trimmer.Key, &trimmer.LoginSession}
}

// Iter is an iterator for lists of Tags.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// Tag returns the most recent Tag visited by a call to Next.
func (i *Iter) Tag() *trimmer.Tag {
	return i.Current().(*trimmer.Tag)
}

func Get(ctx context.Context, tagId string, params *trimmer.TagParams) (*trimmer.Tag, error) {
	return getC().Get(ctx, tagId, params)
}

func Update(ctx context.Context, tagId string, params *trimmer.TagParams) (*trimmer.Tag, error) {
	return getC().Update(ctx, tagId, params)
}

func Delete(ctx context.Context, tagId string) error {
	return getC().Delete(ctx, tagId)
}

func Reply(ctx context.Context, tagId string, params *trimmer.TagParams) (*trimmer.Tag, error) {
	return getC().Reply(ctx, tagId, params)
}

func ListReplies(ctx context.Context, tagId string, params *trimmer.TagListParams) *Iter {
	return getC().ListReplies(ctx, tagId, params)
}

func (c Client) Get(ctx context.Context, tagId string, params *trimmer.TagParams) (*trimmer.Tag, error) {
	if tagId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/tags/%v", tagId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Tag{}
	err := c.B.Call(ctx, "GET", u, c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) Update(ctx context.Context, tagId string, params *trimmer.TagParams) (*trimmer.Tag, error) {
	if tagId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.Tag{}
	u := fmt.Sprintf("/tags/%v", tagId)
	err := c.B.Call(ctx, "PATCH", u, c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) Delete(ctx context.Context, tagId string) error {
	if tagId == "" {
		return trimmer.EIDMissing
	}
	u := fmt.Sprintf("/tags/%v", tagId)
	return c.B.Call(ctx, "DELETE", u, c.Key, c.Sess, nil, nil, nil)
}

func (c Client) Reply(ctx context.Context, tagId string, params *trimmer.TagParams) (*trimmer.Tag, error) {
	if tagId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.Tag{}
	u := fmt.Sprintf("/tags/%v/replies", tagId)
	err := c.B.Call(ctx, "POST", u, c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) ListReplies(ctx context.Context, tagId string, params *trimmer.TagListParams) *Iter {
	if tagId == "" {
		return &Iter{trimmer.GetIterErr(trimmer.EIDMissing)}
	}

	type tagList struct {
		trimmer.ListMeta
		Values trimmer.TagList `json:"replies"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}
		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &Iter{trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &tagList{}
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/tags/%v/replies?%v", tagId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}
