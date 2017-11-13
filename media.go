// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package trimmer

import (
	"strings"
	"time"
	"trimmer.io/go-trimmer/hash"
)

// MediaState is the list of allowed values for a media's status.
// Allowed values are "created", "uploading", "uploaded", "analyzing",
// "transcoding", "failed", "ready", "deleting", "deleted",
type MediaState string
type MediaStateList []MediaState

func (l MediaStateList) Contains(f MediaState) bool {
	if len(l) == 0 {
		return false
	}
	for _, v := range l {
		if v == f {
			return true
		}
	}
	return false
}

func (l *MediaStateList) Add(f MediaState) {
	for !l.Contains(f) {
		*l = append(*l, f)
	}
}

func (l *MediaStateList) Del(f MediaState) {
	i := -1
	for j, v := range *l {
		if v == f {
			i = j
		}
	}
	if i > -1 {
		(*l)[i] = (*l)[len(*l)-1]
		*l = (*l)[:len(*l)-1]
	}
}

func (l MediaStateList) String() string {
	if len(l) == 0 {
		return ""
	}
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = string(v)
	}
	return strings.Join(s, ",")
}

// MediaType is the list of allowed values for a media's type.
// Allowed values are "unknown", "audiovideo", "audio", "video",
// "subtitle", "geo", "image", "grid", "waveform", "index", "sample",
// "sequence", "text", "url".
type MediaType string
type MediaTypeList []MediaType

func (l MediaTypeList) Contains(f MediaType) bool {
	if len(l) == 0 {
		return false
	}
	for _, v := range l {
		if v == f {
			return true
		}
	}
	return false
}

func (l *MediaTypeList) Add(f MediaType) {
	for !l.Contains(f) {
		*l = append(*l, f)
	}
}

func (l *MediaTypeList) Del(f MediaType) {
	i := -1
	for j, v := range *l {
		if v == f {
			i = j
		}
	}
	if i > -1 {
		(*l)[i] = (*l)[len(*l)-1]
		*l = (*l)[:len(*l)-1]
	}
}

func (l MediaTypeList) String() string {
	if len(l) == 0 {
		return ""
	}
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = string(v)
	}
	return strings.Join(s, ",")
}

// MediaFamily is the list of allowed values for media format families.
// Allowed values are "capture", "post", "vfx", "web", "vod", "cinema",
// "broadcast", "archive", "office".
type MediaFamily string
type MediaFamilyList []MediaFamily

// families are hierarchical, so e.g. 'capture' contains 'capture.arri'
func (a MediaFamily) Contains(b MediaFamily) bool {
	return a == b || !strings.Contains(string(a), ".") && strings.HasPrefix(string(b), string(a)+".")
}

func (l MediaFamilyList) Contains(f MediaFamily) bool {
	if len(l) == 0 {
		return false
	}
	for _, v := range l {
		if v.Contains(f) {
			return true
		}
	}
	return false
}

func (l MediaFamilyList) ContainsStrict(f MediaFamily) bool {
	if len(l) == 0 {
		return false
	}
	for _, v := range l {
		if v == f {
			return true
		}
	}
	return false
}

func (l *MediaFamilyList) Add(f MediaFamily) {
	for !l.ContainsStrict(f) {
		*l = append(*l, f)
	}
}

func (l *MediaFamilyList) Del(f MediaFamily) {
	i := -1
	for j, v := range *l {
		if v == f {
			i = j
		}
	}
	if i > -1 {
		(*l)[i] = (*l)[len(*l)-1]
		*l = (*l)[:len(*l)-1]
	}
}

func (l MediaFamilyList) String() string {
	if len(l) == 0 {
		return ""
	}
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = string(v)
	}
	return strings.Join(s, ",")
}

// MediaRelation is the list of allowed values for media format relation types.
// Allowed values are "source", "inter", "proxy", "master".
type MediaRelation string
type MediaRelationList []MediaRelation

func (l MediaRelationList) Contains(f MediaRelation) bool {
	if len(l) == 0 {
		return false
	}
	for _, v := range l {
		if v == f {
			return true
		}
	}
	return false
}

