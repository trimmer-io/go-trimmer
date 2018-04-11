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

package media

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	trimmer "trimmer.io/go-trimmer"
	"trimmer.io/go-trimmer/hash"
	"trimmer.io/go-trimmer/job"
	"trimmer.io/go-trimmer/rfc"
)

type PartState string

const (
	PartStateIdle     PartState = "idle"
	PartStateProgress PartState = "progress"
	PartStateComplete PartState = "complete"
)

type UploadState string

const (
	UploadStateInvalid  UploadState = "invalid"
	UploadStateIdle     UploadState = "idle"
	UploadStateProgress UploadState = "progress"
	UploadStateComplete UploadState = "complete"
	UploadStateAborted  UploadState = "aborted"
)

type UploadInfo struct {
	VolumeUUID    string         `json:"volumeUuid"`              // volume UUID
	Key           string         `json:"key"`                     // filename on storage
	ContentType   string         `json:"mimetype"`                // upload content type
	UploadId      string         `json:"uploadId"`                // multi-part upload id
	State         UploadState    `json:"state"`                   // upload status
	Hashes        hash.HashBlock `json:"hashes,omitempty"`        // hashes for total content
	Size          int64          `json:"size,omitempty"`          // total size
	Expires       time.Time      `json:"expires"`                 // upload expiry time
	TotalParts    int64          `json:"totalParts,omitempty"`    // number of parts finished and in progress
	ProgressParts int64          `json:"progressParts,omitempty"` // number of parts in progress
	Part          *PartInfo      `json:"part,omitempty"`          // current part
}

type PartListResponse struct {
	UploadId string       `json:"uploadId"`
	Key      string       `json:"key"`
	Count    int          `json:"count"`
	Skip     int          `json:"skip"`
	Parts    PartInfoList `json:"parts"`
}

type PartInfo struct {
	PartId     int            `json:"partId"`
	Hashes     hash.HashBlock `json:"hashes"`
	Etag       string         `json:"etag"`
	Size       int64          `json:"size"`
	State      PartState      `json:"state"`
	UploadedAt time.Time      `json:"uploadedAt"`
}

type PartInfoList []*PartInfo

type ProgressFunc func(ctx context.Context, r *UploadRequest, size int64)

type UploadRequest struct {
	C            Client
	Reader       io.ReadSeeker
	Media        *trimmer.Media
	Size         int64
	JobId        string
	Url          string
	Query        url.Values
	Filename     string
	UUID         string
	Mimetype     string
	Hashes       hash.HashBlock
	UploadId     string
	PartNum      int64
	PartSize     int64
	UploadedSize int64
	Manifest     *trimmer.VolumeManifest
	Progress     ProgressFunc
}

type ManifestCache struct {
	cache map[string]*trimmer.VolumeManifest
	m     sync.RWMutex
}

var manifestCache = &ManifestCache{
	cache: make(map[string]*trimmer.VolumeManifest),
}

func (c *ManifestCache) GetManifest(prefix string) *trimmer.VolumeManifest {
	c.m.RLock()
	m, _ := c.cache[prefix]
	c.m.RUnlock()
	return m
}

func (c *ManifestCache) SetManifest(m *trimmer.VolumeManifest) {
	if m == nil {
		return
	}
	prefix := strings.Trim(m.UrlPrefix, "/")
	c.m.Lock()
	c.cache[prefix] = m
	c.m.Unlock()
}

func Upload(ctx context.Context, m *trimmer.Media, src io.ReadSeeker) (*trimmer.FileInfo, error) {
	return getC().Upload(ctx, m, src)
}

type MultiFileLoader func(fi *trimmer.FileInfo) (io.ReadSeeker, error)

func UploadMulti(ctx context.Context, m *trimmer.Media, load MultiFileLoader) (trimmer.FileInfoList, error) {
	return getC().UploadMulti(ctx, m, load)
}

func NewUploadRequest(fi *trimmer.FileInfo, dst *trimmer.Media, src io.ReadSeeker) *UploadRequest {
	return getC().NewUploadRequest(fi, dst, src)
}

func UploadImage(ctx context.Context, uri string, params *trimmer.FileInfo, src io.Reader) (*trimmer.Media, error) {
	return getC().UploadImage(ctx, uri, params, src)
}

