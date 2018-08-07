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
	"context"
	"io"
	"net/http"
	"path/filepath"

	trimmer "trimmer.io/go-trimmer"
	"trimmer.io/go-trimmer/hash"
	"trimmer.io/go-trimmer/rfc"
)

type MultiFileSaver func(fi *trimmer.FileInfo) (io.Writer, error)

func Download(ctx context.Context, src *trimmer.Media, dst io.Writer) (*trimmer.FileInfo, error) {
	return getC().Download(ctx, src, dst)
}

func DownloadUrl(ctx context.Context, uri string, h hash.HashBlock, dst io.Writer) (*trimmer.FileInfo, error) {
	return getC().DownloadUrl(ctx, uri, h, dst)
}

func DownloadMulti(ctx context.Context, src *trimmer.Media, save MultiFileSaver) error {
	return getC().DownloadMulti(ctx, src, save)
}

func (c Client) Download(ctx context.Context, src *trimmer.Media, dst io.Writer) (*trimmer.FileInfo, error) {
	if src == nil || dst == nil {
		return nil, trimmer.ENilPointer
	}
	return c.DownloadUrl(ctx, src.Url, src.Hashes, dst)
}

func (c Client) DownloadUrl(ctx context.Context, uri string, h hash.HashBlock, dst io.Writer) (*trimmer.FileInfo, error) {
	if uri == "" {
		return nil, trimmer.EParamMissing
	}

	ch := &trimmer.CallHeaders{
		Accept: "*/*",
	}

	size, clientHashes, serverHashes, err := c.CDN.CallChecksum(ctx, http.MethodGet, uri, c.Key, c.Sess, ch, h.AnyFlag(), nil, dst, nil)
	if err != nil {
		return nil, err
	}

	if err = clientHashes.Check(h, true); err != nil {
		if trimmer.LogLevel > 0 {
			trimmer.Logger.Println("ERROR: checksum mismatch", err.Error())
		}
		return nil, err
	}

	cd := rfc.ParseContentDisposition(ch.ContentDisposition)

	fi := &trimmer.FileInfo{
		Size:     size,
		Mimetype: ch.ContentType,
		Etag:     serverHashes.Etag(),
		Hashes:   serverHashes,
		Filename: cd.Get("filename"),
		UUID:     cd.Get("uuid"),
		Url:      uri,
	}

	if fi.Filename == "" {
		fi.Filename = rfc.Basename(uri)
	}

	return fi, err
}

// support multi-file media like multi-resolution images and image sequences
func (c Client) DownloadMulti(ctx context.Context, src *trimmer.Media, save MultiFileSaver) error {
	if src == nil || save == nil {
		return trimmer.ENilPointer
	}

	if !IsMultiFileMediaType(src.Type) {
		return trimmer.EParamInvalid
	}

	if src.Attr == nil {
		return trimmer.EParamMissing
	}

	// sequence and grid media
	for _, s := range src.Attr.Sequence {
		for _, v := range s.MediaList {
			if v.Url == "" {
				continue
			}
			fi := &trimmer.FileInfo{
				Size:     v.Size,
				Hashes:   v.Hashes,
				Etag:     v.Hashes.Etag(),
				Filename: filepath.Join(src.Filename, v.Filename),
				UUID:     v.UUID,
				Mimetype: src.Mimetype,
				Url:      v.Url,
			}
			w, err := save(fi)
			if err != nil {
				return err
			}

			m := &trimmer.Media{
				Hashes: v.Hashes,
				Url:    v.Url,
			}
			if _, err = c.Download(ctx, m, w); err != nil {
				return err
			}
		}
	}

	// image media
	for _, v := range src.Attr.Image {
		if v.Url == "" {
			continue
		}
		fi := &trimmer.FileInfo{
			Size:     v.Size,
			Hashes:   v.Hashes,
			Etag:     v.Hashes.Etag(),
			Filename: filepath.Join(src.Filename, v.Filename),
			UUID:     v.UUID,
			Mimetype: src.Mimetype,
			Url:      v.Url,
		}
		w, err := save(fi)
		if err != nil {
			return err
		}

		m := &trimmer.Media{
			Hashes: v.Hashes,
			Url:    v.Url,
		}
		if _, err = c.Download(ctx, m, w); err != nil {
			return err
		}
	}

	// grid media
	for _, v := range src.Attr.Grid {
		if v.Url == "" {
			continue
		}
		fi := &trimmer.FileInfo{
			Size:     v.Size,
			Hashes:   v.Hashes,
			Etag:     v.Hashes.Etag(),
			Filename: filepath.Join(src.Filename, v.Filename),
			UUID:     v.UUID,
			Mimetype: src.Mimetype,
			Url:      v.Url,
		}
		w, err := save(fi)
		if err != nil {
			return err
		}

		m := &trimmer.Media{
			Hashes: v.Hashes,
			Url:    v.Url,
		}
		if _, err = c.Download(ctx, m, w); err != nil {
			return err
		}
	}
	return nil
}
