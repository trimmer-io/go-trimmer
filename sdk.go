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

package trimmer // import "trimmer.io/go-trimmer"

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
	"trimmer.io/go-trimmer/hash"
)

// ---------------------------------------------------------------------------
// SDK version settings
//

// apiversion is the currently supported API version
const ApiVersion = "2018-08-04"

// clientversion is the binding version
const ClientVersion = "1.3"

// defaultHTTPTimeout is the default timeout on the http.Client used by the library.
const defaultHTTPTimeout = 15 * time.Second

// Totalbackends is the total number of Trimmer API endpoints supported by the binding.
const TotalBackends = 2

// MaxRetries is the maximum number of times the client tries issuing a call after
// transient errors (io, network or 5xx server errors), the first retry is performed
// immediately, subsequent retries will wait
const MaxRetries = 10

// RetryBackoffTime defines the time instant that is added to the wait time before
// a new retry attempt is made. The first retry is sent immediately, any subsequent
// retry will wait an additional RetryBackoffTime longer.
const RetryBackoffTime = time.Duration(5) * time.Second

// DefaultPartSize defines the minimum size in bytes for upload parts (= 16 MiB).
const DefaultPartSize = int64(16) << 20

// Backend is an interface for making calls against a Trimmer service.
// This interface exists to enable mocking for tests if needed.
type Backend interface {
	GetUrl() string
	Call(ctx context.Context, method, path string, key ApiKey, sess *Session, headers *CallHeaders, data, v interface{}) error
	CallMultipart(ctx context.Context, method, path string, key ApiKey, sess *Session, headers *CallHeaders, body io.Reader, v interface{}) error
	CallChecksum(ctx context.Context, method, path string, key ApiKey, sess *Session, headers *CallHeaders, flags hash.HashFlags, body io.Reader, resp io.Writer, v interface{}) (int64, hash.HashBlock, hash.HashBlock, error)
}

// BackendConfiguration is the internal implementation for making HTTP calls to Trimmer.
type BackendConfiguration struct {
	Type       SupportedBackend
	URL        string
	HTTPClient *http.Client
}

// SupportedBackend is an enumeration of supported Trimmer endpoints.
// Currently supported values are "api" and "cdn".
type SupportedBackend string

const (
	APIBackend SupportedBackend = "api"
	CDNBackend SupportedBackend = "cdn"
)

// Backends are the currently supported endpoints.
type Backends struct {
	API, CDN Backend
}

// Optional HTTP header data used as call in and out parameter
type CallHeaders struct {
	// in & out parameters
	ContentType        string         // Content-Type
	ContentDisposition string         // Content-Disposition
	Accept             string         // Accept
	Size               int64          // Content-Length
	Hashes             hash.HashBlock // Content-MD5, X-Trimmer-Hash

	// out only
	OAuthScopes string // X-OAuth-Scopes
	SessionId   string // X-Session-Id
	RequestId   string // X-Request-Id
	Runtime     string // X-Runtime
}

// Key is the Trimmer API key used globally in the binding.
const (
	TRIMMER_API_KEY_KEY    = "TRIMMER_API_KEY"
	TRIMMER_API_SERVER_KEY = "TRIMMER_API_SERVER"
	TRIMMER_CDN_SERVER_KEY = "TRIMMER_CDN_SERVER"
)

var (
	apiURL = "https://api.trimmer.io"
	cdnURL = "https://cdn.trimmer.io"
)

// user-settable API key
type ApiKey string

var Key ApiKey

// user-defined User agent string
var UserAgent string

// minimal size of upload parts
var UploadPartSize = DefaultPartSize

// LogLevel is the logging level for this library.
// 0: no logging
// 1: errors only
// 2: errors + informational (default)
// 3: errors + informational + debug
var LogLevel = 2

// Logger controls how the SDK performs logging at a package level. It is useful
// to customise if you need it prefixed for your application to meet other
// requirements
var Logger *log.Logger

func init() {
	// setup the logger
	Logger = log.New(os.Stderr, "", log.LstdFlags)

	// read API Key from environment
	if s := os.Getenv(TRIMMER_API_KEY_KEY); s != "" {
		Key = ApiKey(s)
	}

	// read API/CDN servers from environment
	if s := os.Getenv(TRIMMER_API_SERVER_KEY); s != "" {
		apiURL = s
	}

	if s := os.Getenv(TRIMMER_CDN_SERVER_KEY); s != "" {
		cdnURL = s
	}
}

