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
	"time"
)

type WorkflowAction string

type AccessAction struct {
	SubjectId   string         `json:"subjectId"`
	SubjectType AclSubjectType `json:"subjectType"`
	AccessClass AccessClass    `json:"access"`
	IpRange     string         `json:"ipRange"`
	ValidAfter  time.Time      `json:"validAfter"`
	ValidUntil  time.Time      `json:"validUntil"`
	Perm        Permission     `json:"permissions"`
}

type MetadataAction struct {
	Actions MetaValueList `json:"actions"`
}

type MediaAction struct {
	Match             MediaMatch         `json:"match"`
	TranscoderOptions *TranscoderOptions `json:"transcoderOpts,omitempty"`
	AnalyzerOptions   *AnalyzerOptions   `json:"analyzerOpts,omitempty"`
	RenderOptions     *RenderOptions     `json:"renderOpts,omitempty"`
	VolumeOptions     *VolumeOptions     `json:"volumeOpts,omitempty"`
}

type MediaMatch struct {
	Types     MediaTypeList     `json:"types"`
	Families  MediaFamilyList   `json:"families"`
	Roles     MediaRoleList     `json:"roles"`
	Relations MediaRelationList `json:"relations"`
}

// type AssetAction struct {
// }

// type NotifyAction struct {
// }

// type VolumeAction struct {
// 	WorkspaceId string `json:"workspaceId"`
// }

type AnalyzerOptions struct {
	ExtractMetadata        bool `json:"metadata"`
	ExtractDynamicMetadata bool `json:"dynamicMetadata"`
	ExtractLook            bool `json:"look"`
	ExtractThumbnail       bool `json:"thumbnail"`
}

type VolumeOptions struct {
	TargetVolumeId string `json:"volumeId,omitempty"`
}

type TranscoderOptions struct {
	General   GeneralConfig      `json:"general"`             // target profile and categories
	Metadata  MetaValueList      `json:"metadata,omitempty"`  // KV list for embedding into files
	Quality   *QualityConfig     `json:"quality,omitempty"`   // de-/encoder quality settings
	Crop      *CropConfig        `json:"crop,omitempty"`      // source video crop
	Transform *TransformConfig   `json:"transform,omitempty"` // source video transform
	Trim      *TrimConfig        `json:"trim,omitempty"`      // single-source audio/video trim
	Color     *ColorConfig       `json:"color,omitempty"`     // source video color processing
	Watermark []*WatermarkConfig `json:"watermark,omitempty"` // video watermark
	AudioMix  *AudioMixConfig    `json:"audio,omitempty"`     // audio routing
	ArriSDK   *ArriSDKConfig     `json:"arri_sdk,omitempty"`  // ARRI specific options
}

type RenderOptions struct {
	General  GeneralConfig    `json:"general"`            // target profile and categories
	Metadata MetaValueList    `json:"metadata,omitempty"` // to be embedded into files
	Quality  *QualityConfig   `json:"quality,omitempty"`  // encoder quality settings
	EditList []*SourceOptions `json:"editlist,omitempty"` // per-source options
}

type SourceOptions struct {
	MediaId   string           `json:"mediaId,omitempty"`   // source media id
	Crop      *CropConfig      `json:"crop,omitempty"`      // source video crop
	Transform *TransformConfig `json:"transform,omitempty"` // source video transform
	Trim      *TrimConfig      `json:"trim,omitempty"`      // source audio/video trim
	Color     *ColorConfig     `json:"color,omitempty"`     // source video color processing
}

type GeneralConfig struct {
	Profile           string         `json:"profile"`             // target transcode profile
	Width             int            `json:"width"`               // optional width override
	Height            int            `json:"height"`              // optional height override
	Role              *MediaRole     `json:"role"`                // target role
	Relation          *MediaRelation `json:"relation"`            // target relation
	Timecode          *string        `json:"timecode"`            // start timecode
	StartFrameNumber  int            `json:"start_frame_number"`  // start frame number for output (default: 1)
	FrameNumberDigits int            `json:"frame_number_digits"` // frame number precision for output (default: 7)
}

type QualityConfig struct {
	Preset        string  `json:"preset,omitempty"`   // codec-specific: e.g. x264 preset
	Tune          string  `json:"tune,omitempty"`     // codec-specific: e.g. grain, zerolatency, psnr for x264/5
	VideoProfile  string  `json:"vprofile,omitempty"` // codec-specific: video profile
	VideoLevel    string  `json:"vlevel,omitempty"`   // codec-specific: video profile level
	AudioProfile  string  `json:"aprofile,omitempty"` // codec-specific: audio profile
	RCMode        string  `json:"rc_mode,omitempty"`  // codec-specific: CRF, VBR, CBR, LOSSLESS, ...
	RCValue       float32 `json:"rc_value,omitempty"` // codec-specific: CRF: 18
	MeMethod      string  `json:"me_method"`          // x264/5 motion estimation override
	DenoiseMethod string  `json:"denoise_method"`     // ffmpeg filter (nlmeans, hqdn3d)
	DenoisePreset string  `json:"denoise_preset"`     // denoising preset
	DenoiseTune   string  `json:"denoise_tune"`       // denoising tune (nlmeans-only)
	DenoiseValue  string  `json:"denoise_value"`      // denoising custom setting
}