func (l *MediaRelationList) Add(f MediaRelation) {
	for !l.Contains(f) {
		*l = append(*l, f)
	}
}

func (l *MediaRelationList) Del(f MediaRelation) {
	i := -1
	for j, v := range *l {
		if v == f {
			i = j
		}
	}
	if i > -1 {
		(*l)[i] = (*l)[len(*l)-1]
		*l = (*l)[:len(*l)-1]
	}
}

func (l MediaRelationList) String() string {
	if len(l) == 0 {
		return ""
	}
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = string(v)
	}
	return strings.Join(s, ",")
}

// MediaFormat is the list of allowed values for media format container types.
// Allowed values are "unkown", "webm", "avi", "m2ts", "mov", "mp4", "flac",
// "wav", "avchd", "acvintra", "mxf", "cinemadng", "arriraw", "svg", "json",
// "xml", "jpeg", "png", "r3d", "pdf".
type MediaFormat string
type MediaFormatList []MediaFormat

func (l MediaFormatList) Contains(f MediaFormat) bool {
	if len(l) == 0 {
		return false
	}
	for _, v := range l {
		if v == f {
			return true
		}
	}
	return false
}

func (l *MediaFormatList) Add(f MediaFormat) {
	for !l.Contains(f) {
		*l = append(*l, f)
	}
}

func (l *MediaFormatList) Del(f MediaFormat) {
	i := -1
	for j, v := range *l {
		if v == f {
			i = j
		}
	}
	if i > -1 {
		(*l)[i] = (*l)[len(*l)-1]
		*l = (*l)[:len(*l)-1]
	}
}

func (l MediaFormatList) String() string {
	if len(l) == 0 {
		return ""
	}
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = string(v)
	}
	return strings.Join(s, ",")
}

// MediaRole is the list of allowed values for media roles.
type MediaRole string
type MediaRoleList []MediaRole
type MediaRoleMatch string

// roles are hierarchical, so e.g. 'video' contains 'video.dailies'
func (a MediaRole) Contains(b MediaRole) bool {
	return a == b || strings.HasPrefix(string(b), string(a)+".")
}

func (l MediaRoleList) Contains(r MediaRole) bool {
	if len(l) == 0 {
		return false
	}
	for _, v := range l {
		if v.Contains(r) {
			return true
		}
	}
	return false
}

func (l MediaRoleList) ContainsStrict(r MediaRole) bool {
	if len(l) == 0 {
		return false
	}
	for _, v := range l {
		if v == r {
			return true
		}
	}
	return false
}

func (l *MediaRoleList) Add(r MediaRole) {
	for !l.ContainsStrict(r) {
		*l = append(*l, r)
	}
}

func (l *MediaRoleList) Del(r MediaRole) {
	i := -1
	for j, v := range *l {
		if v == r {
			i = j
		}
	}
	if i > -1 {
		(*l)[i] = (*l)[len(*l)-1]
		*l = (*l)[:len(*l)-1]
	}
}

func (l MediaRoleList) String() string {
	if len(l) == 0 {
		return ""
	}
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = string(v)
	}
	return strings.Join(s, ",")
}

// MediaListKind is the list of allowed values for a media's kind field.
// Allowed values are "all", "own", "online", "offline".
type MediaListKind string

// MediaParams is the set of parameters that can be used to create and
// update media.
//
type MediaParams struct {
	Filename   string         `json:"filename"`
	Size       int64          `json:"size"`
	Type       MediaType      `json:"type"`
	Family     MediaFamily    `json:"family"`
	Format     MediaFormat    `json:"format"`
	Relation   MediaRelation  `json:"relation,omitempty"`
	Role       MediaRole      `json:"role,omitempty"`
	Mimetype   string         `json:"mimetype,omitempty"`
	UUID       string         `json:"uuid,omitempty"`
	Timecode   string         `json:"timecode,omitempty"`
	Duration   time.Duration  `json:"duration,omitempty"`
	Bitrate    int64          `json:"bitrate,omitempty"`
	Profile    string         `json:"profile,omitempty"`
	VolumeId   string         `json:"volumeId,omitempty"`
	Hashes     hash.HashBlock `json:"hashes,omitempty"`
	RecordedAt time.Time      `json:"recordedAt,omitempty"`
	Attr       *MediaAttr     `json:"attr,omitempty"`
	Metadata   *MetaDocument  `json:"meta,omitempty"`
	Embed      ApiEmbedFlags  `json:"embed,omitempty"`
}