func (c Client) Upload(ctx context.Context, dst *trimmer.Media, src io.ReadSeeker) (*trimmer.FileInfo, error) {
	if dst == nil {
		return nil, trimmer.ENilPointer
	}

	if dst.Url == "" {
		return nil, trimmer.EParamMissing
	}

	fi := &trimmer.FileInfo{
		Size:     dst.Size,
		Hashes:   dst.Hashes,
		Etag:     dst.Hashes.Etag(),
		Filename: dst.Filename,
		UUID:     dst.UUID,
		Mimetype: dst.Mimetype,
		Url:      dst.Url,
	}

	// overwrites media on completion
	var err error
	r := c.NewUploadRequest(fi, dst, src)
	fi, err = r.Do(ctx)
	if err != nil {
		return nil, err
	}

	// complete upload job when callback is missing
	if dst.JobId != "" {
		if !r.HasCallback() {
			up := &trimmer.MediaUploadCompletionParams{
				Files: trimmer.FileInfoList{fi},
				Embed: trimmer.API_EMBED_META | trimmer.API_EMBED_DETAILS,
			}
			if m, err := c.CompleteUpload(ctx, dst.ID, up); err != nil {
				return fi, err
			} else {
				*dst = *m
			}
		} else {
			// re-read media after upload
			mp := &trimmer.MediaParams{
				Embed: trimmer.API_EMBED_META | trimmer.API_EMBED_DETAILS,
			}
			if m, err := c.Get(ctx, dst.ID, mp); err != nil {
				return fi, err
			} else {
				*dst = *m
			}
		}
	}

	return fi, nil
}

func appendKeyPath(src, part string) string {
	u, err := url.Parse(src)
	if err != nil {
		return src
	}
	q := u.Query()
	key := q.Get("key")
	q.Set("key", strings.Join([]string{key, part}, "/"))
	u.RawQuery = q.Encode()
	return u.String()
}