var apiHttpClient = &http.Client{Timeout: defaultHTTPTimeout}
var cdnHttpClient = &http.Client{}
var backends Backends

// SetHTTPClient overrides the default HTTP client.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
func SetHTTPClient(backend SupportedBackend, client *http.Client) {
	switch backend {
	case APIBackend:
		apiHttpClient = client
	case CDNBackend:
		cdnHttpClient = client
	}
}

// NewBackends creates a new set of backends with the given HTTP client. You
// should only need to use this for testing purposes or on App Engine.
func NewBackends(httpClient *http.Client) *Backends {
	return &Backends{
		API: BackendConfiguration{
			APIBackend, apiURL, apiHttpClient},
		CDN: BackendConfiguration{
			CDNBackend, cdnURL, cdnHttpClient},
	}
}

// GetBackend returns the currently used backend in the binding.
func GetBackend(backend SupportedBackend) Backend {
	switch backend {
	case APIBackend:
		if backends.API == nil {
			backends.API = BackendConfiguration{backend, apiURL, apiHttpClient}
		}
		return backends.API
	case CDNBackend:
		if backends.CDN == nil {
			backends.CDN = BackendConfiguration{backend, cdnURL, cdnHttpClient}
		}
		return backends.CDN
	}

	return nil
}

// SetBackend sets the backend used in the binding.
func SetBackend(backend SupportedBackend, b Backend) {
	switch backend {
	case APIBackend:
		backends.API = b
	case CDNBackend:
		backends.CDN = b
	}
}

func (s BackendConfiguration) GetUrl() string {
	return s.URL
}

// Call is the Backend.Call implementation for invoking Trimmer APIs.
func (s BackendConfiguration) Call(ctx context.Context, method, path string, key ApiKey, sess *Session, headers *CallHeaders, data, v interface{}) error {

	// prepare POST/PUT/PATCH payload
	var body io.Reader
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(b)
	}

	if headers == nil {
		headers = &CallHeaders{}
	}

	if headers.ContentType == "" {
		headers.ContentType = "application/json"
	}

	if headers.Accept == "" {
		headers.Accept = "application/json"
	}

	req, err := s.NewRequest(method, path, key, sess, headers, body)
	if err != nil {
		return err
	}

	// reuse request headers as response header to return
	responseHeaders := headers

	if _, _, err := s.Do(ctx, req, sess, v, responseHeaders); err != nil {
		return err
	}

	return nil
}

func (s BackendConfiguration) CallMultipart(ctx context.Context, method, path string, key ApiKey, sess *Session, headers *CallHeaders, body io.Reader, v interface{}) error {

	if headers == nil {
		headers = &CallHeaders{}
	}

	if headers.ContentType == "" {
		headers.ContentType = "application/json"
	}

	if headers.Accept == "" {
		headers.Accept = "application/json"
	}

	req, err := s.NewRequest(method, path, key, sess, headers, body)
	if err != nil {
		return err
	}

	// reuse request headers as response header to return
	responseHeaders := headers

	if _, _, err = s.Do(ctx, req, sess, v, responseHeaders); err != nil {
		return err
	}
	return nil
}

