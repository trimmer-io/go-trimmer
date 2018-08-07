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

// Package user provides the /users APIs
package user

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	trimmer "trimmer.io/go-trimmer"
	"trimmer.io/go-trimmer/event"
	"trimmer.io/go-trimmer/media"
	"trimmer.io/go-trimmer/org"
	"trimmer.io/go-trimmer/volume"
	"trimmer.io/go-trimmer/workspace"
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

// Iter is an iterator for lists of Users.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// User returns the most recent User visited by a call to Next.
func (i *Iter) User() *trimmer.User {
	return i.Current().(*trimmer.User)
}

func Me(ctx context.Context, params *trimmer.UserParams) (*trimmer.User, error) {
	return getC().Me(ctx, params)
}

func UpdateMe(ctx context.Context, params *trimmer.UserParams) (*trimmer.User, error) {
	return getC().UpdateMe(ctx, params)
}

func Lookup(ctx context.Context, params *trimmer.UserLookupParams) (trimmer.UserList, error) {
	return getC().Lookup(ctx, params)
}

func Search(ctx context.Context, params *trimmer.UserSearchParams) *Iter {
	return getC().Search(ctx, params)
}

func UploadImage(ctx context.Context, params *trimmer.FileInfo, src io.Reader) (*trimmer.Media, error) {
	return media.UploadImage(ctx, "/users/me/media", params, src)
}

func ListMedia(ctx context.Context, params *trimmer.MediaListParams) *media.Iter {
	return getC().ListMedia(ctx, params)
}

func ListEvents(ctx context.Context, params *trimmer.EventListParams) *event.Iter {
	return getC().ListEvents(ctx, params)
}

func ListOrgs(ctx context.Context, params *trimmer.OrgListParams) *org.Iter {
	return getC().ListOrgs(ctx, params)
}

func NewWorkspace(ctx context.Context, params *trimmer.WorkspaceParams) (*trimmer.Workspace, error) {
	return getC().NewWorkspace(ctx, params)
}

func ListWorkspaces(ctx context.Context, params *trimmer.WorkspaceListParams) *workspace.Iter {
	return getC().ListWorkspaces(ctx, params)
}

func NewVolume(ctx context.Context, params *trimmer.VolumeParams) (*trimmer.Volume, error) {
	return getC().NewVolume(ctx, params)
}

func ListVolumes(ctx context.Context, params *trimmer.VolumeListParams) *volume.Iter {
	return getC().ListVolumes(ctx, params)
}

