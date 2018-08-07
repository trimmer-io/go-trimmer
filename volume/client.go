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

// Package volume provides the /volumes APIs
package volume

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	trimmer "trimmer.io/go-trimmer"
	"trimmer.io/go-trimmer/mount"
	"trimmer.io/go-trimmer/replica"
)

// Client is used to invoke /orgs APIs.
type Client struct {
	B    trimmer.Backend
	Key  trimmer.ApiKey
	Sess *trimmer.Session
}

func getC() Client {
	return Client{trimmer.GetBackend(trimmer.APIBackend), trimmer.Key, &trimmer.LoginSession}
}

// Iter is an iterator for lists of Volumes.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// Volume returns the most recent Volume visited by a call to Next.
func (i *Iter) Volume() *trimmer.Volume {
	return i.Current().(*trimmer.Volume)
}

func Get(ctx context.Context, volId string, params *trimmer.VolumeParams) (*trimmer.Volume, error) {
	return getC().Get(ctx, volId, params)
}

func Update(ctx context.Context, volId string, params *trimmer.VolumeParams) (*trimmer.Volume, error) {
	return getC().Update(ctx, volId, params)
}

func Delete(ctx context.Context, volId string) error {
	return getC().Delete(ctx, volId)
}

func ListReplicas(ctx context.Context, volId string, params *trimmer.MediaListParams) *replica.Iter {
	return getC().ListReplicas(ctx, volId, params)
}

func ListMounts(ctx context.Context, volId string, params *trimmer.WorkspaceListParams) *mount.Iter {
	return getC().ListMounts(ctx, volId, params)
}

func (c Client) Get(ctx context.Context, volId string, params *trimmer.VolumeParams) (*trimmer.Volume, error) {
	if volId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/volumes/%v", volId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Volume{}
	err := c.B.Call(ctx, http.MethodGet, u, c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) Update(ctx context.Context, volId string, params *trimmer.VolumeParams) (*trimmer.Volume, error) {
	if volId == "" {
		return nil, trimmer.EIDMissing
	}

	if params == nil {
		return nil, trimmer.ENilPointer
	}

	u := fmt.Sprintf("/volumes/%v", volId)
	if params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Volume{}
	err := c.B.Call(ctx, http.MethodPatch, u, c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) Delete(ctx context.Context, volId string) error {
	if volId == "" {
		return trimmer.EIDMissing
	}
	err := c.B.Call(ctx, http.MethodDelete, fmt.Sprintf("/volumes/%v", volId), c.Key, c.Sess, nil, nil, nil)
	return err
}

func (c Client) ListReplicas(ctx context.Context, volId string, params *trimmer.MediaListParams) *replica.Iter {
	if volId == "" {
		return &replica.Iter{Iter: trimmer.GetIterErr(trimmer.EIDMissing)}
	}

	type replicaList struct {
		trimmer.ListMeta
		Values trimmer.ReplicaList `json:"replicas"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

		if params.WorkspaceId != "" {
			q.Add("workspaceId", params.WorkspaceId)
		}
		if params.AuthorId != "" {
			q.Add("authorId", params.AuthorId)
		}
		if len(params.States) > 0 {
			q.Add("state", params.States.String())
		}
		if len(params.Types) > 0 {
			q.Add("type", params.Types.String())
		}
		if len(params.Formats) > 0 {
			q.Add("format", params.Formats.String())
		}
		if len(params.Families) > 0 {
			q.Add("family", params.Families.String())
		}
		if len(params.Roles) > 0 {
			q.Add("role", params.Roles.String())
		}
		if len(params.Relations) > 0 {
			q.Add("relation", params.Relations.String())
		}
		if params.Kind != "" {
			q.Add("kind", string(params.Kind))
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &replica.Iter{Iter: trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &replicaList{}
		err := c.B.Call(ctx, http.MethodGet, fmt.Sprintf("/volumes/%v/media?%v", volId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) ListMounts(ctx context.Context, volId string, params *trimmer.WorkspaceListParams) *mount.Iter {
	if volId == "" {
		return &mount.Iter{Iter: trimmer.GetIterErr(trimmer.EIDMissing)}
	}

	type eventList struct {
		trimmer.ListMeta
		Values trimmer.MountList `json:"mounts"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

		if len(params.States) > 0 {
			q.Add("state", params.States.String())
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &mount.Iter{Iter: trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &eventList{}
		err := c.B.Call(ctx, http.MethodGet, fmt.Sprintf("/volumes/%v/mounts?%v", volId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) GetManifest(ctx context.Context, volId string) (*trimmer.VolumeManifest, error) {
	if volId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/volumes/%v/manifest", volId)
	v := &trimmer.VolumeManifest{}
	err := c.B.Call(ctx, http.MethodGet, u, c.Key, c.Sess, nil, nil, v)
	return v, err
}
