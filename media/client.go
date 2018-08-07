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

// Package media provides the /media APIs
package media

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	trimmer "trimmer.io/go-trimmer"
	"trimmer.io/go-trimmer/profile"
	"trimmer.io/go-trimmer/replica"
	"trimmer.io/go-trimmer/volume"
)

// Client is used to invoke /users APIs.
type Client struct {
	B            trimmer.Backend
	CDN          trimmer.Backend
	Key          trimmer.ApiKey
	Sess         *trimmer.Session
	lastProgress time.Time
}

func getC() Client {
	return Client{trimmer.GetBackend(trimmer.APIBackend), trimmer.GetBackend(trimmer.CDNBackend), trimmer.Key, &trimmer.LoginSession, time.Time{}}
}

// Iter is an iterator for lists of Media.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// Media returns the most recent Media visited by a call to Next.
func (i *Iter) Media() *trimmer.Media {
	return i.Current().(*trimmer.Media)
}

func Get(ctx context.Context, mediaId string, params *trimmer.MediaParams) (*trimmer.Media, error) {
	return getC().Get(ctx, mediaId, params)
}

func Update(ctx context.Context, mediaId string, params *trimmer.MediaParams) (*trimmer.Media, error) {
	return getC().Update(ctx, mediaId, params)
}

func Delete(ctx context.Context, mediaId string) error {
	return getC().Delete(ctx, mediaId)
}

func CompleteUpload(ctx context.Context, mediaId string, params *trimmer.MediaUploadCompletionParams) (*trimmer.Media, error) {
	return getC().CompleteUpload(ctx, mediaId, params)
}

func ListProfiles(ctx context.Context, mediaId string, params *trimmer.ProfileListParams) *profile.Iter {
	return getC().ListProfiles(ctx, mediaId, params)
}

func ListReplicas(ctx context.Context, mediaId string, params *trimmer.ReplicaListParams) *replica.Iter {
	return getC().ListReplicas(ctx, mediaId, params)
}

func NewReplica(ctx context.Context, mediaId, volumeId string, params *trimmer.ReplicaParams) (*trimmer.Job, error) {
	return getC().NewReplica(ctx, mediaId, volumeId, params)
}

func RegisterReplica(ctx context.Context, mediaId, volumeId string, params *trimmer.ReplicaParams) (*trimmer.Replica, error) {
	return getC().RegisterReplica(ctx, mediaId, volumeId, params)
}

func DeleteReplica(ctx context.Context, mediaId, volumeId string, params *trimmer.ReplicaDeleteParams) error {
	return getC().DeleteReplica(ctx, mediaId, volumeId, params)
}

func (c Client) Get(ctx context.Context, mediaId string, params *trimmer.MediaParams) (*trimmer.Media, error) {
	if mediaId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/media/%v", mediaId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Media{}
	err := c.B.Call(ctx, http.MethodGet, u, c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) Update(ctx context.Context, mediaId string, params *trimmer.MediaParams) (*trimmer.Media, error) {
	if mediaId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	StripMetadataUrls(params.Attr)
	v := &trimmer.Media{}
	err := c.B.Call(ctx, http.MethodPatch, fmt.Sprintf("/media/%v", mediaId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) Delete(ctx context.Context, mediaId string) error {
	if mediaId == "" {
		return trimmer.EIDMissing
	}
	return c.B.Call(ctx, http.MethodDelete, fmt.Sprintf("/media/%v", mediaId), c.Key, c.Sess, nil, nil, nil)
}

func (c Client) CompleteUpload(ctx context.Context, mediaId string, params *trimmer.MediaUploadCompletionParams) (*trimmer.Media, error) {
	if mediaId == "" {
		return nil, trimmer.EIDMissing
	}
	v := &trimmer.Media{}
	err := c.B.Call(ctx, http.MethodPost, fmt.Sprintf("/media/%v/complete", mediaId), c.Key, c.Sess, nil, params, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (c Client) ListProfiles(ctx context.Context, mediaId string, params *trimmer.ProfileListParams) *profile.Iter {
	type profileList struct {
		trimmer.ListMeta
		Values trimmer.ProfileList `json:"profiles"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

		if len(params.Types) > 0 {
			q.Add("type", params.Types.String())
		}
		if len(params.Formats) > 0 {
			q.Add("format", params.Formats.String())
		}
		if len(params.Families) > 0 {
			q.Add("family", params.Families.String())
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &profile.Iter{Iter: trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &profileList{}
		err := c.B.Call(ctx, http.MethodGet, fmt.Sprintf("/media/%v/profiles?%s", mediaId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) ListReplicas(ctx context.Context, mediaId string, params *trimmer.ReplicaListParams) *replica.Iter {
	if mediaId == "" {
		return &replica.Iter{Iter: trimmer.GetIterErr(trimmer.EIDMissing)}
	}

	type replicaList struct {
		trimmer.ListMeta
		Values trimmer.ReplicaList `json:"replica"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

		if params.Type != "" {
			q.Add("type", string(params.Type))
		}
		if params.Provider != "" {
			q.Add("provider", params.Provider)
		}
		if params.Region != "" {
			q.Add("region", params.Region)
		}
		if params.Readonly != "" {
			q.Add("readonly", strconv.FormatBool(params.Readonly == volume.VolumeReadonlyStateOn))
		}
		if params.Online != "" {
			q.Add("online", strconv.FormatBool(params.Online == volume.VolumeOnlineStateOn))
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &replica.Iter{Iter: trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &replicaList{}
		err := c.B.Call(ctx, http.MethodGet, fmt.Sprintf("/media/%v/replicas?%s", mediaId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) NewReplica(ctx context.Context, mediaId, volumeId string, params *trimmer.ReplicaParams) (*trimmer.Job, error) {
	if mediaId == "" || volumeId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/media/%v/replicas/%v", mediaId, volumeId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Job{}
	err := c.B.Call(ctx, http.MethodPost, u, c.Key, c.Sess, nil, params, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (c Client) RegisterReplica(ctx context.Context, mediaId, volumeId string, params *trimmer.ReplicaParams) (*trimmer.Replica, error) {
	if mediaId == "" || volumeId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/media/%v/replicas/%v", mediaId, volumeId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Replica{}
	err := c.B.Call(ctx, http.MethodPut, u, c.Key, c.Sess, nil, params, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (c Client) DeleteReplica(ctx context.Context, mediaId, volumeId string, params *trimmer.ReplicaDeleteParams) error {
	if mediaId == "" || volumeId == "" {
		return trimmer.EIDMissing
	}
	return c.B.Call(ctx, http.MethodDelete, fmt.Sprintf("/media/%v/replicas/%v", mediaId, volumeId), c.Key, c.Sess, nil, params, nil)
}
