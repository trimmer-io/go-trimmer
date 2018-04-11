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

// Package job provides the /jobs APIs
package job

import (
	"context"
	"fmt"
	"net/url"

	trimmer "trimmer.io/go-trimmer"
)

// Iter is an iterator for lists of Jobs.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// JOb returns the most recent Job visited by a call to Next.
func (i *Iter) Job() *trimmer.Job {
	return i.Current().(*trimmer.Job)
}

// Client is used to invoke /users APIs.
type Client struct {
	B    trimmer.Backend
	CDN  trimmer.Backend
	Key  trimmer.ApiKey
	Sess *trimmer.Session
}

func getC() Client {
	return Client{trimmer.GetBackend(trimmer.APIBackend), trimmer.GetBackend(trimmer.CDNBackend), trimmer.Key, &trimmer.LoginSession}
}

func Get(ctx context.Context, jobId string, params *trimmer.JobParams) (*trimmer.Job, error) {
	return getC().Get(ctx, jobId, params)
}

func Update(ctx context.Context, jobId string, params *trimmer.JobParams) (*trimmer.Job, error) {
	return getC().Update(ctx, jobId, params)
}

func Cancel(ctx context.Context, jobId string) (*trimmer.Job, error) {
	return getC().Cancel(ctx, jobId)
}

func (c Client) Get(ctx context.Context, jobId string, params *trimmer.JobParams) (*trimmer.Job, error) {
	if jobId == "" {
		return nil, trimmer.EIDMissing
	}
	u := fmt.Sprintf("/jobs/%v", jobId)
	if params != nil && params.Embed.IsValid() {
		q := &url.Values{}
		q.Add("embed", params.Embed.String())
		u += fmt.Sprintf("?%v", q.Encode())
	}
	v := &trimmer.Job{}
	err := c.B.Call(ctx, "GET", u, c.Key, c.Sess, nil, nil, v)
	return v, err
}

func (c Client) Update(ctx context.Context, jobId string, params *trimmer.JobParams) (*trimmer.Job, error) {
	if jobId == "" {
		return nil, trimmer.EIDMissing
	}
	if params == nil {
		return nil, trimmer.ENilPointer
	}
	v := &trimmer.Job{}
	err := c.B.Call(ctx, "PATCH", fmt.Sprintf("/jobs/%v", jobId), c.Key, c.Sess, nil, params, v)
	return v, err
}

func (c Client) Cancel(ctx context.Context, jobId string) (*trimmer.Job, error) {
	params := &trimmer.JobParams{
		State: JobStateAborted,
	}
	return c.Update(ctx, jobId, params)
}