type ArriSDKConfig struct { // Arri SDK configuration options
	Version   string  `json:"version"`    // colorimetric processing version: 5.0
	Quality   string  `json:"quality"`    // quality mode: HQ, proxy1, proxy2
	Debayer   string  `json:"debayer"`    // debayer algorithm: ADA-1 to ADA-5 (SW/HW)
	Denoise   float32 `json:"denoise"`    // denoising strength: 0.0(off) 1.0 to 3.5
	TuneRed   int     `json:"tune_red"`   // fine-tuning ADA-5 SW: 0 to 100
	TuneGreen int     `json:"tune_green"` // fine-tuning ADA-5 SW: 0 to 100
	TuneBlue  int     `json:"tune_blue"`  // fine-tuning ADA-5 SW: 0 to 100
	Cct       float32 `json:"cct"`        // correlated color temperatore: 2000/3200 to 7000/11000
	Tint      float32 `json:"tint"`       // -12.0 to 12.0
	Crispness float32 `json:"crispness"`  // downscale crispness: 0.0 to 3.0
	ISO       int     `json:"iso"`        // AsaLUT ISO value: 50 to 500/1600/3200 depending on camera
}

type Box2i struct {
	Left   int `json:"left"`
	Top    int `json:"top"`
	Right  int `json:"right"`
	Bottom int `json:"bottom"`
}

type Point2i struct {
	Left int `json:"left"`
	Top  int `json:"top"`
}

type CropMode string

type CropConfig struct {
	Mode      CropMode `json:"mode,omitempty"`
	Inset     *Box2i   `json:"inset,omitempty"`
	Padding   *Box2i   `json:"padding,omitempty"`
	FillColor string   `json:"fill_color,omitempty"` // 0xFFFFFFFF [RRGGBBAA]
}

type FlipMode string

type RotateMode string

type TransformConfig struct {
	Flip   FlipMode    `json:"flip,omitempty"`
	Rotate *RotateMode `json:"rotate,omitempty"`
}

// source media in/out-points (relative to start of source)
type TrimConfig struct {
	StartFrame int64         `json:"start_frame,omitempty"` // inclusive, relative to media start
	EndFrame   int64         `json:"end_frame,omitempty"`   // not inclusive, relative to media start
	StartTime  time.Duration `json:"start_time,omitempty"`  // inclusive, relative to media start
	EndTime    time.Duration `json:"end_time,omitempty"`    // not inclusive, relative to media start
	StartTC    string        `json:"start_tc,omitempty"`    // inclusive, relative to source media TC
	EndTC      string        `json:"end_tc,omitempty"`      // not inclusive, relative to source media TC
}

type ColorConfig struct {
	Model     ColorModel     `json:"model,omitempty"`     // request color space conversion
	Primaries ColorPrimaries `json:"primaries,omitempty"` // request color space conversion
	Transfer  ColorTransfer  `json:"transfer,omitempty"`  // request color space conversion
	Range     ColorRange     `json:"range,omitempty"`     // request color range conversion
	CDL       *CDL           `json:"cdl,omitempty"`       // ASC CDL values
}

// color pipeline step when CDL is applied
type CDLMode string

type CDL struct {
	Mode       CDLMode `json:"mode,omitempty"` // apply in transcode step
	OffsetB    float32 `json:"offsetB"`        // = lift (dark tones)
	OffsetG    float32 `json:"offsetG"`        //
	OffsetR    float32 `json:"offsetR"`        //
	PowerB     float32 `json:"powerB"`         // = gamma (mid-tones)
	PowerG     float32 `json:"powerG"`         //
	PowerR     float32 `json:"powerR"`         //
	SlopeB     float32 `json:"slopeB"`         // = gain (highlights)
	SlopeG     float32 `json:"slopeG"`         //
	SlopeR     float32 `json:"slopeR"`         //
	Saturation float32 `json:"saturation"`     // all channels
}

type WatermarkMode string

type WatermarkConfig struct {
	Mode      WatermarkMode `json:"mode,omitempty"`
	Position  Point2i       `json:"position,omitempty"`   // watermark top/left coordinates
	Reference CropMode      `json:"reference,omitempty"`  // reference rectangle after cropping+padding, default: display
	MediaId   string        `json:"media_id,omitempty"`   // mode: logo
	Text      string        `json:"text,omitempty"`       // mode: text
	Label     MetaPath      `json:"label,omitempty"`      // mode: metadata key
	FontSize  int           `json:"font_size,omitempty"`  // modes: text, timecode, meta
	FontColor string        `json:"font_color,omitempty"` // modes: text, timecode, meta
	ClipColor string        `json:"clip_color,omitempty"` // mode: logo (transparency)
}

type WatermarkList []*WatermarkConfig

type AudioMixConfig struct {
	Layout AudioLayout `json:"layout"`
}