// FileInfo represents upload and download file metadata used for direct
// image uploads and media upload into volumes.
type FileInfo struct {
	Filename   string         `json:"filename,omitempty"`
	Role       MediaRole      `json:"role,omitempty"`
	Mimetype   string         `json:"mimetype,omitempty"`
	Size       int64          `json:"size,omitempty"`
	Etag       string         `json:"etag,omitempty"`
	Hashes     hash.HashBlock `json:"hashes,omitempty"`
	UUID       string         `json:"uuid,omitempty"`
	VolumeUUID string         `json:"volumeUuid,omitempty"`
	Url        string         `json:"url,omitempty"`
	Embed      ApiEmbedFlags  `json:"embed,omitempty"`
}

type FileInfoList []*FileInfo

type MediaUploadCompletionParams struct {
	Files FileInfoList  `json:"files"`
	Embed ApiEmbedFlags `json:"embed,omitempty"`
}

// MediaEventType is the list of allowed values for the asset list event field.
// Allowed values are "created", "uploaded", "updated", "recorded"
type MediaListEvent string

// MediaListParams is the set of parameters that can be used when listing media.
type MediaListParams struct {
	ListParams
	WorkspaceId string            `json:"-"`
	AuthorId    string            `json:"authorId,omitempty"`
	States      MediaStateList    `json:"state,omitempty"`
	Types       MediaTypeList     `json:"type,omitempty"`
	Formats     MediaFormatList   `json:"format,omitempty"`
	Families    MediaFamilyList   `json:"family,omitempty"`
	Roles       MediaRoleList     `json:"role,omitempty"`
	Relations   MediaRelationList `json:"relation,omitempty"`
	Kind        MediaListKind     `json:"kind,omitempty"`
	UUID        string            `json:"uuid,omitempty"`
	Embed       ApiEmbedFlags     `json:"embed,omitempty"`
}

// MediaEmbed is a short representation of the Trimmer media resource,
// used when media is embedded into other resources.
type MediaEmbed struct {
	ID        string            `json:"mediaId"`
	Type      MediaType         `json:"type"`
	Family    MediaFamily       `json:"family"`
	Format    MediaFormat       `json:"format"`
	Relation  MediaRelation     `json:"relation"`
	Mime      string            `json:"mime"`
	Role      MediaRole         `json:"role"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
	ExpiresAt time.Time         `json:"expiresAt"`
	Urls      map[string]string `json:"urls"`
}

// Media is the resource representing a Trimmer media.
type Media struct {
	ID          string         `json:"mediaId"`
	UUID        string         `json:"uuid"`
	State       MediaState     `json:"state"`
	AccountId   string         `json:"accountId"`
	WorkspaceId string         `json:"workspaceId"`
	AuthorId    string         `json:"authorId"`
	Revision    int            `json:"revision"`
	Type        MediaType      `json:"type"`
	Family      MediaFamily    `json:"family"`
	Format      MediaFormat    `json:"format"`
	Role        MediaRole      `json:"role"`
	Mimetype    string         `json:"mimetype"`
	Relation    MediaRelation  `json:"relation"`
	Profile     string         `json:"profile"`
	Timecode    string         `json:"timecode"`
	Duration    time.Duration  `json:"duration"`
	Bitrate     int64          `json:"bitrate"`
	Filename    string         `json:"filename"`
	Size        int64          `json:"size"`
	Hashes      hash.HashBlock `json:"hashes"`
	RecordedAt  time.Time      `json:"recordedAt"`
	UploadedAt  time.Time      `json:"uploadedAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	CreatedAt   time.Time      `json:"createdAt"`
	ExpiresAt   time.Time      `json:"expiresAt"`
	Attr        *MediaAttr     `json:"attr"`
	Metadata    *MetaDocument  `json:"meta"`
	Url         string         `json:"url"`
	JobId       string         `json:"jobId"`
	Workspace   *Workspace     `json:"workspace"`
	Account     *User          `json:"account"`
	Author      *User          `json:"author"`
}

