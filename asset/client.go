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

// Package asset provides the /assets APIs
package asset

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	trimmer "trimmer.io/go-trimmer"
	"trimmer.io/go-trimmer/link"
	"trimmer.io/go-trimmer/media"
	"trimmer.io/go-trimmer/meta"
	"trimmer.io/go-trimmer/tag"
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

// Iter is an iterator for lists of Assets.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// Asset returns the most recent Asset visited by a call to Next.
func (i *Iter) Asset() *trimmer.Asset {
	return i.Current().(*trimmer.Asset)
}

func Get(ctx context.Context, assetId string, params *trimmer.AssetParams) (*trimmer.Asset, error) {
	return getC().Get(ctx, assetId, params)
}

func Update(ctx context.Context, assetId string, params *trimmer.AssetUpdateParams) (*trimmer.Asset, error) {
	return getC().Update(ctx, assetId, params)
}

func Delete(ctx context.Context, assetId string) error {
	return getC().Delete(ctx, assetId)
}

func ForkCopy(ctx context.Context, assetId string, params *trimmer.AssetForkParams) (*trimmer.Asset, error) {
	return getC().ForkCopy(ctx, assetId, params)
}

func ForkVersion(ctx context.Context, assetId string, params *trimmer.AssetForkParams) (*trimmer.Asset, error) {
	return getC().ForkVersion(ctx, assetId, params)
}

func ListVersions(ctx context.Context, assetId string, params *trimmer.AssetListParams) *Iter {
	return getC().ListVersions(ctx, assetId, params)
}

func DeleteVersions(ctx context.Context, assetId string) error {
	return getC().DeleteVersions(ctx, assetId)
}

func Trash(ctx context.Context, assetId string) (*trimmer.Asset, error) {
	return getC().Trash(ctx, assetId)
}

func Undelete(ctx context.Context, assetId string) (*trimmer.Asset, error) {
	return getC().Undelete(ctx, assetId)
}

func GetRevision(ctx context.Context, assetId string, params *trimmer.MetaQueryParams) (*trimmer.MetaRevision, error) {
	return getC().GetRevision(ctx, assetId, params)
}

func DiffRevisions(ctx context.Context, assetId string, params *trimmer.MetaDiffParams) ([]byte, error) {
	return getC().DiffRevisions(ctx, assetId, params)
}

func ListRevisions(ctx context.Context, assetId string, params *trimmer.MetaListParams) *meta.Iter {
	return getC().ListRevisions(ctx, assetId, params)
}

func CommitRevision(ctx context.Context, assetId string, params *trimmer.MetaUpdateParams) (*trimmer.MetaRevision, error) {
	return getC().CommitRevision(ctx, assetId, params)
}

func ListLinks(ctx context.Context, assetId string, params *trimmer.LinkListParams) *link.Iter {
	return getC().ListLinks(ctx, assetId, params)
}

func ListTags(ctx context.Context, assetId string, params *trimmer.TagListParams) *tag.Iter {
	return getC().ListTags(ctx, assetId, params)
}

func NewTag(ctx context.Context, assetId string, params *trimmer.TagParams) (*trimmer.Tag, error) {
	return getC().NewTag(ctx, assetId, params)
}

func ListMedia(ctx context.Context, assetId string, params *trimmer.MediaListParams) *media.Iter {
	return getC().ListMedia(ctx, assetId, params)
}

func NewMedia(ctx context.Context, assetId string, params *trimmer.MediaParams) (*trimmer.Media, error) {
	return getC().NewMedia(ctx, assetId, params)
}

func NewUpload(ctx context.Context, assetId string, params *trimmer.MediaParams) (*trimmer.Media, error) {
	return getC().NewUpload(ctx, assetId, params)
}

func UploadMedia(ctx context.Context, assetId string, params *trimmer.MediaParams, src io.ReadSeeker) (*trimmer.Media, error) {
	return getC().UploadMedia(ctx, assetId, params, src)
}

func DeleteMedia(ctx context.Context, assetId, mediaId string) error {
	return getC().DeleteMedia(ctx, assetId, mediaId)
}

func Analyze(ctx context.Context, assetId string, params *trimmer.AssetAnalyzeParams) (*trimmer.Job, error) {
	return getC().Analyze(ctx, assetId, params)
}

