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

// MediaListParams is the set of parameters that can be used when listing media.
type ProfileListParams struct {
	ListParams
	Types    MediaTypeList   `json:"type,omitempty"`
	Formats  MediaFormatList `json:"format,omitempty"`
	Families MediaFamilyList `json:"family,omitempty"`
	Embed    ApiEmbedFlags   `json:"embed,omitempty"`
}

// Profile is the resource representing a Trimmer transcoding profile.
type Profile struct {
	ID          int64           `json:"profileId"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Family      MediaFamily     `json:"family"`
	Type        MediaType       `json:"type"`
	Format      MediaFormat     `json:"format"`
	Role        MediaRole       `json:"role"`
	Mime        string          `json:"mimetype"`
	Details     *ProfileDetails `json:"details"`
}

type ProfileList []*Profile

type ProfileDetails struct {
	Image    *ProfileImageDetails    `json:"image"`
	Video    *ProfileVideoDetails    `json:"video"`
	Audio    *ProfileAudioDetails    `json:"audio"`
	Subtitle *ProfileSubtitleDetails `json:"subtitle"`
}

type ProfileVideoDetails struct {
	Codec           VideoCodec     `json:"codec"`
	Width           int            `json:"width"`
	Height          int            `json:"height"`
	Strict          bool           `json:"strict"`
	Bitrate         string         `json:"bitrate"`
	Pixel           PixelFormat    `json:"pixelFormat"`
	Sar             string         `json:"sar"`
	ColorModel      ColorModel     `json:"colorModel"`
	ColorTransfer   ColorTransfer  `json:"colorTransfer"`
	ColorPrimaries  ColorPrimaries `json:"colorPrimaries"`
	ColorRange      ColorRange     `json:"colorRange"`
	Profile         string         `json:"profile"`
	Level           string         `json:"level"`
	Preset          string         `json:"preset"`
	Tune            string         `json:"tune"`
	RCMode          string         `json:"rcMode"`
	RCValue         float32        `json:"rcValue"`
	MeMethod        string         `json:"meMethod"`
	DenoiseMethod   string         `json:"denoiseMethod"`
	DenoisePreset   string         `json:"denoisePreset"`
	DenoiseTune     string         `json:"denoiseTune"`
	DenoiseValue    string         `json:"denoiseValue"`
	SegmentDuration string         `json:"segmentDuration"`
	MinTracks       int            `json:"minTracks"`
	MaxTracks       int            `json:"maxTracks"`
}

type ProfileAudioDetails struct {
	Codec           AudioCodec  `json:"codec"`
	Bitrate         string      `json:"bitrate"`
	Channels        int         `json:"channels"`
	Layout          AudioLayout `json:"layout"`
	Profile         string      `json:"profile"`
	SampleBits      int         `json:"sampleBits"`
	SampleRate      int         `json:"sampleRate"`
	SampleFormat    string      `json:"sampleFormat"`
	SegmentDuration string      `json:"segmentDuration"`
	MinTracks       int         `json:"minTracks"`
	MaxTracks       int         `json:"maxTracks"`
}

type ProfileSubtitleDetails struct {
	Codec   SubtitleCodec `json:"codec"`
	Profile string        `json:"profile"`
}

type ProfileImageDetails struct {
	Width  int     `json:"width"`
	Height int     `json:"height"`
	GridX  int     `json:"gridX"`
	GridY  int     `json:"gridY"`
	Frames int     `json:"frames"`
	Rate   float32 `json:"rate"`
}
