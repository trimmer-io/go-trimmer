// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

// Package stash provides the /stashes APIs
package stash

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	trimmer "trimmer.io/go-trimmer"
	"trimmer.io/go-trimmer/link"
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

// Iter is an iterator for lists of Stashs.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// Stash returns the most recent User visited by a call to Next.
func (i *Iter) Stash() *trimmer.Stash {
	return i.Current().(*trimmer.Stash)
}

func Get(ctx context.Context, stashId string, params *trimmer.StashParams) (*trimmer.Stash, error) {
	return getC().Get(ctx, stashId, params)
}

func Update(ctx context.Context, stashId string, params *trimmer.StashParams) (*trimmer.Stash, error) {
	return getC().Update(ctx, stashId, params)
}

func Delete(ctx context.Context, stashId string) error {
	return getC().Delete(ctx, stashId)
}

func Watch(ctx context.Context, stashId string) error {
	return getC().Watch(ctx, stashId)
}

func Unwatch(ctx context.Context, stashId string) error {
	return getC().Unwatch(ctx, stashId)
}

func IsWatching(ctx context.Context, stashId string) (bool, error) {
	return getC().IsWatching(ctx, stashId)
}

func ListLinks(ctx context.Context, stashId string, params *trimmer.LinkListParams) *link.Iter {
	return getC().ListLinks(ctx, stashId, params)
}

func NewLink(ctx context.Context, stashId string, params *trimmer.LinkParams) (*trimmer.Link, error) {
	return getC().NewLink(ctx, stashId, params)
}

func GetLink(ctx context.Context, stashId, linkId string, params *trimmer.LinkParams) (*trimmer.Link, error) {
	return getC().GetLink(ctx, stashId, linkId, params)
}

func UpdateLink(ctx context.Context, stashId, linkId string, params *trimmer.LinkParams) (*trimmer.Link, error) {
	return getC().UpdateLink(ctx, stashId, linkId, params)
}

func DeleteLink(ctx context.Context, stashId, linkId string) error {
	return getC().DeleteLink(ctx, stashId, linkId)
}

func ClearLinks(ctx context.Context, stashId string) error {
	return getC().ClearLinks(ctx, stashId)
}

func (c Client) Get(ctx context.Context, stashId string, params *trimmer.StashParams) (*trimmer.Stash, error) {
	if stashId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/stashes/%v", stashId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Stash{}
	err := c.B.Call(ctx, "GET", u, c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) Update(ctx context.Context, stashId string, params *trimmer.StashParams) (*trimmer.Stash, error) {
	if stashId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.Stash{}
	err := c.B.Call(ctx, "PATCH", fmt.Sprintf("/stashes/%v", stashId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) Delete(ctx context.Context, stashId string) error {
	if stashId == "" {
		return trimmer.EIDMissing
	}
	return c.B.Call(ctx, "DELETE", fmt.Sprintf("/stashes/%v", stashId), c.Key, c.Sess, nil, nil, nil)
}

func (c Client) Watch(ctx context.Context, stashId string) error {
	if stashId == "" {
		return trimmer.EIDMissing
	}
	return c.B.Call(ctx, "PUT", fmt.Sprintf("/stashes/%v/watch", stashId), c.Key, c.Sess, nil, nil, nil)
}

func (c Client) Unwatch(ctx context.Context, stashId string) error {
	if stashId == "" {
		return trimmer.EIDMissing
	}
	return c.B.Call(ctx, "DELETE", fmt.Sprintf("/stashes/%v/watch", stashId), c.Key, c.Sess, nil, nil, nil)
}

func (c Client) IsWatching(ctx context.Context, stashId string) (bool, error) {
	if stashId == "" {
		return false, trimmer.EIDMissing
	}
	err := c.B.Call(ctx, "GET", fmt.Sprintf("/stashes/%v/watch", stashId), c.Key, c.Sess, nil, nil, nil)
	switch e := err.(type) {
	case trimmer.TrimmerError:
		switch e.StatusCode {
		case http.StatusNotFound:
			return false, nil
		case http.StatusNoContent:
			return true, nil
		default:
			return false, err
		}
	default:
		return false, err
	}
}

func (c Client) ListLinks(ctx context.Context, stashId string, params *trimmer.LinkListParams) *link.Iter {

	if stashId == "" {
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

		if params.AssetId != "" {
			q.Add("assetId", string(params.AssetId))
		}
		if params.AuthorId != "" {
			q.Add("authorId", string(params.AuthorId))
		}
		if params.Embed.IsValid() {
			q.Add("embed", params.Embed.String())
		}

		params.AppendTo(q)
		lp = &params.ListParams
	}

	return &link.Iter{trimmer.GetIter(lp, q, func(b url.Values) ([]interface{}, trimmer.ListMeta, error) {
		list := &linkList{}
		err := c.B.Call(ctx, "GET", fmt.Sprintf("/stashes/%v/links?%v", stashId, b.Encode()), c.Key, c.Sess, nil, nil, list)
		ret := make([]interface{}, len(list.Values))

		// pass concrete values as abstract interface into iterator
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) NewLink(ctx context.Context, stashId string, params *trimmer.LinkParams) (*trimmer.Link, error) {
	if stashId == "" {
		return nil, trimmer.EIDMissing
	}
	v := &trimmer.Link{}
	err := c.B.Call(ctx, "POST", fmt.Sprintf("/stashes/%v/links", stashId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) GetLink(ctx context.Context, stashId, linkId string, params *trimmer.LinkParams) (*trimmer.Link, error) {
	if stashId == "" || linkId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/stashes/%v/links/%v", stashId, linkId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Link{}
	err := c.B.Call(ctx, "GET", u, c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) UpdateLink(ctx context.Context, stashId, linkId string, params *trimmer.LinkParams) (*trimmer.Link, error) {
	if stashId == "" || linkId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.Link{}
	err := c.B.Call(ctx, "PATCH", fmt.Sprintf("/stashes/%v/links/%v", stashId, linkId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) DeleteLink(ctx context.Context, stashId, linkId string) error {
	if stashId == "" || linkId == "" {
		return trimmer.EIDMissing
	}
	return c.B.Call(ctx, "DELETE", fmt.Sprintf("/stashes/%v/links/%v", stashId, linkId), c.Key, c.Sess, nil, nil, nil)
}

func (c Client) ClearLinks(ctx context.Context, stashId string) error {
	if stashId == "" {
		return trimmer.EIDMissing
	}
	return c.B.Call(ctx, "DELETE", fmt.Sprintf("/stashes/%v/links", stashId), c.Key, c.Sess, nil, nil, nil)
}