// MediaList is representing a slice of Media structs.
type MediaList []*Media

func (l MediaList) SearchId(id string) (int, *Media) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}

type IsoLanguage string
type PixelFormat string
type ColorModel string
type ColorRange string
type ColorPrimaries string
type ColorTransfer string
type ColorLocation string
type ColorSampling string
type VideoCodec string
type AudioLayout string
type AudioCodec string
type SubtitleCodec string

type MediaAttr struct {
	References []*MediaReference `json:"refs,omitempty"`
	TrackIndex []*MediaTrack     `json:"tracks,omitempty"`
	Video      []*VideoAttr      `json:"video,omitempty"`
	Audio      []*AudioAttr      `json:"audio,omitempty"`
	Sequence   []*SequenceAttr   `json:"sequence,omitempty"`
	Image      []*ImageAttr      `json:"image,omitempty"`
	Grid       []*GridAttr       `json:"grid,omitempty"`
	Subtitle   []*SubtitleAttr   `json:"subtitle,omitempty"`
	Data       []*DataAttr       `json:"data,omitempty"`
	Document   []*DocumentAttr   `json:"document,omitempty"`
	Manifest   []*ManifestAttr   `json:"manifest,omitempty"`
}

func NewMediaAttr() *MediaAttr {
	return &MediaAttr{
		References: make([]*MediaReference, 0),
		TrackIndex: make([]*MediaTrack, 0),
		Video:      make([]*VideoAttr, 0),
		Audio:      make([]*AudioAttr, 0),
		Sequence:   make([]*SequenceAttr, 0),
		Image:      make([]*ImageAttr, 0),
		Grid:       make([]*GridAttr, 0),
		Subtitle:   make([]*SubtitleAttr, 0),
		Data:       make([]*DataAttr, 0),
		Document:   make([]*DocumentAttr, 0),
		Manifest:   make([]*ManifestAttr, 0),
	}
}

// pointer to source media for bookkeeping relations after edits, muxing, etc
// one reference per track
type MediaReference struct {
	UUID           string        `json:"uuid,omitempty"`       // current track referencing the origin track
	Timecode       string        `json:"startTc,omitempty"`    // current media timecode
	Duration       time.Duration `json:"duration,omitempty"`   // duration [ms] for origin and current
	OriginID       string        `json:"originId,omitempty"`   // origin media id
	OriginUUID     string        `json:"originUuid,omitempty"` // global unique id for origin track
	OriginTimecode string        `json:"originTc,omitempty"`   // source media reference timecode
}

// track index entry
type MediaTrack struct {
	UUID     string    `json:"uuid,omitempty"`     // global unique id for track
	TrackNum int       `json:"trackNum,omitempty"` // track number in multi-track containers
	Type     MediaType `json:"type,omitempty"`     // track type
}

type VideoAttr struct {
	OriginID       string         `json:"originId,omitempty"`       // origin media id
	UUID           string         `json:"uuid,omitempty"`           // global unique id for track
	OriginUUID     string         `json:"originUuid,omitempty"`     // global unique id for origin track
	TrackNum       int            `json:"trackNum,omitempty"`       // track number in multi-track containers
	FrameCount     int64          `json:"frameCount,omitempty"`     // number of frames
	RateNum        int            `json:"rateNum,omitempty"`        // video frame rate numerator
	RateDen        int            `json:"rateDen,omitempty"`        // video frame rate denumerator
	Duration       time.Duration  `json:"duration,omitempty"`       // video stream runtime
	Width          int            `json:"width,omitempty"`          // native width
	Height         int            `json:"height,omitempty"`         // native height
	Depth          int            `json:"depth,omitempty"`          // color depth 8/10/12/14/16/24 bit
	Sar            float32        `json:"sar,omitempty"`            // sample aspect ratio
	Dar            float32        `json:"dar,omitempty"`            // display aspect ratio
	Codec          VideoCodec     `json:"codec,omitempty"`          // codec name
	Profile        string         `json:"profile,omitempty"`        // codec profile/level
	Bitrate        int64          `json:"bitrate,omitempty"`        // encoding bitrate
	PixelFormat    PixelFormat    `json:"pixelFormat,omitempty"`    // video pixel format (RGB, YUV, +chroma subsampling)
	ColorModel     ColorModel     `json:"colorModel,omitempty"`     // color space
	ColorRange     ColorRange     `json:"colorRange,omitempty"`     // color range (full or limited range)
	ColorPrimaries ColorPrimaries `json:"colorPrimaries,omitempty"` // color primaries RGB/XYZ mapping
	ColorTransfer  ColorTransfer  `json:"colorTransfer,omitempty"`  // color transfer (linearization, gamma)
	ColorLocation  ColorLocation  `json:"colorLocation,omitempty"`  // color location (chroma subsample position)
	Hashes         hash.HashBlock `json:"hashes,omitempty"`         // content checksums
	Size           int64          `json:"size,omitempty"`           // file size
	Filename       string         `json:"filename,omitempty"`       // media filename, e.g. for DASH
	Url            string         `json:"url,omitempty"`            // dynamic access url
}