func (c Client) UploadMulti(ctx context.Context, dst *trimmer.Media, load MultiFileLoader) (trimmer.FileInfoList, error) {
	if dst == nil || load == nil {
		return nil, trimmer.ENilPointer
	}

	if !IsMultiFileMediaType(dst.Type) {
		return nil, trimmer.EParamInvalid
	}

	if dst.Attr == nil {
		return nil, trimmer.EParamMissing
	}

	if dst.Url == "" {
		return nil, trimmer.EParamMissing
	}

	var uploadedBytes int64
	uploadedFiles := make(trimmer.FileInfoList, 0)

	// whatever happens, reset progress state
	defer func() { c.lastProgress = time.Time{} }()

	// sequence and grid media
	for _, s := range dst.Attr.Sequence {
		for _, v := range s.MediaList {
			if v.Filename == "" {
				continue
			}
			// skip already uploaded files (one's that have a URL and Hashes set)
			if v.Url != "" && !v.Hashes.IsZero() {
				uploadedBytes += v.Size
				continue
			}
			fi := &trimmer.FileInfo{
				Size:     v.Size,
				Hashes:   v.Hashes,
				Etag:     v.Hashes.Etag(),
				Filename: v.Filename,
				UUID:     v.UUID,
				Mimetype: dst.Mimetype,
				Url:      appendKeyPath(dst.Url, v.Filename),
			}
			f, err := load(fi)
			if err != nil {
				return uploadedFiles, err
			}

			// append filename to base upload URL (key query parameter)
			// store size and hashes with metadata
			r := c.NewUploadRequest(fi, dst, f)
			r.Progress = nil // no auto-update because uploadedBytes is unknown
			fi, err = r.Do(ctx)
			if err != nil {
				return uploadedFiles, err
			}
			// store upload hashes in attr
			v.Hashes = fi.Hashes
			uploadedBytes += fi.Size
			uploadedFiles = append(uploadedFiles, fi)
			c.ProgressUpload(ctx, r, uploadedBytes)
		}
	}

	// image media
	for _, v := range dst.Attr.Image {
		// skip files without name
		if v.Filename == "" {
			continue
		}
		// skip already uploaded files (one's that have a URL and Hashes set)
		if v.Url != "" && !v.Hashes.IsZero() {
			uploadedBytes += v.Size
			continue
		}
		fi := &trimmer.FileInfo{
			Size:     v.Size,
			Hashes:   v.Hashes,
			Etag:     v.Hashes.Etag(),
			Filename: v.Filename,
			UUID:     v.UUID,
			Mimetype: dst.Mimetype,
			Url:      appendKeyPath(dst.Url, v.Filename),
		}
		f, err := load(fi)
		if err != nil {
			return uploadedFiles, err
		}

		// append filename to base upload URL (key query parameter)
		// store size and hashes with metadata
		r := c.NewUploadRequest(fi, dst, f)
		r.Progress = nil // no auto-update because uploadedBytes is unknown
		fi, err = r.Do(ctx)
		if err != nil {
			return uploadedFiles, err
		}
		// store upload hashes in attr
		v.Hashes = fi.Hashes
		uploadedBytes += fi.Size
		uploadedFiles = append(uploadedFiles, fi)
		c.ProgressUpload(ctx, r, uploadedBytes)
	}

	// grid media
	for _, v := range dst.Attr.Grid {
		if v.Filename == "" {
			continue
		}
		// skip already uploaded files (one's that have a URL and Hashes set)
		if v.Url != "" && !v.Hashes.IsZero() {
			uploadedBytes += v.Size
			continue
		}
		fi := &trimmer.FileInfo{
			Size:     v.Size,
			Hashes:   v.Hashes,
			Etag:     v.Hashes.Etag(),
			Filename: v.Filename,
			UUID:     v.UUID,
			Mimetype: dst.Mimetype,
			Url:      appendKeyPath(dst.Url, v.Filename),
		}
		f, err := load(fi)
		if err != nil {
			return uploadedFiles, err
		}

		// append filename to base upload URL (key query parameter)
		// store size and hashes with metadata
		r := c.NewUploadRequest(fi, dst, f)
		r.Progress = nil // no auto-update because uploadedBytes is unknown
		fi, err = r.Do(ctx)
		if err != nil {
			return uploadedFiles, err
		}
		// store upload hashes in attr
		v.Hashes = fi.Hashes
		uploadedBytes += fi.Size
		uploadedFiles = append(uploadedFiles, fi)
		c.ProgressUpload(ctx, r, uploadedBytes)
	}

	// complete upload job when job id is missing (multi-file media never uses
	// callbacks in urls)
	if dst.JobId != "" {
		up := &trimmer.MediaUploadCompletionParams{
			Files: uploadedFiles,
			Embed: trimmer.API_EMBED_META | trimmer.API_EMBED_DETAILS,
		}
		if m, err := c.CompleteUpload(ctx, dst.ID, up); err != nil {
			return uploadedFiles, err
		} else {
			*dst = *m
		}
	}

	return uploadedFiles, nil
}

func (c Client) ProgressUpload(ctx context.Context, r *UploadRequest, size int64) {

	// check all preconditions for progress updates
	if r == nil || r.Media == nil || r.Media.JobId == "" || r.HasCallback() || r.Media.Size == 0 {
		return
	}

	// throttle progress callbacks
	now := time.Now().UTC()
	if !c.lastProgress.IsZero() && c.lastProgress.Add(10*time.Second).Before(now) {
		return
	}

	p := &trimmer.JobParams{
		Progress: int(size * 100 / r.Media.Size),
	}
	if _, err := job.Update(ctx, r.Media.JobId, p); err == nil {
		c.lastProgress = now
	}
	return
}

func (c Client) NewUploadRequest(fi *trimmer.FileInfo, dst *trimmer.Media, src io.ReadSeeker) *UploadRequest {

	r := &UploadRequest{
		C:        c,
		Reader:   src,
		Media:    dst,
		Size:     fi.Size,
		PartNum:  1,
		Filename: fi.Filename,
		UUID:     fi.UUID,
		Mimetype: fi.Mimetype,
		Hashes:   fi.Hashes,
		Progress: c.ProgressUpload,
		Query:    url.Values{},
	}

	parts := strings.Split(fi.Url, "?")
	switch len(parts) {
	case 1:
		r.Url = strings.Trim(parts[0], "/")
	case 2:
		r.Url = parts[0]
		r.Query, _ = url.ParseQuery(parts[1])
	default:
		return &UploadRequest{}
	}

	// always pass expected size as query parameter
	r.Query.Set("size", strconv.FormatInt(r.Size, 10))
	return r
}