func (c Client) Me(ctx context.Context, params *trimmer.UserParams) (*trimmer.User, error) {
	v := &trimmer.User{}
	u := "/users/me"
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	err := c.B.Call(ctx, http.MethodGet, u, c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) UpdateMe(ctx context.Context, params *trimmer.UserParams) (*trimmer.User, error) {
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.User{}
	err := c.B.Call(ctx, http.MethodPatch, "/users/me", c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) Lookup(ctx context.Context, params *trimmer.UserLookupParams) (trimmer.UserList, error) {
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	type userList struct {
		trimmer.ListMeta
		Values trimmer.UserList `json:"users"`
	}
	type searchResult struct {
		Users userList `json:"users"`
	}
	res := &searchResult{}
	var err error
	if len(params.IDs)+len(params.Names) <= trimmer.LIST_GET_LIMIT {
		q := &url.Values{}
		if len(params.IDs) > 0 {
			q.Add("id", strings.Join(params.IDs, ","))
		}
		if len(params.Names) > 0 {
			q.Add("name", strings.Join(params.Names, ","))
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}
		err = c.B.Call(ctx, http.MethodGet, "/lookup?"+q.Encode(), c.Key, c.Sess, nil, nil, res)
	} else {
		err = c.B.Call(ctx, http.MethodPost, "/lookup", c.Key, c.Sess, nil, params, res)
	}
	return res.Users.Values, err
}

func (c Client) Search(ctx context.Context, params *trimmer.UserSearchParams) *Iter {
	type userList struct {
		trimmer.ListMeta
		Values trimmer.UserList `json:"users"`
	}
	type searchResult struct {
		Users userList `json:"users"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}
		q.Add("t", "user")
		if params.Query != "" {
			q.Add("q", params.Query)
		}
		if params.Fields != "" {
			q.Add("f", string(params.Fields))
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}
		params.AppendTo(q)

		lp = &params.ListParams
	}

	return &Iter{trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		res := &searchResult{}
		err := c.B.Call(ctx, http.MethodGet, fmt.Sprintf("/search?%v", b.Encode()), c.Key, c.Sess, nil, nil, res)
		ret := make([]interface{}, len(res.Users.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range res.Users.Values {
			ret[i] = v
		}

		return ret, res.Users.ListMeta, err
	})}
}

func (c Client) ListMedia(ctx context.Context, params *trimmer.MediaListParams) *media.Iter {

	type mediaList struct {
		trimmer.ListMeta
		Values trimmer.MediaList `json:"media"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

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

	return &media.Iter{Iter: trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &mediaList{}
		err := c.B.Call(ctx, http.MethodGet, "/users/me/media?"+b.Encode(), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) ListOrgs(ctx context.Context, params *trimmer.OrgListParams) *org.Iter {

	type orgList struct {
		trimmer.ListMeta
		Values trimmer.OrgList `json:"orgs"`
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

	return &org.Iter{trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &orgList{}
		err := c.B.Call(ctx, http.MethodGet, "/users/me/orgs?"+b.Encode(), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) NewWorkspace(ctx context.Context, params *trimmer.WorkspaceParams) (*trimmer.Workspace, error) {
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	u := "/users/me/workspaces"
	if params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Workspace{}
	err := c.B.Call(ctx, http.MethodPost, u, c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) ListWorkspaces(ctx context.Context, params *trimmer.WorkspaceListParams) *workspace.Iter {

	type workspaceList struct {
		trimmer.ListMeta
		Values trimmer.WorkspaceList `json:"workspaces"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

		if len(params.States) > 0 {
			q.Add("state", params.States.String())
		}
		if len(params.IDs) > 0 {
			q.Add("id", strings.Join(params.IDs, ","))
		}
		if len(params.Names) > 0 {
			q.Add("name", strings.Join(params.Names, ","))
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &workspace.Iter{trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &workspaceList{}
		err := c.B.Call(ctx, http.MethodGet, "/users/me/workspaces?"+b.Encode(), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) NewVolume(ctx context.Context, params *trimmer.VolumeParams) (*trimmer.Volume, error) {
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	u := "/users/me/volumes"
	if params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Volume{}
	err := c.B.Call(ctx, http.MethodPost, u, c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) ListVolumes(ctx context.Context, params *trimmer.VolumeListParams) *volume.Iter {

	type volumeList struct {
		trimmer.ListMeta
		Values trimmer.VolumeList `json:"volumes"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

		if len(params.States) > 0 {
			q.Add("state", params.States.String())
		}
		if len(params.Types) > 0 {
			q.Add("type", params.Types.String())
		}
		if params.AccessClass != "" {
			q.Add("access", string(params.AccessClass))
		}
		if params.Name != "" {
			q.Add("name", params.Name)
		}
		if !trimmer.IsNilUUID(params.UUID) {
			q.Add("uuid", params.UUID)
		}
		if params.Serial != "" {
			q.Add("serial", params.Serial)
		}
		if params.Provider != "" {
			q.Add("provider", string(params.Provider))
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

	return &volume.Iter{Iter: trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &volumeList{}
		err := c.B.Call(ctx, http.MethodGet, "/users/me/volumes?"+b.Encode(), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) ListEvents(ctx context.Context, params *trimmer.EventListParams) *event.Iter {

	type eventList struct {
		trimmer.ListMeta
		Values trimmer.EventList `json:"events"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

		if params.Key != "" {
			q.Add("eventKey", string(params.Key))
		}
		if params.Type != "" {
			q.Add("eventType", string(params.Type))
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &event.Iter{trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &eventList{}
		err := c.B.Call(ctx, http.MethodGet, "/users/me/events?"+b.Encode(), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}