// HLS, DASH, DPX, JPEG2000, etc = manifest + sequence(s)
//
type SequenceFile struct {
	Frame    int64          `json:"frame,omitempty"`    // frame number
	Size     int64          `json:"size,omitempty"`     // file size
	Hashes   hash.HashBlock `json:"hashes,omitempty"`   // content checksums
	Filename string         `json:"filename,omitempty"` // original or generated filename
	UUID     string         `json:"uuid,omitempty"`     // global unique id for media file
	Url      string         `json:"url,omitempty"`      // dynamic access url
}

type SequenceFileList []*SequenceFile

func (x SequenceFileList) Size() (size int64) {
	for _, v := range x {
		size += v.Size
	}
	return
}

type SequenceAttr struct {
	OriginID       string           `json:"originId,omitempty"`       // origin media id
	UUID           string           `json:"uuid,omitempty"`           // global unique id for track
	OriginUUID     string           `json:"originUuid,omitempty"`     // global unique id for origin track
	TrackNum       int              `json:"trackNum,omitempty"`       // track number in multi-track containers
	FrameCount     int64            `json:"frameCount,omitempty"`     // number of frames
	RateNum        int              `json:"rateNum,omitempty"`        // video frame rate numerator
	RateDen        int              `json:"rateDen,omitempty"`        // video frame rate denumerator
	Duration       time.Duration    `json:"duration,omitempty"`       // video stream runtime
	Width          int              `json:"width,omitempty"`          // native width
	Height         int              `json:"height,omitempty"`         // native height
	Depth          int              `json:"depth,omitempty"`          // color depth 8/10/12/14/16/24 bit
	Sar            float32          `json:"sar,omitempty"`            // sample aspect ratio
	Dar            float32          `json:"dar,omitempty"`            // display aspect ratio
	Codec          VideoCodec       `json:"codec,omitempty"`          // codec name
	Profile        string           `json:"profile,omitempty"`        // codec profile/level
	Bitrate        int64            `json:"bitrate,omitempty"`        // encoding bitrate
	PixelFormat    PixelFormat      `json:"pixelFormat,omitempty"`    // video pixel format (RGB, YUV, +chroma subsampling)
	ColorModel     ColorModel       `json:"colorModel,omitempty"`     // color space
	ColorRange     ColorRange       `json:"colorRange,omitempty"`     // color range (full or limited range)
	ColorPrimaries ColorPrimaries   `json:"colorPrimaries,omitempty"` // color primaries RGB/XYZ mapping
	ColorTransfer  ColorTransfer    `json:"colorTransfer,omitempty"`  // color transfer (linearization, gamma)
	ColorLocation  ColorLocation    `json:"colorLocation,omitempty"`  // color location (chroma subsample position)
	MediaList      SequenceFileList `json:"media,omitempty"`          // list of media file details
}