// Extract volume path from upload URL of form
//
// http://127.0.0.1/public/uploads?key=..
//   -> public
// http://127.0.0.1/my/storage/path/uploads?key=..
//   -> my/storage/path
//
func (r *UploadRequest) VolumePrefix() string {
	u, err := url.Parse(r.Url)
	if err != nil {
		return ""
	}
	// check if this is an upload URL
	if !strings.Contains(u.Path, "/uploads") {
		return ""
	}
	return strings.Trim(strings.Split(u.Path, "/uploads")[0], "/")
}

func (r UploadRequest) HasCallback() bool {
	return r.Query.Get("cb") != ""
}

func (r *UploadRequest) SingleUrl() string {
	return fmt.Sprintf("%s?%s", r.Url, r.Query.Encode())
}

func (r *UploadRequest) ManifestUrl() string {
	return strings.Replace(r.SingleUrl(), "uploads", "manifest.json", 1)
}

func (r *UploadRequest) InitMultiUrl() string {
	return r.SingleUrl()
}

func (r *UploadRequest) UploadMultiUrl(part int64) string {
	r.Query.Set("part", strconv.FormatInt(part, 10))
	return fmt.Sprintf("%s/%s?%s", r.Url, r.UploadId, r.Query.Encode())
}

func (r *UploadRequest) EndMultiUrl() string {
	r.Query.Del("part")
	return fmt.Sprintf("%s/%s?%s", r.Url, r.UploadId, r.Query.Encode())
}

func (r *UploadRequest) uploadSingle(ctx context.Context) (int64, hash.HashBlock, error) {

	ct := rfc.NewContentDisposition("attachement")
	ct.Set("filename", r.Filename)
	ct.Set("uuid", r.UUID)

	// Volume requires special headers
	h := &trimmer.CallHeaders{
		ContentType:        r.Mimetype,
		ContentDisposition: ct.Encode(),
		Hashes:             r.Hashes,
		Size:               r.Size,
	}

	if trimmer.LogLevel > 2 {
		trimmer.Logger.Printf("Uploading single file %s %d bytes", r.Filename, r.Size)
	}

	i := &UploadInfo{}
	_, clientHashes, serverHashes, err := r.C.CDN.CallChecksum(ctx, "PUT", r.SingleUrl(), r.C.Key, r.C.Sess, h, r.Hashes.AnyFlag(), r.Reader, nil, i)
	if err != nil {
		return 0, hash.HashBlock{}, err
	}

	if err = clientHashes.Check(serverHashes, true); err != nil {
		if trimmer.LogLevel > 0 {
			trimmer.Logger.Println("ERROR: checksum mismatch", err.Error())
		}
		return i.Size, i.Hashes, err
	}

	if trimmer.LogLevel > 2 {
		trimmer.Logger.Printf("Upload success VolumeUUID=%s hashes=%s", i.VolumeUUID, i.Hashes.String())
	}

	return i.Size, i.Hashes, nil
}

func (r *UploadRequest) initMulti(ctx context.Context) error {

	ct := rfc.NewContentDisposition("attachement")
	ct.Set("filename", r.Filename)
	ct.Set("uuid", r.UUID)

	// CDN requires special headers
	h := &trimmer.CallHeaders{
		ContentType:        r.Mimetype,
		ContentDisposition: ct.Encode(),
	}

	upload := &UploadInfo{}
	err := r.C.CDN.Call(ctx, "POST", r.InitMultiUrl(), r.C.Key, r.C.Sess, h, nil, upload)
	if err != nil {
		return err
	}

	// check upload state, assuming sequential upload progress
	if upload.State == UploadStateProgress {
		r.PartNum = upload.TotalParts + 1
	}

	r.UploadId = upload.UploadId
	return nil
}

