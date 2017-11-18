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

package media

import (
	trimmer "trimmer.io/go-trimmer"
)

func IndexTracks(d *trimmer.MediaAttr) {
	if d == nil {
		return
	}
	d.TrackIndex = make([]*trimmer.MediaTrack, 0)
	for _, v := range d.Video {
		d.TrackIndex = append(d.TrackIndex, &trimmer.MediaTrack{
			UUID:     v.UUID,
			TrackNum: v.TrackNum,
			Type:     MediaTypeVideo,
		})
	}

	for _, v := range d.Audio {
		d.TrackIndex = append(d.TrackIndex, &trimmer.MediaTrack{
			UUID:     v.UUID,
			TrackNum: v.TrackNum,
			Type:     MediaTypeAudio,
		})
	}

	for _, v := range d.Sequence {
		d.TrackIndex = append(d.TrackIndex, &trimmer.MediaTrack{
			UUID:     v.UUID,
			TrackNum: v.TrackNum,
			Type:     MediaTypeSequence,
		})
	}

	for _, v := range d.Image {
		d.TrackIndex = append(d.TrackIndex, &trimmer.MediaTrack{
			UUID:     v.UUID,
			TrackNum: v.TrackNum,
			Type:     MediaTypeImage,
		})
	}

	for _, v := range d.Grid {
		d.TrackIndex = append(d.TrackIndex, &trimmer.MediaTrack{
			UUID:     v.UUID,
			TrackNum: v.TrackNum,
			Type:     MediaTypeGrid,
		})
	}

	for _, v := range d.Subtitle {
		d.TrackIndex = append(d.TrackIndex, &trimmer.MediaTrack{
			UUID:     v.UUID,
			TrackNum: v.TrackNum,
			Type:     MediaTypeSubtitle,
		})
	}

	for _, v := range d.Data {
		d.TrackIndex = append(d.TrackIndex, &trimmer.MediaTrack{
			UUID:     v.UUID,
			TrackNum: v.TrackNum,
			Type:     MediaTypeData,
		})
	}

	for _, v := range d.Document {
		d.TrackIndex = append(d.TrackIndex, &trimmer.MediaTrack{
			UUID:     v.UUID,
			TrackNum: v.TrackNum,
			Type:     MediaTypeDocument,
		})
	}

	for _, v := range d.Manifest {
		d.TrackIndex = append(d.TrackIndex, &trimmer.MediaTrack{
			UUID:     v.UUID,
			TrackNum: v.TrackNum,
			Type:     MediaTypeManifest,
		})
	}
}

// deep copy media
func DeepCopyMedia(m *trimmer.Media) *trimmer.Media {
	if m == nil {
		return nil
	}

	target := &trimmer.Media{}
	*target = *m

	if m.Workspace != nil {
		target.Workspace = &trimmer.Workspace{}
		*target.Workspace = *m.Workspace
	}

	if m.Account != nil {
		target.Account = &trimmer.User{}
		*target.Account = *m.Account
	}

	if m.Author != nil {
		target.Author = &trimmer.User{}
		*target.Author = *m.Author
	}

	if m.Attr != nil {
		target.Attr = trimmer.NewMediaAttr()

		// copy track index
		for _, v := range m.Attr.TrackIndex {
			n := &trimmer.MediaTrack{}
			*n = *v
			target.Attr.TrackIndex = append(target.Attr.TrackIndex, n)
		}

		// copy track references
		for _, v := range m.Attr.References {
			n := &trimmer.MediaReference{}
			*n = *v
			target.Attr.References = append(target.Attr.References, n)
		}

		// copy video tracks
		for _, v := range m.Attr.Video {
			n := &trimmer.VideoAttr{}
			*n = *v
			target.Attr.Video = append(target.Attr.Video, n)
		}

		// copy audio tracks
		for _, v := range m.Attr.Audio {
			n := &trimmer.AudioAttr{}
			*n = *v
			target.Attr.Audio = append(target.Attr.Audio, n)
		}

		// copy sequence tracks
		for _, v := range m.Attr.Sequence {
			n := &trimmer.SequenceAttr{}
			*n = *v
			n.MediaList = make(trimmer.SequenceFileList, len(v.MediaList))
			for i, vv := range v.MediaList {
				nn := &trimmer.SequenceFile{}
				*nn = *vv
				n.MediaList[i] = nn
			}
			target.Attr.Sequence = append(target.Attr.Sequence, n)
		}

		// copy image tracks
		for _, v := range m.Attr.Image {
			n := &trimmer.ImageAttr{}
			*n = *v
			target.Attr.Image = append(target.Attr.Image, n)
		}

		// copy grid tracks
		for _, v := range m.Attr.Grid {
			n := &trimmer.GridAttr{}
			*n = *v
			target.Attr.Grid = append(target.Attr.Grid, n)
		}

		// copy subtitle tracks
		for _, v := range m.Attr.Subtitle {
			n := &trimmer.SubtitleAttr{}
			*n = *v
			target.Attr.Subtitle = append(target.Attr.Subtitle, n)
		}

		// copy data tracks
		for _, v := range m.Attr.Data {
			n := &trimmer.DataAttr{}
			*n = *v
			target.Attr.Data = append(target.Attr.Data, n)
		}

		// copy document tracks
		for _, v := range m.Attr.Document {
			n := &trimmer.DocumentAttr{}
			*n = *v
			target.Attr.Document = append(target.Attr.Document, n)
		}

		// copy manifest tracks
		for _, v := range m.Attr.Manifest {
			n := &trimmer.ManifestAttr{}
			*n = *v
			target.Attr.Manifest = append(target.Attr.Manifest, n)
		}
	}
	return target
}

// use for dropping all urls before returning metadata to the API
func StripMetadataUrls(d *trimmer.MediaAttr) {
	if d == nil {
		return
	}
	for _, v := range d.Video {
		v.Url = ""
	}

	for _, v := range d.Audio {
		v.Url = ""
	}

	for _, s := range d.Sequence {
		for _, v := range s.MediaList {
			v.Url = ""
		}
	}

	for _, v := range d.Image {
		v.Url = ""
	}

	for _, v := range d.Grid {
		v.Url = ""
	}

	for _, v := range d.Subtitle {
		v.Url = ""
	}

	for _, v := range d.Data {
		v.Url = ""
	}

	for _, v := range d.Document {
		v.Url = ""
	}

	for _, v := range d.Manifest {
		v.Url = ""
	}
}

// use for cleaning up single-file media
func StripMetadataFilenames(d *trimmer.MediaAttr) {
	if d == nil {
		return
	}
	for _, v := range d.Video {
		v.Filename = ""
	}

	for _, v := range d.Audio {
		v.Filename = ""
	}

	for _, v := range d.Subtitle {
		v.Filename = ""
	}

	for _, v := range d.Data {
		v.Filename = ""
	}

	for _, v := range d.Document {
		v.Filename = ""
	}

	for _, v := range d.Manifest {
		v.Url = ""
	}
}