type AudioAttr struct {
	OriginID      string         `json:"originId,omitempty"`      // origin media (used by proxies)
	UUID          string         `json:"uuid,omitempty"`          // global unique id for track
	OriginUUID    string         `json:"originUuid,omitempty"`    // global unique id for origin track
	TrackNum      int            `json:"trackNum,omitempty"`      // track number in multi-track containers
	Lang          IsoLanguage    `json:"lang,omitempty"`          // ISO language code
	Channels      int            `json:"channels,omitempty"`      // audio channels
	ChannelLayout AudioLayout    `json:"channelLayout,omitempty"` // audio channel layout
	SampleRate    int            `json:"sampleRate,omitempty"`    // audio sampling rate (44100, 48000, 96000, ...)
	SampleBits    int            `json:"sampleBits,omitempty"`    // sample width 8/16/24/32 bit
	SampleFormat  string         `json:"sampleFormat,omitempty"`  // sample format
	Codec         AudioCodec     `json:"codec,omitempty"`         // codec name
	Profile       string         `json:"profile,omitempty"`       // codec profile/level
	Bitrate       int64          `json:"bitrate,omitempty"`       // audio encoding bitrate
	Duration      time.Duration  `json:"duration,omitempty"`      // runtime in ms
	Hashes        hash.HashBlock `json:"hashes,omitempty"`        // content checksums
	Filename      string         `json:"filename,omitempty"`      // media filename (e.g for DASH)
	Size          int64          `json:"size,omitempty"`          // file size
	Url           string         `json:"url,omitempty"`           // dynamic access url
}

// single and multi-resolution image media
type ImageAttr struct {
	OriginID    string         `json:"originId,omitempty"`    // origin media (used by proxies)
	UUID        string         `json:"uuid,omitempty"`        // global unique id for track
	OriginUUID  string         `json:"originUuid,omitempty"`  // global unique id for origin track
	TrackNum    int            `json:"trackNum,omitempty"`    // track number in multi-track containers
	Width       int            `json:"width,omitempty"`       // native width
	Height      int            `json:"height,omitempty"`      // native height
	Depth       int            `json:"depth,omitempty"`       // color depth 8/10/12/14/16/24 bit
	Sar         float32        `json:"sar,omitempty"`         // sample aspect ratio
	Dar         float32        `json:"dar,omitempty"`         // display aspect ratio
	PixelFormat PixelFormat    `json:"pixelFormat,omitempty"` // image pixel format (RGB24)
	ColorModel  ColorModel     `json:"colorModel,omitempty"`  // image color space (sRGB)
	Hashes      hash.HashBlock `json:"hashes,omitempty"`      // content checksums
	Size        int64          `json:"size,omitempty"`        // file size
	Filename    string         `json:"filename,omitempty"`    // media filename
	Url         string         `json:"url,omitempty"`         // dynamic access url
}

// grid of images for quick thumbnail skimming
type GridAttr struct {
	OriginID    string         `json:"originId,omitempty"`    // origin media (used by proxies)
	UUID        string         `json:"uuid,omitempty"`        // global unique id for track
	OriginUUID  string         `json:"originUuid,omitempty"`  // global unique id for origin track
	TrackNum    int            `json:"trackNum,omitempty"`    // track number in multi-track containers
	FrameCount  int            `json:"frameCount,omitempty"`  // number of frames
	GridX       int            `json:"gridX,omitempty"`       // horizontal grid fields
	GridY       int            `json:"gridY,omitempty"`       // vertical grid fields
	Width       int            `json:"width,omitempty"`       // frame/image width
	Height      int            `json:"height,omitempty"`      // frame/image height
	Depth       int            `json:"depth,omitempty"`       // color depth 8/10/12/14/16/24 bit
	PixelFormat PixelFormat    `json:"pixelFormat,omitempty"` // image pixel format (RGB)
	ColorModel  ColorModel     `json:"colorModel,omitempty"`  // image color space (sRGB)
	Sar         float32        `json:"sar,omitempty"`         // sample aspect ratio
	Dar         float32        `json:"dar,omitempty"`         // display aspect ratio
	RateNum     int            `json:"rateNum,omitempty"`     // frame rate numerator
	RateDen     int            `json:"rateDen,omitempty"`     // frame rate denumerator
	Duration    time.Duration  `json:"duration,omitempty"`    // grid runtime
	Hashes      hash.HashBlock `json:"hashes,omitempty"`      // content checksums
	Size        int64          `json:"size,omitempty"`        // file size
	Filename    string         `json:"filename,omitempty"`    // media filename
	Url         string         `json:"url,omitempty"`         // dynamic access url
}