func (r *UploadRequest) uploadPart(ctx context.Context, reader io.Reader, overwrite bool) error {
	// don't send on empty upload id
	if r.UploadId == "" {
		return trimmer.NewInternalError("empty upload id", nil)
	}

	// CDN requires special headers
	h := &trimmer.CallHeaders{
		ContentType: r.Mimetype,
	}

	// signal a part is supposed to be overwritten
	// (i.e. because of earlier checksum failure)
	if overwrite {
		r.Query.Set("overwrite", "true")
	} else {
		r.Query.Del("overwrite")
	}

	i := &UploadInfo{}
	_, clientHash, serverHash, err := r.C.CDN.CallChecksum(ctx, "PUT", r.UploadMultiUrl(r.PartNum), r.C.Key, r.C.Sess, h, hash.HASH_TYPE_SHA256, io.LimitReader(reader, r.PartSize), nil, i)
	if err != nil {
		return err
	}

	// verify part checksums
	//
	// Note: we usually don't have part-wise checksums before having sent a part, hence we
	//       cannot let the CDN server check integrity for us. Although the entire file
	//       gets checked at commit(), it would be nice if we could check each individual
	//       part already during upload to detect and repair transmission errors early.
	//
	//       For this reason the CDN server returns its default part checksum (SHA256)
	//       even if not explicitly asked to generate such a checksum (when our
	//       x-trimmer-hash header is missing)
	//
	if err := clientHash.Check(serverHash, true); err != nil {
		if trimmer.LogLevel > 0 {
			trimmer.Logger.Println("ERROR: checksum mismatch on part", r.PartNum, err.Error())
		}
		return err
	}

	r.UploadedSize += i.Part.Size

	return nil
}

func (r *UploadRequest) abortMulti(ctx context.Context) error {
	// don't abort empty uploads
	if r.UploadId == "" {
		return trimmer.NewInternalError("empty upload id", nil)
	}

	ct := rfc.NewContentDisposition("attachement")
	ct.Set("filename", r.Filename)
	ct.Set("uuid", r.UUID)

	// CDN requires special headers
	h := &trimmer.CallHeaders{
		ContentType:        r.Mimetype,
		ContentDisposition: ct.Encode(),
	}

	err := r.C.CDN.Call(ctx, "DELETE", r.EndMultiUrl(), r.C.Key, r.C.Sess, h, nil, nil)
	if err != nil {
		return err
	}

	// clear upload id
	r.UploadId = ""
	return nil
}

func (r *UploadRequest) commitMulti(ctx context.Context) (hash.HashBlock, error) {
	// don't commit on empty upload id
	if r.UploadId == "" {
		return hash.HashBlock{}, trimmer.NewInternalError("empty upload id", nil)
	}

	ct := rfc.NewContentDisposition("attachement")
	ct.Set("filename", r.Filename)
	ct.Set("uuid", r.UUID)

	// CDN requires special headers
	h := &trimmer.CallHeaders{
		ContentType:        r.Mimetype,
		ContentDisposition: ct.Encode(),
		Hashes:             r.Hashes,
	}

	i := &UploadInfo{}
	err := r.C.CDN.Call(ctx, "POST", r.EndMultiUrl(), r.C.Key, r.C.Sess, h, nil, i)
	if err != nil {
		return i.Hashes, err
	}

	// compare final server-side hash with original client-side hash
	if err := r.Hashes.Check(i.Hashes, true); err != nil {
		return i.Hashes, err
	}

	// clear upload id (so the subsequent `defer abort()` will not execute)
	r.UploadId = ""
	return i.Hashes, nil
}

func (r *UploadRequest) CalculatePartSize() int64 {

	// use user-defined minimum/default
	s := trimmer.UploadPartSize

	// requires volume manifest to be present
	if r.Manifest == nil {
		return s
	}

	if s < r.Manifest.Limits.PartSizeMin {
		s = r.Manifest.Limits.PartSizeMin
	}

	// double the part size as long as the entire file requires more
	// than half the allowed PartsMax, but stay below PartSizeMax
	halfmax := r.Manifest.Limits.PartsMax / 2
	for s*2 <= r.Manifest.Limits.PartSizeMax && s*halfmax < r.Size {
		s = 2 * s
	}

	// cross-check if part size is sufficient to upload all parts (Note: doubling
	// the value in the last step above might cross the PartSizeMax limit)
	if s*r.Manifest.Limits.PartsMax < r.Size {
		s = r.Manifest.Limits.PartSizeMax
	}

	return s
}

