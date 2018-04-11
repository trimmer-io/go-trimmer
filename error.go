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

package trimmer

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type Error interface {
	error
	IsUsage() bool
	IsInternal() bool
	IsApi() bool
	Cause() error
}

type errorType int

const (
	usageError = errorType(1 << iota)
	internalError
	apiError
)

type TrimmerError struct {
	typ         errorType
	RequestId   string `json:"requestId,omitempty"`
	SessionId   string `json:"sessionId,omitempty"`
	OauthScopes string `json:"oauthScopes,omitempty"`
	ErrorCode   int    `json:"code,omitempty"`
	StatusCode  int    `json:"status,omitempty"`
	Message     string `json:"message,omitempty"`
	Scope       string `json:"scope,omitempty"`
	Detail      string `json:"detail,omitempty"`
	Cause       error  `json:"cause,omitempty"`
}

func (e TrimmerError) IsUsage() bool    { return (e.typ & usageError) != 0 }
func (e TrimmerError) IsInternal() bool { return (e.typ & internalError) != 0 }
func (e TrimmerError) IsApi() bool      { return (e.typ & apiError) != 0 }

type apiErrorList []TrimmerError

type apiErrorResponse struct {
	Errors apiErrorList `json:"errors"`
}

func NewApiError(status int) TrimmerError {
	return TrimmerError{
		typ:        apiError,
		StatusCode: status,
	}
}

func NewUsageError(msg string, err error) TrimmerError {
	return TrimmerError{
		typ:     usageError,
		Message: msg,
		Cause:   err,
	}
}

func NewInternalError(msg string, err error) TrimmerError {
	return TrimmerError{
		typ:     internalError,
		Message: msg,
		Cause:   err,
	}
}

func quote(s string) string {
	return strings.Join([]string{"\"", s, "\""}, "")
}

func (e TrimmerError) Error() string {
	s := []string{}
	if e.RequestId != "" {
		s = append(s, strings.Join([]string{"request_id", e.RequestId}, "="))
	}
	if e.SessionId != "" {
		s = append(s, strings.Join([]string{"session_id", e.SessionId}, "="))
	}
	if e.OauthScopes != "" {
		s = append(s, strings.Join([]string{"oauth_scopes", quote(e.OauthScopes)}, "="))
	}
	if e.StatusCode != 0 {
		s = append(s, strings.Join([]string{"status", strconv.Itoa(e.StatusCode)}, "="))
	}
	if e.ErrorCode != 0 {
		s = append(s, strings.Join([]string{"code", strconv.Itoa(e.ErrorCode)}, "="))
	}
	if e.Scope != "" {
		s = append(s, strings.Join([]string{"scope", e.Scope}, "="))
	}
	s = append(s, strings.Join([]string{"message", quote(e.Message)}, "="))
	if e.Detail != "" {
		s = append(s, strings.Join([]string{"detail", quote(e.Detail)}, "="))
	}
	if e.Cause != nil {
		s = append(s, strings.Join([]string{"cause", quote(e.Cause.Error())}, "="))
	}
	return strings.Join(s, " ")
}

func ParseApiError(i io.Reader) TrimmerError {
	var response apiErrorResponse
	jsonDecoder := json.NewDecoder(i)
	if err := jsonDecoder.Decode(&response); err != nil {
		return NewInternalError("parsing API error failed", err)
	}
	e := response.Errors[0]
	e.typ = apiError
	return e
}

func ParseApiErrorFromByte(b []byte) TrimmerError {
	var response apiErrorResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return NewInternalError("parsing API error failed", err)
	}
	e := response.Errors[0]
	e.typ = apiError
	return e
}

func (e TrimmerError) MarshalIndent() []byte {
	errResp := apiErrorResponse{
		Errors: []TrimmerError{e},
	}
	b, _ := json.MarshalIndent(errResp, "", "  ")
	return b
}

func (e TrimmerError) Marshal() []byte {
	errResp := apiErrorResponse{
		Errors: []TrimmerError{e},
	}
	b, _ := json.Marshal(errResp)
	return b
}

var (
	ENilPointer   = TrimmerError{typ: usageError, Message: "unexpected nil pointer"}
	EIDMissing    = TrimmerError{typ: usageError, Message: "id value missing"}
	EParamMissing = TrimmerError{typ: usageError, Message: "missing parameter"}
	EParamInvalid = TrimmerError{typ: usageError, Message: "invalid parameter"}
)