type SubtitleAttr struct {
	OriginID   string         `json:"originId,omitempty"`   // origin media (used by proxies)
	UUID       string         `json:"uuid,omitempty"`       // global unique id for track
	OriginUUID string         `json:"originUuid,omitempty"` // global unique id for origin track
	TrackNum   int            `json:"trackNum,omitempty"`   // track number in multi-track containers
	Duration   time.Duration  `json:"duration,omitempty"`   // subtitle stream runtime
	Codec      SubtitleCodec  `json:"codec,omitempty"`      // codec name
	Lang       IsoLanguage    `json:"lang,omitempty"`       // ISO language code
	Hashes     hash.HashBlock `json:"hashes,omitempty"`     // content checksums
	Filename   string         `json:"filename,omitempty"`   // media filename, e.g. for DASH
	Size       int64          `json:"size,omitempty"`       // file size
	Url        string         `json:"url,omitempty"`        // dynamic access url
}

type DocumentAttr struct {
	OriginID   string         `json:"originId,omitempty"`   // origin media (used by proxies)
	UUID       string         `json:"uuid,omitempty"`       // global unique id for track
	OriginUUID string         `json:"originUuid,omitempty"` // global unique id for origin track
	TrackNum   int            `json:"trackNum,omitempty"`   // track number in multi-track containers
	Pages      int            `json:"pages,omitempty"`      // page number
	Lang       IsoLanguage    `json:"lang,omitempty"`       // ISO language code
	DataType   string         `json:"dataType,omitempty"`   // doc datatype
	Hashes     hash.HashBlock `json:"hashes,omitempty"`     // content checksums
	Filename   string         `json:"filename,omitempty"`   // media filename, e.g. for sidecar data
	Size       int64          `json:"size,omitempty"`       // file size
	Url        string         `json:"url,omitempty"`        // dynamic access url
}

type DataAttr struct {
	OriginID   string         `json:"originId,omitempty"`   // origin media (used by proxies)
	UUID       string         `json:"uuid,omitempty"`       // global unique id for track
	OriginUUID string         `json:"originUuid,omitempty"` // global unique id for origin track
	TrackNum   int            `json:"trackNum,omitempty"`   // track number in multi-track containers
	Duration   time.Duration  `json:"duration,omitempty"`   // data stream runtime
	Codec      string         `json:"codec,omitempty"`      // codec name
	DataType   string         `json:"dataType,omitempty"`   // track datatype
	Hashes     hash.HashBlock `json:"hashes,omitempty"`     // content checksums
	Filename   string         `json:"filename,omitempty"`   // media filename, e.g. for sidecar data
	Size       int64          `json:"size,omitempty"`       // file size
	Url        string         `json:"url,omitempty"`        // dynamic access url
}

type ManifestAttr struct {
	OriginID   string `json:"originId,omitempty"`   // origin media (used by proxies)
	UUID       string `json:"uuid,omitempty"`       // global unique id for track
	OriginUUID string `json:"originUuid,omitempty"` // global unique id for origin track
	TrackNum   int    `json:"trackNum,omitempty"`   // track number in multi-track containers
	// only used for embedded manifests (i.e. a manifest is part of a media file collection
	// such as with DASH or HLS, but not when the manifest file is independent, such as
	// for IMF)
	Duration time.Duration  `json:"duration,omitempty"` // track runtime
	Format   MediaFormat    `json:"format,omitempty"`   // storage container: mpd, m3u8, imf, xml
	Mime     string         `json:"mimetype,omitempty"` // mime type
	Hashes   hash.HashBlock `json:"hashes,omitempty"`   // content checksums
	Filename string         `json:"filename,omitempty"` // media filename
	Size     int64          `json:"size,omitempty"`     // file size
	Url      string         `json:"url,omitempty"`      // dynamic access url
}