func (r *UploadRequest) uploadMulti(ctx context.Context) (size int64, hashes hash.HashBlock, err error) {

	// first, choose a reasonable part size between PartSizeMin & PartSizeMax
	r.PartSize = r.CalculatePartSize()

	// init upload
	if err = r.initMulti(ctx); err != nil {
		return
	}

	// on subsequent failure abort upload
	defer r.abortMulti(ctx)

	// upload all parts (assuming sizeHint is correct); consider continuation
	var sz int64 = (r.PartNum - 1) * r.PartSize
	var retries int = trimmer.MaxRetries
	var overwritePart bool = false
	r.Reader.Seek(sz, io.SeekStart)

	for sz < r.Size {
		// retry on checksum errors
		if err = r.uploadPart(ctx, r.Reader, overwritePart); err != nil {
			// fail when upload has been cancelled
			if e, ok := err.(trimmer.TrimmerError); ok && e.IsApi() && e.StatusCode == 404 {
				return
			}

			// FIXME: support retries on more transient error conditions
			//        such as io/network errors and 5xx server errors
			//
			retries--
			if retries == 0 || err != hash.EInvalidHash {
				return
			}
			r.Reader.Seek(sz, io.SeekStart)
			overwritePart = true
			err = nil
			wait := trimmer.RetryBackoffTime * time.Duration(trimmer.MaxRetries-retries-1)
			if trimmer.LogLevel > 1 {
				trimmer.Logger.Printf("Retrying upload in %v", wait)
			}
			time.Sleep(wait)

		} else {
			// auto-progress update
			if r.Progress != nil {
				r.Progress(ctx, r, sz)
			}

			// on success upload next part
			overwritePart = false
			retries = trimmer.MaxRetries
			r.PartNum++
			sz += r.PartSize
		}
	}

	// finish upload
	hashes, err = r.commitMulti(ctx)
	size = r.UploadedSize
	return
}

// hides the complexity of single/multipart upload handling
func (r *UploadRequest) Do(ctx context.Context) (*trimmer.FileInfo, error) {

	// must have a valid URL
	if r.Url == "" {
		return nil, trimmer.EParamMissing
	}

	// fetch volume manifest from cache
	r.Manifest = manifestCache.GetManifest(r.VolumePrefix())
	if r.Manifest == nil {
		r.Manifest = &trimmer.VolumeManifest{}
		if err := r.C.CDN.Call(ctx, "GET", r.ManifestUrl(), r.C.Key, r.C.Sess, nil, nil, &r.Manifest); err != nil {
			return nil, err
		}
		// check the manifest is valid
		if trimmer.IsNilUUID(r.Manifest.UUID) {
			return nil, trimmer.NewUsageError("invalid volume manifest", nil)
		}
		manifestCache.SetManifest(r.Manifest)
	}

	var (
		err    error
		size   int64
		hashes hash.HashBlock
	)

	if r.Size < r.Manifest.Limits.SinglePartMax {
		size, hashes, err = r.uploadSingle(ctx)
	} else {
		size, hashes, err = r.uploadMulti(ctx)
	}

	if err != nil {
		return nil, err
	}

	fi := &trimmer.FileInfo{
		Filename:   r.Filename,
		Mimetype:   r.Mimetype,
		Size:       size,
		Hashes:     hashes,
		UUID:       r.UUID,
		VolumeUUID: r.Manifest.UUID,
	}

	return fi, nil
}

func (c Client) UploadImage(ctx context.Context, uri string, params *trimmer.FileInfo, src io.Reader) (*trimmer.Media, error) {
	if params == nil {
		return nil, trimmer.ENilPointer
	}

	// upload multipart-formdata
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	q := &url.Values{}

	if params.Filename != "" {
		q.Add("filename", params.Filename)
	}
	if params.Mimetype != "" {
		q.Add("mimetype", string(params.Mimetype))
	}
	if params.Role != "" {
		q.Add("role", string(params.Role))
	}
	if params.Embed.IsValid() {
		q.Add("embed", params.Embed.String())
	}

	// create a multipart-formdata
	h := make(textproto.MIMEHeader)
	ct := rfc.NewContentDisposition("form-data")
	ct.Set("filename", params.Filename)
	ct.Set("name", string(params.Role))
	h.Set("Content-Disposition", ct.Encode())
	h.Set("Content-Type", params.Mimetype)

	part, err := writer.CreatePart(h)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, src)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	ch := &trimmer.CallHeaders{
		ContentType: "multipart/form-data; boundary=" + writer.Boundary(),
	}

	v := &trimmer.Media{}
	u := fmt.Sprintf("%s?%s", uri, q.Encode())
	err = c.B.CallMultipart(ctx, "POST", u, c.Key, c.Sess, ch, body, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
