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

package trimmer

import (
	"os"
	"strings"
	"sync"
	"time"
)

// ApiAccessScope is the list of allowed values for login session scopes
// allowed values are "public", "private", "upload", "publish", "admin",
// "callback"
type ApiAccessScope string

const (
	API_SCOPE_PUBLIC   = "public"
	API_SCOPE_PRIVATE  = "private"
	API_SCOPE_UPLOAD   = "upload"
	API_SCOPE_PUBLISH  = "publish"
	API_SCOPE_ADMIN    = "admin"
	API_SCOPE_CALLBACK = "callback"
)

// ApiScopes is a wrapper function used for concatenating multiple
// scopes for requests.
func ApiScopes(scopes ...ApiAccessScope) string {
	if len(scopes) == 0 {
		return ""
	}
	ss := make([]string, len(scopes))
	for i, v := range scopes {
		ss[i] = string(v)
	}
	return strings.Join(ss, ", ")
}

// ---------------------------------------------------------------------------
// Environment Variables
//
const (
	TRIMMER_USERNAME_KEY     = "TRIMMER_API_USERNAME"
	TRIMMER_PASSWORD_KEY     = "TRIMMER_API_PASSWORD"
	TRIMMER_CLIENT_TOKEN_KEY = "TRIMMER_CLIENT_TOKEN"
)

// global session, safe to use from goroutines
var (
	LoginSession Session
)

type ClientToken string

type LoginParams struct {
	Username string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Scopes   string `json:"scopes"`
}

func (p LoginParams) IsValid() bool {
	return (p.Username != "" || p.Email != "") && p.Password != ""
}

func (p *LoginParams) Reset() {
	p.Username = ""
	p.Email = ""
	p.Password = ""
	p.Scopes = ""
}

// Global Login Session State
type Session struct {
	TokenId      string    `json:"tokenId"`
	AccessToken  string    `json:"accessToken"`
	TokenType    string    `json:"tokenType"`
	Scopes       string    `json:"scopes"`
	RefreshToken string    `json:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
	User         *User     `json:"user"`
	mu           sync.RWMutex
}

func NewSession() *Session {
	LoginSession.Update(&Session{})
	return &LoginSession
}

func NewClientSession(token ClientToken) (*Session, error) {
	if token == "" {
		token = ClientToken(os.Getenv(TRIMMER_CLIENT_TOKEN_KEY))
	}
	if token == "" {
		return nil, EParamMissing
	}
	LoginSession.Update(&Session{
		AccessToken: string(token),
		TokenType:   "Bearer",
	})
	return &LoginSession, nil
}

func (s *Session) Update(ss *Session) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.TokenId = ss.TokenId
	s.AccessToken = ss.AccessToken
	s.TokenType = ss.TokenType
	s.Scopes = ss.Scopes
	s.RefreshToken = ss.RefreshToken
	s.ExpiresAt = ss.ExpiresAt
	s.User = ss.User
}

func (s *Session) GetAuthorization() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.TokenType + " " + s.AccessToken
}

func (s *Session) GetRefreshToken() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.RefreshToken
}

func (s *Session) IsValid() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.AccessToken != "" && !s.isExpired()
}

func (s *Session) IsExpired() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.isExpired()
}

func (s *Session) isExpired() bool {
	return !s.ExpiresAt.IsZero() && s.ExpiresAt.Before(time.Now().UTC())
}

func (s *Session) ValidFor() time.Duration {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.ExpiresAt.IsZero() {
		return 0
	} else {
		return s.ExpiresAt.Sub(time.Now().UTC())
	}
}

func (s *Session) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.TokenId = ""
	s.AccessToken = ""
	s.TokenType = ""
	s.Scopes = ""
	s.RefreshToken = ""
	s.ExpiresAt = time.Time{}
	s.User = nil
}