func (s BackendConfiguration) CallChecksum(ctx context.Context, method, path string, key ApiKey, sess *Session, headers *CallHeaders, flags hash.HashFlags, r io.Reader, w io.Writer, v interface{}) (int64, hash.HashBlock, hash.HashBlock, error) {

	if headers == nil {
		headers = &CallHeaders{}
	}

	if headers.ContentType == "" && r != nil {
		headers.ContentType = "application/json"
	}

	if headers.Accept == "" && v != nil {
		headers.Accept = "application/json"
	}

	var (
		requestBody  io.Reader = r
		responseBody io.Writer = w
		clientHash   hash.HashBlock
		serverHash   hash.HashBlock
		size         int64
	)

	if flags > 0 {

		// Send only flagged hash values to the upstream server (clone() will strip
		// excluded values). The server will calculate and respond with the same set
		// of hashes.
		//
		headers.Hashes = headers.Hashes.Clone(flags)

		// on upload (when io.Reader is provided), we're going to hash outgoing data
		// on download (when io.Writer is provided), we're going to hash incoming data

		if r != nil {

			// Hash data while sending using the flagged hash functions. This adds
			// an additional layer of integrity protection (i.e. it is possible to
			// detect data corruption that has occured at the sender), but all used
			// hashes must be available BEFORE! sending the data.
			//
			// Note that for fresh data that is ingested into Trimmer for the first
			// time it may be necessary to calculate hash values before calling this
			// function if this has not happened yet.
			//
			requestBody = clientHash.NewReader(requestBody, flags)
		}

		if w != nil {
			// Hash data while receiving using the flagged hash functions. This adds
			// an additional layer of integrity protection (i.e. it is possible to
			// detect data corruption that occured during transit).
			//
			// Note that for validating hashes the original values must be available.
			//
			responseBody = clientHash.NewWriter(responseBody, flags)
		}
	}

	req, err := s.NewRequest(method, path, key, sess, headers, requestBody)
	if err != nil {
		return 0, hash.HashBlock{}, hash.HashBlock{}, err
	}

	// reuse request headers as response header to return
	responseHeaders := headers

	// on requests, the server response with his version of hash values
	if responseBody != nil {
		size, serverHash, err = s.Do(ctx, req, sess, responseBody, responseHeaders)
	} else {
		size, serverHash, err = s.Do(ctx, req, sess, v, responseHeaders)
	}
	if err != nil {
		return 0, hash.HashBlock{}, hash.HashBlock{}, err
	}

	// save hashes of sent/received data
	clientHash.Sum()

	// return local and remote hashes for what's been sent/received
	return size, clientHash, serverHash, nil
}

// NewRequest is used by Call to generate an http.Request. It adds appropriate headers.
func (s *BackendConfiguration) NewRequest(method, path string, key ApiKey, sess *Session, headers *CallHeaders, body io.Reader) (*http.Request, error) {

	// build full API/CDN URL unless path already starts with http
	if !strings.HasPrefix(path, "http") {
		path = s.URL + path
	}

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		if LogLevel > 0 {
			Logger.Printf("ERROR: cannot create request: %v\n", err)
		}
		return nil, NewUsageError("preparing request failed", err)
	}

	// add content-type header to POST, PUT, PATCH
	switch method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		req.Header.Add("Content-Type", headers.ContentType)
	}

	// add accept header
	if headers.Accept != "" {
		req.Header.Add("Accept", headers.Accept)
	}

	// add extra hash header
	if !headers.Hashes.IsZero() {
		req.Header.Add("X-Trimmer-Hash", headers.Hashes.String())

		// also add standard checksum header (HTTP supports only MD5 here)
		if etag := headers.Hashes.EtagHeader(); len(etag) > 2 {
			req.Header.Add("Content-MD5", etag)
		}
	}

	// API Key
	if key != "" {
		req.Header.Add("X-API-Key", string(key))
	} else {
		if LogLevel > 0 {
			Logger.Println("ERROR: API Key missing")
		}
		return nil, NewUsageError("API Key missing", nil)
	}

	// Authorization header
	if sess != nil && sess.IsValid() {
		req.Header.Add("Authorization", sess.GetAuthorization())
	}

	req.Header.Add("X-API-Version", ApiVersion)
	if UserAgent != "" {
		req.Header.Add("User-Agent", UserAgent)
	} else {
		req.Header.Add("User-Agent", "Trimmer-SDK-Go/"+ClientVersion)
	}

	return req, nil
}

