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

// Package session provides the /auth APIs
package session

import (
	"context"
	"os"

	trimmer "trimmer.io/go-trimmer"
)

// Client is used to invoke /auth APIs.
type Client struct {
	B    trimmer.Backend
	Key  trimmer.ApiKey
	Sess *trimmer.Session
}

func ParseEnv() *trimmer.LoginParams {
	p := &trimmer.LoginParams{
		Scopes: trimmer.ApiScopes(trimmer.API_SCOPE_PUBLIC, trimmer.API_SCOPE_PRIVATE, trimmer.API_SCOPE_UPLOAD),
	}
	if v := os.Getenv(trimmer.TRIMMER_USERNAME_KEY); v != "" {
		p.Username = v
	}
	if v := os.Getenv(trimmer.TRIMMER_PASSWORD_KEY); v != "" {
		p.Password = v
	}
	return p
}

func Check(ctx context.Context) error {
	return getC().Check(ctx)
}

func (c Client) Check(ctx context.Context) error {
	// GET /users/me
	v := &trimmer.User{}
	err := c.B.Call(ctx, "GET", "/users/me", c.Key, c.Sess, nil, nil, v)
	if err == nil {
		c.Sess.User = v
		// TODO: extract auth scopes from header
	}
	return err
}

func Refresh(ctx context.Context) error {
	return getC().Refresh(ctx)
}

func (c Client) Refresh(ctx context.Context) error {
	authRefresh := struct {
		RefreshToken string `json:"refreshToken"`
	}{
		RefreshToken: c.Sess.GetRefreshToken(),
	}

	s := &trimmer.Session{}
	err := c.B.Call(ctx, "POST", "/auth/refresh", c.Key, c.Sess, nil, authRefresh, s)
	if err != nil {
		return err
	}
	c.Sess.Update(s)
	return nil
}

func Login(ctx context.Context, params *trimmer.LoginParams) error {
	return getC().Login(ctx, params)
}

func (c Client) Login(ctx context.Context, params *trimmer.LoginParams) error {
	s := &trimmer.Session{}
	err := c.B.Call(ctx, "POST", "/auth/login", c.Key, nil, nil, params, s)
	if err == nil {
		c.Sess.Update(s)

		if trimmer.LogLevel > 1 {
			trimmer.Logger.Printf("INFO: Logged in as [%v] %v %v\n", c.Sess.User.ID, c.Sess.User.Name, c.Sess.User.DisplayName)
		}

	}
	return err
}

func Logout(ctx context.Context) error {
	return getC().Logout(ctx)
}

func (c Client) Logout(ctx context.Context) error {
	err := c.B.Call(ctx, "POST", "/auth/logout", c.Key, c.Sess, nil, nil, nil)
	if err == nil {
		if trimmer.LogLevel > 1 {
			trimmer.Logger.Printf("INFO: Logged out\n")
		}
		c.Sess.Reset()
	}
	return err
}

func getC() Client {
	return Client{trimmer.GetBackend(trimmer.APIBackend), trimmer.Key, &trimmer.LoginSession}
}
