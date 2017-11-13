// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

// Package org provides the /orgs APIs
package org

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"

	trimmer "trimmer.io/go-trimmer"
	"trimmer.io/go-trimmer/event"
	"trimmer.io/go-trimmer/media"
	"trimmer.io/go-trimmer/member"
	"trimmer.io/go-trimmer/volume"
	"trimmer.io/go-trimmer/workspace"
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

// Iter is an iterator for lists of Orgs.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// Org returns the most recent Org visited by a call to Next.
func (i *Iter) Org() *trimmer.Org {
	return i.Current().(*trimmer.Org)
}

func Get(ctx context.Context, orgId string, params *trimmer.OrgParams) (*trimmer.Org, error) {
	return getC().Get(ctx, orgId, params)
}

func Update(ctx context.Context, orgId string, params *trimmer.OrgParams) (*trimmer.Org, error) {
	return getC().Update(ctx, orgId, params)
}

func UploadImage(ctx context.Context, orgId string, params *trimmer.FileInfo, src io.Reader) (*trimmer.Media, error) {
	uri := fmt.Sprintf("/orgs/%v/media", orgId)
	return media.UploadImage(ctx, uri, params, src)
}

func ListMedia(ctx context.Context, orgId string, params *trimmer.MediaListParams) *media.Iter {
	return getC().ListMedia(ctx, orgId, params)
}

func ListEvents(ctx context.Context, orgId string, params *trimmer.EventListParams) *event.Iter {
	return getC().ListEvents(ctx, orgId, params)
}

func NewWorkspace(ctx context.Context, orgId string, params *trimmer.WorkspaceParams) (*trimmer.Workspace, error) {
	return getC().NewWorkspace(ctx, orgId, params)
}

func ListWorkspaces(ctx context.Context, orgId string, params *trimmer.WorkspaceListParams) *workspace.Iter {
	return getC().ListWorkspaces(ctx, orgId, params)
}

func NewVolume(ctx context.Context, orgId string, params *trimmer.VolumeParams) (*trimmer.Volume, error) {
	return getC().NewVolume(ctx, orgId, params)
}

func ListVolumes(ctx context.Context, orgId string, params *trimmer.VolumeListParams) *volume.Iter {
	return getC().ListVolumes(ctx, orgId, params)
}

func GetMember(ctx context.Context, orgId, userId string) (*trimmer.Member, error) {
	return getC().GetMember(ctx, orgId, userId)
}

func NewMember(ctx context.Context, orgId, userId string, params *trimmer.MemberParams) (*trimmer.Member, error) {
	return getC().NewMember(ctx, orgId, userId, params)
}

func DeleteMember(ctx context.Context, orgId, userId string) error {
	return getC().DeleteMember(ctx, orgId, userId)
}

func ListMembers(ctx context.Context, orgId string, params *trimmer.MemberListParams) *member.Iter {
	return getC().ListMembers(ctx, orgId, params)
}

func (c Client) Get(ctx context.Context, orgId string, params *trimmer.OrgParams) (*trimmer.Org, error) {
	if orgId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/orgs/%v", orgId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Org{}
	err := c.B.Call(ctx, "GET", u, c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) Update(ctx context.Context, orgId string, params *trimmer.OrgParams) (*trimmer.Org, error) {
	if orgId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.Org{}
	err := c.B.Call(ctx, "PATCH", fmt.Sprintf("/orgs/%v", orgId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) ListMedia(ctx context.Context, orgId string, params *trimmer.MediaListParams) *media.Iter {

	if orgId == "" {
		return &media.Iter{Iter: trimmer.GetIterErr(trimmer.EIDMissing)}
	}

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
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/orgs/%v/media?%v", orgId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) ListEvents(ctx context.Context, orgId string, params *trimmer.EventListParams) *event.Iter {

	if orgId == "" {
		return &event.Iter{trimmer.GetIterErr(trimmer.EIDMissing)}
	}

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
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/orgs/%v/events?%v", orgId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) NewWorkspace(ctx context.Context, orgId string, params *trimmer.WorkspaceParams) (*trimmer.Workspace, error) {
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.Workspace{}
	err := c.B.Call(ctx, "POST", fmt.Sprintf("/orgs/%v/workspaces", orgId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) ListWorkspaces(ctx context.Context, orgId string, params *trimmer.WorkspaceListParams) *workspace.Iter {

	if orgId == "" {
		return &workspace.Iter{trimmer.GetIterErr(trimmer.EIDMissing)}
	}

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
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/orgs/%v/workspaces?%v", orgId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) NewVolume(ctx context.Context, orgId string, params *trimmer.VolumeParams) (*trimmer.Volume, error) {
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.Volume{}
	err := c.B.Call(ctx, "POST", fmt.Sprintf("/users/%v/volumes", orgId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) ListVolumes(ctx context.Context, orgId string, params *trimmer.VolumeListParams) *volume.Iter {

	if orgId == "" {
		return &volume.Iter{Iter: trimmer.GetIterErr(trimmer.EIDMissing)}
	}

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
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/orgs/%v/volumes?%v", orgId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) GetMember(ctx context.Context, orgId, userId string) (*trimmer.Member, error) {
	if orgId == "" || userId == "" {
		return nil, trimmer.EIDMissing
	}
	v := &trimmer.Member{}
	err := c.B.Call(ctx, "GET", fmt.Sprintf("/orgs/%v/members/%v", orgId, userId), c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) NewMember(ctx context.Context, orgId, userId string, params *trimmer.MemberParams) (*trimmer.Member, error) {
	if orgId == "" || userId == "" {
		return nil, trimmer.EIDMissing
	}
	v := &trimmer.Member{}
	err := c.B.Call(ctx, "PUT", fmt.Sprintf("/orgs/%v/members/%v", orgId, userId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) DeleteMember(ctx context.Context, orgId, userId string) error {
	if orgId == "" || userId == "" {
		return trimmer.EIDMissing
	}
	return c.B.Call(ctx, "DELETE", fmt.Sprintf("/orgs/%v/members/%v", orgId, userId), c.Key, c.Sess, nil, nil, nil)
}

func (c Client) ListMembers(ctx context.Context, orgId string, params *trimmer.MemberListParams) *member.Iter {
	if orgId == "" {
		return &member.Iter{trimmer.GetIterErr(trimmer.EIDMissing)}
	}

	type memberList struct {
		trimmer.ListMeta
		Values trimmer.VolumeList `json:"members"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

		if len(params.State) > 0 {
			q.Add("state", string(params.State))
		}
		if len(params.Role) > 0 {
			q.Add("role", string(params.Role))
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &member.Iter{trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &memberList{}
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/orgs/%v/members?%v", orgId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}