// Do is used by Call to execute an API request and parse the response. It uses
// the backend's HTTP client to execute the request and unmarshals the response
// into v. It also handles unmarshaling errors returned by the API.
//
// This function also handles binary responses like downloading data.
//
func (s *BackendConfiguration) Do(ctx context.Context, req *http.Request, sess *Session, v interface{}, responseHeaders *CallHeaders) (int64, hash.HashBlock, error) {

	if LogLevel > 1 {
		q := req.URL.RawQuery
		if len(q) > 0 {
			Logger.Printf("%v %v%v?%v\n", req.Method, req.URL.Host, req.URL.Path, q)
		} else {
			Logger.Printf("%v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
		}

		if LogLevel > 2 {
			// only dump content-type application/json
			r, _ := httputil.DumpRequestOut(req, req.Header.Get("Content-Type") == "application/json")
			log.Println(string(r))
		}
	}

	start := time.Now()

	// create a new timeout child context
	var (
		callCtx context.Context
		cancel  context.CancelFunc
	)
	if s.HTTPClient.Timeout > 0 {
		callCtx, cancel = context.WithTimeout(ctx, s.HTTPClient.Timeout)
	} else {
		callCtx, cancel = context.WithCancel(ctx)
	}
	defer cancel()

	// wrap http request in context
	req = req.WithContext(callCtx)

	// calling the API, will fail on timeout
	resp, err := s.HTTPClient.Do(req)

	if err != nil {
		if LogLevel > 0 {
			Logger.Println("ERROR: request failed:", err)
		}
		return 0, hash.HashBlock{}, NewInternalError("request failed", err)
	}
	defer resp.Body.Close()

	if LogLevel > 1 {
		Logger.Printf("Completed in %v\n", time.Since(start))
	}

	// extract response headers
	responseHeaders.ContentDisposition = resp.Header.Get("Content-Disposition")
	responseHeaders.ContentType = resp.Header.Get("Content-Type")
	responseHeaders.Hashes = hash.ParseString(resp.Header.Get("X-Trimmer-Hash"))
	responseHeaders.RequestId = resp.Header.Get("X-Request-Id")
	responseHeaders.SessionId = resp.Header.Get("X-Session-Id")
	responseHeaders.OAuthScopes = resp.Header.Get("X-OAuth-Scopes")
	responseHeaders.Runtime = resp.Header.Get("X-Runtime")
	responseHeaders.Size = resp.ContentLength

	isJsonResponse := strings.Contains(resp.Header.Get("Content-Type"), "application/json")
	serverHash := hash.ParseString(resp.Header.Get("X-Trimmer-Hash"))
	if serverHash.Md5 == "" {
		serverHash.ParseEtag(resp.Header.Get("ETag"))
	}

	// on failure, return API error
	if resp.StatusCode >= 400 {

		// API responses in JSON get handled here
		if LogLevel > 2 {
			s, _ := httputil.DumpResponse(resp, true)
			log.Println(string(s))
		}

		// clear session on 401
		if resp.StatusCode == 401 && sess != nil {
			sess.Reset()
		}

		// parse error
		e := NewApiError(resp.StatusCode)
		if isJsonResponse && req.Method != http.MethodHead && resp.ContentLength != 0 {
			e = ParseApiError(resp.Body)
		}
		e.RequestId = resp.Header.Get("X-Request-Id")
		e.SessionId = resp.Header.Get("X-Session-Id")
		e.OauthScopes = resp.Header.Get("X-OAuth-Scopes")
		if LogLevel > 0 {
			Logger.Println("ERROR:", e.Error())
		}
		return resp.ContentLength, serverHash, e
	}

	// on success parse the response
	if isJsonResponse {

		// API responses in JSON get handled here
		if LogLevel > 2 {
			s, _ := httputil.DumpResponse(resp, true)
			log.Println(string(s))
		}

		if v != nil && (resp.ContentLength > 0 || resp.ContentLength == -1) {
			jsonDecoder := json.NewDecoder(resp.Body)
			if err := jsonDecoder.Decode(v); err != nil {
				return resp.ContentLength, serverHash, NewInternalError("parsing response failed", err)
			}
		}

	} else {

		// binary responses like downloading data are handled here
		if w, ok := v.(io.Writer); ok {
			var size int64
			if resp.ContentLength > 0 {
				// read exactly N bytes (returns error if operation ends early)
				size, err = io.CopyN(w, resp.Body, resp.ContentLength)
			} else {
				// read until EOF (returns no error on EOF!)
				size, err = io.Copy(w, resp.Body)
			}

			if err != nil {
				return size, serverHash, NewInternalError("copying response failed", err)
			}

			return size, serverHash, nil
		}
	}

	return resp.ContentLength, serverHash, nil
}
