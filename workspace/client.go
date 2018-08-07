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

// Package workspace provides the /workspaces APIs
package workspace

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"

	trimmer "trimmer.io/go-trimmer"
	"trimmer.io/go-trimmer/asset"
	"trimmer.io/go-trimmer/event"
	"trimmer.io/go-trimmer/job"
	"trimmer.io/go-trimmer/media"
	"trimmer.io/go-trimmer/member"
	"trimmer.io/go-trimmer/mount"
	"trimmer.io/go-trimmer/profile"
	"trimmer.io/go-trimmer/stash"
	"trimmer.io/go-trimmer/volume"
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

// Iter is an iterator for lists of Workspaces.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// Workspace returns the most recent Workspace visited by a call to Next.
func (i *Iter) Workspace() *trimmer.Workspace {
	return i.Current().(*trimmer.Workspace)
}

func Get(ctx context.Context, workId string, params *trimmer.WorkspaceParams) (*trimmer.Workspace, error) {
	return getC().Get(ctx, workId, params)
}

func Update(ctx context.Context, workId string, params *trimmer.WorkspaceParams) (*trimmer.Workspace, error) {
	return getC().Update(ctx, workId, params)
}

func Delete(ctx context.Context, workId string) error {
	return getC().Delete(ctx, workId)
}

func ListAssets(ctx context.Context, workId string, params *trimmer.AssetListParams) *asset.Iter {
	return getC().ListAssets(ctx, workId, params)
}

func NewAsset(ctx context.Context, workId string, params *trimmer.AssetParams) (*trimmer.Asset, error) {
	return getC().NewAsset(ctx, workId, params)
}

func ListStashes(ctx context.Context, workId string, params *trimmer.StashListParams) *stash.Iter {
	return getC().ListStashes(ctx, workId, params)
}

func NewStash(ctx context.Context, workId string, params *trimmer.StashParams) (*trimmer.Stash, error) {
	return getC().NewStash(ctx, workId, params)
}

func UploadImage(ctx context.Context, workId string, params *trimmer.FileInfo, src io.Reader) (*trimmer.Media, error) {
	uri := fmt.Sprintf("/workspaces/%v/media", workId)
	return media.UploadImage(ctx, uri, params, src)
}

func ListMedia(ctx context.Context, workId string, params *trimmer.MediaListParams) *media.Iter {
	return getC().ListMedia(ctx, workId, params)
}

func ListEvents(ctx context.Context, workId string, params *trimmer.EventListParams) *event.Iter {
	return getC().ListEvents(ctx, workId, params)
}

func ListProfiles(ctx context.Context, workId string, params *trimmer.ProfileListParams) *profile.Iter {
	return getC().ListProfiles(ctx, workId, params)
}

func ListJobs(ctx context.Context, workId string, params *trimmer.JobListParams) *job.Iter {
	return getC().ListJobs(ctx, workId, params)
}

func ListMounts(ctx context.Context, workId string, params *trimmer.MountListParams) *mount.Iter {
	return getC().ListMounts(ctx, workId, params)
}

func MountVolume(ctx context.Context, workId, volId string, params *trimmer.MountParams) error {
	return getC().MountVolume(ctx, workId, volId, params)
}

func UnmountVolume(ctx context.Context, workId, volId string) error {
	return getC().UnmountVolume(ctx, workId, volId)
}

func ScanVolume(ctx context.Context, workId, volId string, params *trimmer.VolumeScanParams) (*trimmer.Job, error) {
	return getC().ScanVolume(ctx, workId, volId, params)
}

func ClearVolume(ctx context.Context, workId, volId string, params *trimmer.VolumeClearParams) (*trimmer.Job, error) {
	return getC().ClearVolume(ctx, workId, volId, params)
}

func GetMember(ctx context.Context, workId, userId string) (*trimmer.Member, error) {
	return getC().GetMember(ctx, workId, userId)
}

func NewMember(ctx context.Context, workId, userId string, params *trimmer.MemberParams) (*trimmer.Member, error) {
	return getC().NewMember(ctx, workId, userId, params)
}