func Snapshot(ctx context.Context, assetId string, params *trimmer.AssetSnapshotParams) (*trimmer.Job, error) {
	return getC().Snapshot(ctx, assetId, params)
}

func Transcode(ctx context.Context, assetId string, params *trimmer.AssetTranscodeParams) (*trimmer.Job, error) {
	return getC().Transcode(ctx, assetId, params)
}

func Trim(ctx context.Context, assetId string, params *trimmer.AssetTrimParams) (*trimmer.Media, error) {
	return getC().Trim(ctx, assetId, params)
}

func Count(ctx context.Context, assetId string, params *trimmer.AssetCountParams) (*trimmer.Asset, error) {
	return getC().Count(ctx, assetId, params)
}

func Lock(ctx context.Context, assetId string, params *trimmer.EmbedParams) (*trimmer.Asset, error) {
	return getC().Lock(ctx, assetId, params)
}

func Unlock(ctx context.Context, assetId string, params *trimmer.EmbedParams) (*trimmer.Asset, error) {
	return getC().Unlock(ctx, assetId, params)
}

func (c Client) Get(ctx context.Context, assetId string, params *trimmer.AssetParams) (*trimmer.Asset, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/assets/%v", assetId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Asset{}
	err := c.B.Call(ctx, http.MethodGet, u, c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) Update(ctx context.Context, assetId string, params *trimmer.AssetUpdateParams) (*trimmer.Asset, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.Asset{}
	u := fmt.Sprintf("/assets/%v", assetId)
	err := c.B.Call(ctx, http.MethodPatch, u, c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) Delete(ctx context.Context, assetId string) error {
	if assetId == "" {
		return trimmer.EIDMissing
	}
	err := c.B.Call(ctx, http.MethodDelete, fmt.Sprintf("/assets/%v", assetId), c.Key, c.Sess, nil, nil, nil)
	return err
}

func (c Client) Trash(ctx context.Context, assetId string) (*trimmer.Asset, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	v := &trimmer.Asset{}
	err := c.B.Call(ctx, http.MethodPost, fmt.Sprintf("/assets/%v/trash", assetId), c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) Undelete(ctx context.Context, assetId string) (*trimmer.Asset, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	v := &trimmer.Asset{}
	err := c.B.Call(ctx, http.MethodPost, fmt.Sprintf("/assets/%v/undelete", assetId), c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) ForkCopy(ctx context.Context, assetId string, params *trimmer.AssetForkParams) (*trimmer.Asset, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.Asset{}
	err := c.B.Call(ctx, http.MethodPost, fmt.Sprintf("/assets/%v/fork", assetId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) ForkVersion(ctx context.Context, assetId string, params *trimmer.AssetForkParams) (*trimmer.Asset, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.Asset{}
	err := c.B.Call(ctx, http.MethodPost, fmt.Sprintf("/assets/%v/versions", assetId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) ListVersions(ctx context.Context, assetId string, params *trimmer.AssetListParams) *Iter {
	if assetId == "" {
		return &Iter{trimmer.GetIterErr(trimmer.EIDMissing)}
	}
	type assetList struct {
		trimmer.ListMeta
		Values trimmer.AssetList `json:"versions"`
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
		list := &assetList{}
		err := c.B.Call(ctx, http.MethodGet, fmt.Sprintf("/assets/%v/versions?%v", assetId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) DeleteVersions(ctx context.Context, assetId string) error {
	if assetId == "" {
		return trimmer.EIDMissing
	}
	return c.B.Call(ctx, http.MethodDelete, fmt.Sprintf("/assets/%v/versions", assetId), c.Key, c.Sess, nil, nil, nil)
}
func (c Client) ListLinks(ctx context.Context, assetId string, params *trimmer.LinkListParams) *link.Iter {
	if assetId == "" {
		return &link.Iter{trimmer.GetIterErr(trimmer.EIDMissing)}
	}

	type linkList struct {
		trimmer.ListMeta
		Values trimmer.LinkList `json:"links"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

		if params.StashId != "" {
			q.Add("stashId", params.StashId)
		}
		if params.AuthorId != "" {
			q.Add("authorId", params.AuthorId)
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &link.Iter{trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &linkList{}
		err := c.B.Call(ctx, http.MethodGet, fmt.Sprintf("/assets/%v/links?%v", assetId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) ListTags(ctx context.Context, assetId string, params *trimmer.TagListParams) *tag.Iter {
	if assetId == "" {
		return &tag.Iter{trimmer.GetIterErr(trimmer.EIDMissing)}
	}

	type tagList struct {
		trimmer.ListMeta
		Values trimmer.TagList `json:"tags"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

		if len(params.IDs) > 0 {
			q.Add("id", strings.Join(params.IDs, ","))
		}
		if params.AuthorId != "" {
			q.Add("authorId", params.AuthorId)
		}
		if params.AccessClass != "" {
			q.Add("access", string(params.AccessClass))
		}
		if len(params.Labels) > 0 {
			q.Add("label", params.Labels.String())
		}
		if params.From > 0 {
			q.Add("from", strconv.FormatInt(params.From, 10))
		}
		if params.To > 0 {
			q.Add("to", strconv.FormatInt(params.To, 10))
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &tag.Iter{trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &tagList{}
		err := c.B.Call(ctx, http.MethodGet, fmt.Sprintf("/assets/%v/tags?%v", assetId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) NewTag(ctx context.Context, assetId string, params *trimmer.TagParams) (*trimmer.Tag, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.Tag{}
	err := c.B.Call(ctx, http.MethodPost, fmt.Sprintf("/assets/%v/tags", assetId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) ListMedia(ctx context.Context, assetId string, params *trimmer.MediaListParams) *media.Iter {
	if assetId == "" {
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
		if params.Kind != "" {
			q.Add("kind", string(params.Kind))
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
		if !trimmer.IsNilUUID(params.UUID) {
			q.Add("uuid", params.UUID)
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &media.Iter{Iter: trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &mediaList{}
		err := c.B.Call(ctx, http.MethodGet, fmt.Sprintf("/assets/%v/media?%v", assetId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) NewMedia(ctx context.Context, assetId string, params *trimmer.MediaParams) (*trimmer.Media, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.Media{}
	media.StripMetadataUrls(params.Attr)
	err := c.B.Call(ctx, http.MethodPost, fmt.Sprintf("/assets/%v/media", assetId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) GetRevision(ctx context.Context, assetId string, params *trimmer.MetaQueryParams) (*trimmer.MetaRevision, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	var u string
	if params != nil {
		revision := "head"
		if params.Revision != "" {
			revision = params.Revision
		}
		q := &url.Values{}
		if params.Filter != "" {
			q.Add("filter", params.Filter)
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}
		if len(*q) > 0 {
			u = fmt.Sprintf("/assets/%v/meta/%v?%s", assetId, revision, q.Encode())
		} else {
			u = fmt.Sprintf("/assets/%v/meta/%v", assetId, revision)
		}
	} else {
		u = fmt.Sprintf("/assets/%v/meta/head", assetId)
	}
	v := &trimmer.MetaRevision{}
	err := c.B.Call(ctx, http.MethodGet, u, c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) ListRevisions(ctx context.Context, assetId string, params *trimmer.MetaListParams) *meta.Iter {
	if assetId == "" {
		return &meta.Iter{trimmer.GetIterErr(trimmer.EIDMissing)}
	}

	type metaList struct {
		trimmer.ListMeta
		Values trimmer.MetaRevisionList `json:"revisions"`
	}

	var q *url.Values
	var lp *trimmer.ListParams
	if params != nil {
		q = &url.Values{}

		if params.AuthorId != "" {
			q.Add("authorId", params.AuthorId)
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &meta.Iter{trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &metaList{}
		err := c.B.Call(ctx, http.MethodGet, fmt.Sprintf("/assets/%v/meta?%v", assetId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) DiffRevisions(ctx context.Context, assetId string, params *trimmer.MetaDiffParams) ([]byte, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	v1 := "head"
	v2 := "head^"
	if params != nil {
		if params.V1 != "" {
			v1 = params.V1
		}
		if params.V2 != "" {
			v2 = params.V2
		}
	}
	var buf bytes.Buffer
	h := &trimmer.CallHeaders{
		Accept: "application/vnd.trimmer.diff",
	}
	err := c.B.Call(ctx, http.MethodGet, fmt.Sprintf("/assets/%v/meta?diff=%s:%s", assetId, v1, v2), c.Key, c.Sess, h, nil, buf)
	return buf.Bytes(), err
}

func (c Client) CommitRevision(ctx context.Context, assetId string, params *trimmer.MetaUpdateParams) (*trimmer.MetaRevision, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.MetaRevision{}
	err := c.B.Call(ctx, http.MethodPatch, fmt.Sprintf("/assets/%v/meta", assetId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) NewUpload(ctx context.Context, assetId string, params *trimmer.MediaParams) (*trimmer.Media, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.Media{}
	media.StripMetadataUrls(params.Attr)
	err := c.B.Call(ctx, http.MethodPost, fmt.Sprintf("/assets/%v/upload", assetId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) UploadMedia(ctx context.Context, assetId string, params *trimmer.MediaParams, src io.ReadSeeker) (*trimmer.Media, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}

	// 1 create asset media using params
	m, err := c.NewUpload(ctx, assetId, params)
	if err != nil {
		return nil, err
	}

	fi := &trimmer.FileInfo{
		Size:     m.Size,
		Hashes:   m.Hashes,
		Etag:     m.Hashes.Etag(),
		Filename: m.Filename,
		UUID:     m.UUID,
		Mimetype: m.Mimetype,
		Url:      m.Url,
	}

	// 2 upload file data
	r := media.NewUploadRequest(fi, m, src)
	fi, err = r.Do(ctx)
	if err != nil {
		c.DeleteMedia(ctx, assetId, m.ID)
		return nil, err
	}

	// 3 complete asset upload (not required when callback is used)
	if m.State == media.MediaStateUploading && !r.HasCallback() {
		i := m.ID
		up := &trimmer.MediaUploadCompletionParams{
			Files: trimmer.FileInfoList{fi},
			Embed: trimmer.API_EMBED_META | trimmer.API_EMBED_DETAILS,
		}
		if m, err = media.CompleteUpload(ctx, m.ID, up); err != nil {
			c.DeleteMedia(ctx, assetId, i)
			return nil, err
		}
	}

	return m, err
}

func (c Client) DeleteMedia(ctx context.Context, assetId, mediaId string) error {
	if assetId == "" || mediaId == "" {
		return trimmer.EIDMissing
	}
	err := c.B.Call(ctx, http.MethodDelete, fmt.Sprintf("/assets/%v/media/%v", assetId, mediaId), c.Key, c.Sess, nil, nil, nil)
	return err
}

func (c Client) Analyze(ctx context.Context, assetId string, params *trimmer.AssetAnalyzeParams) (*trimmer.Job, error) {
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	if assetId == "" || params.MediaId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/assets/%v/media/%v/analyze", assetId, params.MediaId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Job{}
	err := c.B.Call(ctx, http.MethodPost, u, c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) Snapshot(ctx context.Context, assetId string, params *trimmer.AssetSnapshotParams) (*trimmer.Job, error) {
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	if assetId == "" || params.MediaId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/assets/%v/media/%v/snapshot", assetId, params.MediaId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Job{}
	err := c.B.Call(ctx, http.MethodPost, u, c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) Transcode(ctx context.Context, assetId string, params *trimmer.AssetTranscodeParams) (*trimmer.Job, error) {
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	v := &trimmer.Job{}
	err := c.B.Call(ctx, http.MethodPost, fmt.Sprintf("/assets/%v/transcode", assetId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) Trim(ctx context.Context, assetId string, params *trimmer.AssetTrimParams) (*trimmer.Media, error) {
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	if assetId == "" || params.MediaId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/assets/%v/media/%v", assetId, params.MediaId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Media{}
	err := c.B.Call(ctx, http.MethodPatch, u, c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) Count(ctx context.Context, assetId string, params *trimmer.AssetCountParams) (*trimmer.Asset, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	u := fmt.Sprintf("/assets/%v/counts", assetId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Asset{}
	err := c.B.Call(ctx, http.MethodPost, u, c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) Lock(ctx context.Context, assetId string, params *trimmer.EmbedParams) (*trimmer.Asset, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/assets/%v/lock", assetId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Asset{}
	err := c.B.Call(ctx, http.MethodPost, u, c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) Unlock(ctx context.Context, assetId string, params *trimmer.EmbedParams) (*trimmer.Asset, error) {
	if assetId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/assets/%v/lock", assetId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Asset{}
	err := c.B.Call(ctx, http.MethodDelete, u, c.Key, c.Sess, nil, nil, v)
	return v, err
}