func DeleteMember(ctx context.Context, workId, userId string) error {
	return getC().DeleteMember(ctx, workId, userId)
}

func ListMembers(ctx context.Context, workId string, params *trimmer.MemberListParams) *member.Iter {
	return getC().ListMembers(ctx, workId, params)
}

func (c Client) Get(ctx context.Context, workId string, params *trimmer.WorkspaceParams) (*trimmer.Workspace, error) {
	if workId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/workspaces/%v", workId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Workspace{}
	err := c.B.Call(ctx, "GET", u, c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) Update(ctx context.Context, workId string, params *trimmer.WorkspaceParams) (*trimmer.Workspace, error) {
	if workId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	u := fmt.Sprintf("/workspaces/%v", workId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Workspace{}
	err := c.B.Call(ctx, "PATCH", u, c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) Delete(ctx context.Context, workId string) error {
	if workId == "" {
		return trimmer.EIDMissing
	}
	return c.B.Call(ctx, "DELETE", fmt.Sprintf("/workspaces/%v", workId), c.Key, c.Sess, nil, nil, nil)
}

func (c Client) ListAssets(ctx context.Context, workId string, params *trimmer.AssetListParams) *asset.Iter {
	if workId == "" {
		return &asset.Iter{trimmer.GetIterErr(trimmer.EIDMissing)}
	}

	type assetList struct {
		trimmer.ListMeta
		Values trimmer.AssetList `json:"assets"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

		if len(params.IDs) > 0 {
			q.Add("id", strings.Join(params.IDs, ","))
		}
		if !trimmer.IsNilUUID(params.UUID) {
			q.Add("uuid", params.UUID)
		}
		if params.AccessClass != "" {
			q.Add("access", string(params.AccessClass))
		}
		if params.State != "" {
			q.Add("state", string(params.State))
		}
		if params.Version != "" {
			q.Add("version", string(params.Version))
		}
		if params.Original {
			q.Add("original", "true")
		}
		if params.Head {
			q.Add("head", "true")
		}
		if params.Event != "" {
			q.Add("event", string(params.Event))
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &asset.Iter{trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &assetList{}
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/workspaces/%v/assets?%v", workId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) NewAsset(ctx context.Context, workId string, params *trimmer.AssetParams) (*trimmer.Asset, error) {
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	if workId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/workspaces/%v/assets", workId)
	if params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Asset{}
	err := c.B.Call(ctx, "POST", u, c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) ListStashes(ctx context.Context, workId string, params *trimmer.StashListParams) *stash.Iter {

	if workId == "" {
		return &stash.Iter{trimmer.GetIterErr(trimmer.EIDMissing)}
	}

	type stashList struct {
		trimmer.ListMeta
		Values trimmer.StashList `json:"stashes"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

		if params.Name != "" {
			q.Add("name", params.Name)
		}
		if params.Type != "" {
			q.Add("type", string(params.Type))
		}
		if params.AccessClass != "" {
			q.Add("access", string(params.AccessClass))
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &stash.Iter{trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &stashList{}
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/workspaces/%v/stashes?%v", workId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}

}

func (c Client) NewStash(ctx context.Context, workId string, params *trimmer.StashParams) (*trimmer.Stash, error) {
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	if workId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/workspaces/%v/stashes", workId)
	if params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Stash{}
	err := c.B.Call(ctx, "POST", u, c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) ListMedia(ctx context.Context, workId string, params *trimmer.MediaListParams) *media.Iter {
	if workId == "" {
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
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/workspaces/%v/media?%v", workId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) ListEvents(ctx context.Context, workId string, params *trimmer.EventListParams) *event.Iter {
	if workId == "" {
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
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/workspaces/%v/events?%v", workId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) ListProfiles(ctx context.Context, workId string, params *trimmer.ProfileListParams) *profile.Iter {
	if workId == "" {
		return &profile.Iter{Iter: trimmer.GetIterErr(trimmer.EIDMissing)}
	}

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
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/workspaces/%v/profiles?%s", workId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) ListJobs(ctx context.Context, workId string, params *trimmer.JobListParams) *job.Iter {
	if workId == "" {
		return &job.Iter{trimmer.GetIterErr(trimmer.EIDMissing)}
	}

	type jobList struct {
		trimmer.ListMeta
		Values trimmer.JobList `json:"jobs"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

		if params.State != "" {
			q.Add("state", string(params.State))
		}
		if params.Type != "" {
			q.Add("type", string(params.Type))
		}
		if params.Queue != "" {
			q.Add("queue", string(params.Queue))
		}
		if params.AuthorId != "" {
			q.Add("authorId", params.AuthorId)
		}
		if params.MediaId != "" {
			q.Add("mediaId", params.MediaId)
		}
		if params.VolumeId != "" {
			q.Add("volumeId", params.VolumeId)
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &job.Iter{trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &jobList{}
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/workspaces/%v/jobs?%v", workId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) ListMounts(ctx context.Context, workId string, params *trimmer.MountListParams) *mount.Iter {
	if workId == "" {
		return &mount.Iter{Iter: trimmer.GetIterErr(trimmer.EIDMissing)}
	}

	type mountList struct {
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
		if len(params.Types) > 0 {
			q.Add("type", params.Types.String())
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

	return &mount.Iter{Iter: trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &mountList{}
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/workspaces/%v/mounts?%v", workId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) MountVolume(ctx context.Context, workId, volId string, params *trimmer.MountParams) error {
	if workId == "" || volId == "" {
		return trimmer.EIDMissing
	}
	u := fmt.Sprintf("/workspaces/%v/volumes/%v", workId, volId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	return c.B.Call(ctx, "PUT", u, c.Key, c.Sess, nil, params, nil)
}

func (c Client) UnmountVolume(ctx context.Context, workId, volId string) error {
	if workId == "" || volId == "" {
		return trimmer.EIDMissing
	}
	u := fmt.Sprintf("/workspaces/%v/volumes/%v", workId, volId)
	return c.B.Call(ctx, "DELETE", u, c.Key, c.Sess, nil, nil, nil)
}

func (c Client) ScanVolume(ctx context.Context, workId, volId string, params *trimmer.VolumeScanParams) (*trimmer.Job, error) {
	if workId == "" || volId == "" {
		return nil, trimmer.EIDMissing
	}
	v := &trimmer.Job{}
	u := fmt.Sprintf("/workspaces/%v/volumes/%v/scan", workId, volId)
	err := c.B.Call(ctx, "POST", u, c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) ClearVolume(ctx context.Context, workId, volId string, params *trimmer.VolumeClearParams) (*trimmer.Job, error) {
	if workId == "" || volId == "" {
		return nil, trimmer.EIDMissing
	}
	v := &trimmer.Job{}
	u := fmt.Sprintf("/workspaces/%v/volumes/%v/clear", workId, volId)
	err := c.B.Call(ctx, "POST", u, c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) GetMember(ctx context.Context, workId, userId string) (*trimmer.Member, error) {
	if workId == "" || userId == "" {
		return nil, trimmer.EIDMissing
	}
	v := &trimmer.Member{}
	err := c.B.Call(ctx, "GET", fmt.Sprintf("/workspaces/%v/members/%v", workId, userId), c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) NewMember(ctx context.Context, workId, userId string, params *trimmer.MemberParams) (*trimmer.Member, error) {
	if workId == "" || userId == "" {
		return nil, trimmer.EIDMissing
	}
	v := &trimmer.Member{}
	err := c.B.Call(ctx, "PUT", fmt.Sprintf("/workspaces/%v/members/%v", workId, userId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) DeleteMember(ctx context.Context, workId, userId string) error {
	if workId == "" || userId == "" {
		return trimmer.EIDMissing
	}
	return c.B.Call(ctx, "DELETE", fmt.Sprintf("/workspaces/%v/members/%v", workId, userId), c.Key, c.Sess, nil, nil, nil)
}

func (c Client) ListMembers(ctx context.Context, workId string, params *trimmer.MemberListParams) *member.Iter {
	if workId == "" {
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
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/workspaces/%v/members?%v", workId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}
