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
	"strings"

	trimmer "trimmer.io/go-trimmer"
)

const (
	MediaStateUndefined  trimmer.MediaState = ""
	MediaStateAll        trimmer.MediaState = "all" // used for listing only
	MediaStateCreated    trimmer.MediaState = "created"
	MediaStateUploading  trimmer.MediaState = "uploading"
	MediaStateUploaded   trimmer.MediaState = "uploaded"
	MediaStateAnalyzing  trimmer.MediaState = "analyzing"
	MediaStateProcessing trimmer.MediaState = "processing"
	MediaStateFailed     trimmer.MediaState = "failed"
	MediaStateReady      trimmer.MediaState = "ready"
	MediaStateDeleting   trimmer.MediaState = "deleting"
	MediaStateDeleted    trimmer.MediaState = "deleted"
	MediaStateMissing    trimmer.MediaState = "missing"
)

func ParseState(s string) trimmer.MediaState {
	switch s {
	case "all":
		return MediaStateAll
	case "created":
		return MediaStateCreated
	case "uploading":
		return MediaStateUploading
	case "uploaded":
		return MediaStateUploaded
	case "analyzing":
		return MediaStateAnalyzing
	case "processing":
		return MediaStateProcessing
	case "failed":
		return MediaStateFailed
	case "ready":
		return MediaStateReady
	case "deleting":
		return MediaStateDeleting
	case "deleted":
		return MediaStateDeleted
	case "missing":
		return MediaStateMissing
	default:
		return MediaStateUndefined
	}
}

// MediaListKind
const (
	MediaListKindUndefined trimmer.MediaListKind = ""
	MediaListKindAll       trimmer.MediaListKind = "all"
	MediaListKindOwn       trimmer.MediaListKind = "own"
	MediaListKindOnline    trimmer.MediaListKind = "online"
	MediaListKindOffline   trimmer.MediaListKind = "offline"
)

func ParseListKind(s string) trimmer.MediaListKind {
	switch s {
	case "all":
		return MediaListKindAll
	case "own":
		return MediaListKindOwn
	case "online":
		return MediaListKindOnline
	case "offline":
		return MediaListKindOffline
	default:
		return MediaListKindUndefined
	}
}

// MediaListEvent
const (
	MediaListEventUndefined trimmer.MediaListEvent = ""
	MediaListEventCreated   trimmer.MediaListEvent = "created"
	MediaListEventUpdated   trimmer.MediaListEvent = "updated"
	MediaListEventUploaded  trimmer.MediaListEvent = "uploaded"
	MediaListEventRecorded  trimmer.MediaListEvent = "recorded"
)

func ParseListEvent(s string) trimmer.MediaListEvent {
	switch s {
	case "created":
		return MediaListEventCreated
	case "updated":
		return MediaListEventUpdated
	case "uploaded":
		return MediaListEventUploaded
	case "recorded":
		return MediaListEventRecorded
	default:
		return MediaListEventUndefined
	}
}

const (
	MediaTypeUndefined  trimmer.MediaType = ""
	MediaTypeData       trimmer.MediaType = "data"
	MediaTypeAudioVideo trimmer.MediaType = "audiovideo"
	MediaTypeAudio      trimmer.MediaType = "audio"
	MediaTypeVideo      trimmer.MediaType = "video"
	MediaTypeSequence   trimmer.MediaType = "sequence"
	MediaTypeSubtitle   trimmer.MediaType = "subtitle"
	MediaTypeImage      trimmer.MediaType = "image"
	MediaTypeGrid       trimmer.MediaType = "grid"
	MediaTypeManifest   trimmer.MediaType = "manifest"
	MediaTypeEditlist   trimmer.MediaType = "editlist"
	MediaTypeDocument   trimmer.MediaType = "document"
	MediaTypeSidecar    trimmer.MediaType = "sidecar"
	MediaTypeScript     trimmer.MediaType = "script"
	MediaTypeEmbed      trimmer.MediaType = "embed"
)

func ParseMediaType(s string) trimmer.MediaType {
	switch s {
	case "data":
		return MediaTypeData
	case "audiovideo":
		return MediaTypeAudioVideo
	case "audio":
		return MediaTypeAudio
	case "video":
		return MediaTypeVideo
	case "sequence":
		return MediaTypeSequence
	case "subtitle":
		return MediaTypeSubtitle
	case "image":
		return MediaTypeImage
	case "grid":
		return MediaTypeGrid
	case "manifest":
		return MediaTypeManifest
	case "editlist":
		return MediaTypeEditlist
	case "document":
		return MediaTypeDocument
	case "sidecar":
		return MediaTypeSidecar
	case "script":
		return MediaTypeScript
	case "embed":
		return MediaTypeEmbed
	default:
		return MediaTypeUndefined
	}
}

func IsValidType(x trimmer.MediaType) bool {
	return x != MediaTypeUndefined
}

func IsMultiFileMediaType(t trimmer.MediaType) bool {
	switch t {
	case MediaTypeSequence, MediaTypeImage, MediaTypeGrid:
		return true
	default:
		return false
	}
}

func IsVideoMediaType(t trimmer.MediaType) bool {
	switch t {
	case MediaTypeAudioVideo, MediaTypeVideo, MediaTypeSequence, MediaTypeEditlist:
		return true
	default:
		return false
	}
}

func IsAudioMediaType(t trimmer.MediaType) bool {
	switch t {
	case MediaTypeAudioVideo, MediaTypeAudio:
		return true
	default:
		return false
	}
}

func IsImageMediaType(t trimmer.MediaType) bool {
	switch t {
	case MediaTypeImage, MediaTypeGrid:
		return true
	default:
		return false
	}
}

const (
	MediaFamilyUndefined         trimmer.MediaFamily = ""
	MediaFamilyCapture           trimmer.MediaFamily = "capture"
	MediaFamilyCaptureAja        trimmer.MediaFamily = "capture.aja"
	MediaFamilyCaptureArri       trimmer.MediaFamily = "capture.arri"
	MediaFamilyCaptureBlackmagic trimmer.MediaFamily = "capture.blackmagic"
	MediaFamilyCaptureCanon      trimmer.MediaFamily = "capture.canon"
	MediaFamilyCaptureGopro      trimmer.MediaFamily = "capture.gopro"
	MediaFamilyCapturePanasonic  trimmer.MediaFamily = "capture.panasonic"
	MediaFamilyCapturePanavision trimmer.MediaFamily = "capture.panavision"
	MediaFamilyCaptureSony       trimmer.MediaFamily = "capture.sony"
	MediaFamilyCaptureRed        trimmer.MediaFamily = "capture.red"
	MediaFamilyPost              trimmer.MediaFamily = "post"
	MediaFamilyPostProres        trimmer.MediaFamily = "post.prores"
	MediaFamilyPostAVCUltra      trimmer.MediaFamily = "post.avcultra"
	MediaFamilyVfx               trimmer.MediaFamily = "vfx"
	MediaFamilyWeb               trimmer.MediaFamily = "web"
	MediaFamilyWebImage          trimmer.MediaFamily = "web.image"
	MediaFamilyWebHtml5          trimmer.MediaFamily = "web.html5"
	MediaFamilyWebHls            trimmer.MediaFamily = "web.hls"
	MediaFamilyWebDash           trimmer.MediaFamily = "web.dash"
	MediaFamilyVod               trimmer.MediaFamily = "vod"
	MediaFamilyVodVimeo          trimmer.MediaFamily = "vod.vimeo"
	MediaFamilyVodYoutube        trimmer.MediaFamily = "vod.youtube"
	MediaFamilyVodNetflix        trimmer.MediaFamily = "vod.netflix"
	MediaFamilyVodItunes         trimmer.MediaFamily = "vod.itunes"
	MediaFamilyCinema            trimmer.MediaFamily = "cinema"
	MediaFamilyBroadcast         trimmer.MediaFamily = "broadcast"
	MediaFamilyArchive           trimmer.MediaFamily = "archive"
	MediaFamilyOffice            trimmer.MediaFamily = "office"
	MediaFamilyOfficeDocs        trimmer.MediaFamily = "office.docs"
	MediaFamilyOfficeSheets      trimmer.MediaFamily = "office.sheets"
	MediaFamilyOfficeSlides      trimmer.MediaFamily = "office.slides"
)

func ParseMediaFamily(s string) trimmer.MediaFamily {
	switch s {
	case "capture":
		return MediaFamilyCapture
	case "capture.aja":
		return MediaFamilyCaptureAja
	case "capture.arri":
		return MediaFamilyCaptureArri
	case "capture.blackmagic":
		return MediaFamilyCaptureBlackmagic
	case "capture.canon":
		return MediaFamilyCaptureCanon
	case "capture.gopro":
		return MediaFamilyCaptureGopro
	case "capture.panasonic":
		return MediaFamilyCapturePanasonic
	case "capture.panavision":
		return MediaFamilyCapturePanavision
	case "capture.sony":
		return MediaFamilyCaptureSony
	case "capture.red":
		return MediaFamilyCaptureRed
	case "post":
		return MediaFamilyPost
	case "post.prores":
		return MediaFamilyPostProres
	case "post.avcultra":
		return MediaFamilyPostAVCUltra
	case "vfx":
		return MediaFamilyVfx
	case "web":
		return MediaFamilyWeb
	case "web.image":
		return MediaFamilyWebImage
	case "web.html5":
		return MediaFamilyWebHtml5
	case "web.hls":
		return MediaFamilyWebHls
	case "web.dash":
		return MediaFamilyWebDash
	case "vod":
		return MediaFamilyVod
	case "vod.vimeo":
		return MediaFamilyVodVimeo
	case "vod.youtube":
		return MediaFamilyVodYoutube
	case "vod.netflix":
		return MediaFamilyVodNetflix
	case "vod.itunes":
		return MediaFamilyVodItunes
	case "cinema":
		return MediaFamilyCinema
	case "broadcast":
		return MediaFamilyBroadcast
	case "archive":
		return MediaFamilyArchive
	case "office":
		return MediaFamilyOffice
	case "office.docs":
		return MediaFamilyOfficeDocs
	case "office.sheets":
		return MediaFamilyOfficeSheets
	case "office.slides":
		return MediaFamilyOfficeSlides
	default:
		return MediaFamilyUndefined
	}
}

func IsValidFamily(x trimmer.MediaFamily) bool {
	return x != MediaFamilyUndefined
}

const (
	MediaRelationUndefined trimmer.MediaRelation = ""
	MediaRelationSource    trimmer.MediaRelation = "source"
	MediaRelationInter     trimmer.MediaRelation = "inter"
	MediaRelationProxy     trimmer.MediaRelation = "proxy"
	MediaRelationMaster    trimmer.MediaRelation = "master"
)

func ParseMediaRelation(s string) trimmer.MediaRelation {
	switch s {
	case "source":
		return MediaRelationSource
	case "inter":
		return MediaRelationInter
	case "proxy":
		return MediaRelationProxy
	case "master":
		return MediaRelationMaster
	default:
		return MediaRelationUndefined
	}
}

func IsValidRelation(x trimmer.MediaRelation) bool {
	return x != MediaRelationUndefined
}

const (
	MediaFormatUndefined trimmer.MediaFormat = ""
	MediaFormatAAC       trimmer.MediaFormat = "aac"
	MediaFormatArriLook  trimmer.MediaFormat = "aml"
	MediaFormatArriRaw   trimmer.MediaFormat = "ari"
	MediaFormatAVI       trimmer.MediaFormat = "avi"
	MediaFormatBWF       trimmer.MediaFormat = "bwf"
	MediaFormatCDL       trimmer.MediaFormat = "cdl"
	MediaFormatCine      trimmer.MediaFormat = "cine"
	MediaFormatCSV       trimmer.MediaFormat = "csv"
	MediaFormatCUBE      trimmer.MediaFormat = "cube"
	MediaFormatDATA      trimmer.MediaFormat = "data"
	MediaFormatDNG       trimmer.MediaFormat = "dng"
	MediaFormatDPX       trimmer.MediaFormat = "dpx"
	MediaFormatEXR       trimmer.MediaFormat = "exr"
	MediaFormatFLAC      trimmer.MediaFormat = "flac"
	MediaFormatJPEG      trimmer.MediaFormat = "jpeg"
	MediaFormatJPEG2000  trimmer.MediaFormat = "jp2"
	MediaFormatJPEG2000X trimmer.MediaFormat = "jpx"
	MediaFormatJSON      trimmer.MediaFormat = "json"
	MediaFormatLOOK      trimmer.MediaFormat = "look"
	MediaFormatLUT       trimmer.MediaFormat = "lut"
	MediaFormatM2TS      trimmer.MediaFormat = "m2ts"
	MediaFormatM3U       trimmer.MediaFormat = "m3u"
	MediaFormatM3U8      trimmer.MediaFormat = "m3u8"
	MediaFormatMKV       trimmer.MediaFormat = "mkv"
	MediaFormatMOV       trimmer.MediaFormat = "mov"
	MediaFormatMP3       trimmer.MediaFormat = "mp3"
	MediaFormatMP4       trimmer.MediaFormat = "mp4"
	MediaFormatMPD       trimmer.MediaFormat = "mpd"
	MediaFormatMXF       trimmer.MediaFormat = "mxf"
	MediaFormatOGG       trimmer.MediaFormat = "ogg"
	MediaFormatPDF       trimmer.MediaFormat = "pdf"
	MediaFormatPNG       trimmer.MediaFormat = "png"
	MediaFormatPSD       trimmer.MediaFormat = "psd"
	MediaFormatR3D       trimmer.MediaFormat = "r3d"
	MediaFormatRMD       trimmer.MediaFormat = "rmd"
	MediaFormatSRT       trimmer.MediaFormat = "srt"
	MediaFormatSVG       trimmer.MediaFormat = "svg"
	MediaFormatTIFF      trimmer.MediaFormat = "tiff"
	MediaFormatTXT       trimmer.MediaFormat = "txt"
	MediaFormatWebVTT    trimmer.MediaFormat = "vtt"
	MediaFormatWAV       trimmer.MediaFormat = "wav"
	MediaFormatWEBM      trimmer.MediaFormat = "webm"
	MediaFormatXML       trimmer.MediaFormat = "xml"
	MediaFormatXMP       trimmer.MediaFormat = "xmp"
)

func IsVideoMediaFormat(f trimmer.MediaFormat) bool {
	switch f {
	case
		MediaFormatAVI,
		MediaFormatCine,
		MediaFormatM2TS,
		MediaFormatMKV,
		MediaFormatMOV,
		MediaFormatMP4,
		MediaFormatMXF,
		MediaFormatR3D,
		MediaFormatWEBM:
		return true
	default:
		return false
	}
}

func IsSequenceMediaFormat(f trimmer.MediaFormat) bool {
	switch f {
	case
		MediaFormatArriRaw,
		MediaFormatDNG,
		MediaFormatDPX,
		MediaFormatEXR,
		MediaFormatJPEG,
		MediaFormatJPEG2000,
		MediaFormatJPEG2000X,
		MediaFormatTIFF:
		return true
	default:
		return false
	}
}

func IsAudioMediaFormat(f trimmer.MediaFormat) bool {
	switch f {
	case
		MediaFormatAAC,
		MediaFormatBWF,
		MediaFormatFLAC,
		MediaFormatMP3,
		MediaFormatOGG,
		MediaFormatWAV:
		return true
	default:
		return false
	}
}

func IsImageMediaFormat(f trimmer.MediaFormat) bool {
	switch f {
	case
		MediaFormatJPEG,
		MediaFormatPNG,
		MediaFormatPSD,
		MediaFormatSVG,
		MediaFormatTIFF:
		return true
	default:
		return false
	}
}

// Note: catches ambiguities in mime-type/format detection systems
//       and is compatible to input from filepath.Ext(filename)
//
func ParseMediaFormat(s string) trimmer.MediaFormat {
	switch strings.ToLower(strings.Trim(s, ".")) {
	case "aac":
		return MediaFormatAAC
	case "aml":
		return MediaFormatArriLook
	case "ari":
		return MediaFormatArriRaw
	case "avi":
		return MediaFormatAVI
	case "bwf":
		return MediaFormatBWF
	case "cdl":
		return MediaFormatCDL
	case "cine":
		return MediaFormatCine
	case "csv":
		return MediaFormatCSV
	case "cube":
		return MediaFormatCUBE
	case "dat", "data":
		return MediaFormatDATA
	case "dng":
		return MediaFormatDNG
	case "dpx":
		return MediaFormatDPX
	case "exr":
		return MediaFormatEXR
	case "flac":
		return MediaFormatFLAC
	case "jpg", "jpeg":
		return MediaFormatJPEG
	case "jp2", "j2k":
		return MediaFormatJPEG2000
	case "jpx":
		return MediaFormatJPEG2000X
	case "json":
		return MediaFormatJSON
	case "look":
		return MediaFormatLOOK
	case "lut":
		return MediaFormatLUT
	case "m2ts":
		return MediaFormatM2TS
	case "m3u":
		return MediaFormatM3U
	case "m3u8":
		return MediaFormatM3U8
	case "mkv":
		return MediaFormatMKV
	case "mov", "quicktime":
		return MediaFormatMOV
	case "mp3":
		return MediaFormatMP3
	case "mp4", "m4v", "m4a":
		return MediaFormatMP4
	case "mpd":
		return MediaFormatMPD
	case "mxf":
		return MediaFormatMXF
	case "ogg":
		return MediaFormatOGG
	case "pdf":
		return MediaFormatPDF
	case "png":
		return MediaFormatPNG
	case "psd":
		return MediaFormatPSD
	case "r3d":
		return MediaFormatR3D
	case "rmd":
		return MediaFormatRMD
	case "srt":
		return MediaFormatSRT
	case "svg":
		return MediaFormatSVG
	case "tif", "tiff":
		return MediaFormatTIFF
	case "txt":
		return MediaFormatTXT
	case "vtt":
		return MediaFormatWebVTT
	case "wav":
		return MediaFormatWAV
	case "webm":
		return MediaFormatWEBM
	case "xml":
		return MediaFormatXML
	case "xmp":
		return MediaFormatXMP
	default:
		return MediaFormatUndefined
	}
}

func IsValidFormat(x trimmer.MediaFormat) bool {
	return x != MediaFormatUndefined
}

const (
	MediaRoleUndefined trimmer.MediaRole = ""
	// top-level roles
	MediaRoleImage   trimmer.MediaRole = "image"
	MediaRoleVideo   trimmer.MediaRole = "video"
	MediaRoleAudio   trimmer.MediaRole = "audio"
	MediaRoleScript  trimmer.MediaRole = "script"
	MediaRoleWeb     trimmer.MediaRole = "web"
	MediaRoleOffice  trimmer.MediaRole = "office"
	MediaRoleSidecar trimmer.MediaRole = "sidecar"
	// pre-defined video sub-roles
	MediaRoleDailies   trimmer.MediaRole = "video.dailies"
	MediaRoleEditorial trimmer.MediaRole = "video.editorial"
	MediaRoleVfx       trimmer.MediaRole = "video.vfx"
	MediaRoleTitles    trimmer.MediaRole = "video.titles"
	// pre-defined audio sub-roles
	MediaRoleDialogue trimmer.MediaRole = "audio.dialogue"
	MediaRoleMusic    trimmer.MediaRole = "audio.music"
	MediaRoleFoley    trimmer.MediaRole = "audio.foley"
	MediaRoleEffects  trimmer.MediaRole = "audio.effects"
	// pre-defined script sub-roles
	MediaRoleDraft trimmer.MediaRole = "script.draft"
	MediaRoleFinal trimmer.MediaRole = "script.final"
	MediaRoleNotes trimmer.MediaRole = "script.notes"
	// pre-defined web sub-roles
	MediaRolePoster     trimmer.MediaRole = "web.poster"
	MediaRoleStoryboard trimmer.MediaRole = "web.storyboard"
	MediaRoleThumbnail  trimmer.MediaRole = "web.thumbnail"
	MediaRoleAvatar     trimmer.MediaRole = "web.avatar"
	MediaRoleBanner     trimmer.MediaRole = "web.banner"
	MediaRoleLogo       trimmer.MediaRole = "web.logo"
	MediaRoleBackground trimmer.MediaRole = "web.background"
	MediaRoleWatermark  trimmer.MediaRole = "web.watermark"
	// pre-defined office roles
	MediaRoleOfficeAccounting trimmer.MediaRole = "office.accounting"
	MediaRoleOfficeLegal      trimmer.MediaRole = "office.legal"
	MediaRoleOfficeReport     trimmer.MediaRole = "office.report"
)

func ParseMediaRole(s string) trimmer.MediaRole {
	switch s {
	case "image":
		return MediaRoleImage
	case "video":
		return MediaRoleVideo
	case "audio":
		return MediaRoleAudio
	case "script":
		return MediaRoleScript
	case "web":
		return MediaRoleWeb
	case "office":
		return MediaRoleOffice
	case "sidecar":
		return MediaRoleSidecar
	case "video.dailies":
		return MediaRoleDailies
	case "video.editorial":
		return MediaRoleEditorial
	case "video.vfx":
		return MediaRoleVfx
	case "video.titles":
		return MediaRoleTitles
	case "audio.dialogue":
		return MediaRoleDialogue
	case "audio.music":
		return MediaRoleMusic
	case "audio.foley":
		return MediaRoleFoley
	case "audio.effects":
		return MediaRoleEffects
	case "script.draft":
		return MediaRoleDraft
	case "script.final":
		return MediaRoleFinal
	case "script.notes":
		return MediaRoleNotes
	case "web.poster":
		return MediaRolePoster
	case "web.storyboard":
		return MediaRoleStoryboard
	case "web.thumbnail":
		return MediaRoleThumbnail
	case "web.avatar":
		return MediaRoleAvatar
	case "web.banner":
		return MediaRoleBanner
	case "web.logo":
		return MediaRoleLogo
	case "web.background":
		return MediaRoleBackground
	case "web.watermark":
		return MediaRoleWatermark
	case "office.accounting":
		return MediaRoleOfficeAccounting
	case "office.legal":
		return MediaRoleOfficeLegal
	case "office.report":
		return MediaRoleOfficeReport
	default:
		// allow user-defined roles here, so just check the top part
		r := trimmer.MediaRole(s)
		if IsValidRole(r) {
			return r
		}
		return MediaRoleUndefined
	}
}

func IsValidRole(r trimmer.MediaRole) bool {
	p := strings.Split(string(r), ".")
	if len(p) == 1 {
		return r != MediaRoleUndefined
	}
	return len(p) > 0 && len(p) <= 3 && IsValidRole(ParseMediaRole(p[0]))
}

// Image and Video Color spaces (i.e. Matrix Coefficients)
//
const (
	ColorModelUndefined   trimmer.ColorModel = ""
	ColorModelRGB         trimmer.ColorModel = "rgb"         // 0: Identity: RGB, XYZ, IEC 61966-2-1 sRGB, SMPTE ST428-1
	ColorModelBT709       trimmer.ColorModel = "bt709"       // 1: BT.709, BT.1361-0, IEC 61966-2-1 sYCC, IEC 61966-2-4 xvYCC_709, SMPTE RP177 Annex B
	ColorModelUnspecified trimmer.ColorModel = "unspecified" // 2
	ColorModelReserved    trimmer.ColorModel = "reserved"    // 3
	ColorModelFCC         trimmer.ColorModel = "fcc"         // 4: FCC Title 47:2003 73.682 (a) (20)
	ColorModelBT470BG     trimmer.ColorModel = "bt470bg"     // 5: BT.470BG, BT.601 625, BT.1700 PAL/SECAM, IEC 61966-2-4 xvYCC_601
	ColorModelSMPTE170M   trimmer.ColorModel = "smpte170m"   // 6: BT.601 525, BT.1700 NTSC, SMPTE 170M:2004
	ColorModelSMPTE240M   trimmer.ColorModel = "smpte240m"   // 7: SMPTE 240M:1999 (functionally identitcal to smpte170m)
	ColorModelYCGCO       trimmer.ColorModel = "ycgco"       // 8: YCgCo, Used by Dirac / VC-2 and H.264 FRext, see ITU-T SG16
	ColorModelBT2020NCL   trimmer.ColorModel = "bt2020ncl"   // 9: BT.2020 NCL, BT.2100 [HDR-TV] YCbCr (non-constant luminance system)
	ColorModelBT2020CL    trimmer.ColorModel = "bt2020cl"    // 10: BT.2020 CL (constant luminance system)
	ColorModelYDZYX       trimmer.ColorModel = "ydzyx"       // 11: SMPTE ST 2085:2015 YDzDx
	ColorModelDERIVED_NCL trimmer.ColorModel = "derived_ncl" // 12: NCL computed, see ISO/IEC 23001-8:2016
	ColorModelDERIVED_CL  trimmer.ColorModel = "derived_cl"  // 13: NCL computed, see ISO/IEC 23001-8:2016
	ColorModelICTCP       trimmer.ColorModel = "ictcp"       // 14: BT.2100 [HDR-TV] ICtCp with constant luminance system
	ColorModelYCBCR       trimmer.ColorModel = "ycbcr"
	ColorModelXVYCC       trimmer.ColorModel = "xvycc"
	ColorModelSRGB        trimmer.ColorModel = "srgb"
	ColorModelAdobeRGB    trimmer.ColorModel = "adobergb"
	ColorModelHSV         trimmer.ColorModel = "hsv"
	ColorModelHLS         trimmer.ColorModel = "hls"
	ColorModelXYZ         trimmer.ColorModel = "xyz"
)

const (
	ColorPrimariesUndefined   trimmer.ColorPrimaries = ""
	ColorPrimariesReserved0   trimmer.ColorPrimaries = "reserved0"   // 0: reserved
	ColorPrimariesBT709       trimmer.ColorPrimaries = "bt709"       // 1: BT.709, BT.1316-0, IEC 61966-2-1 sRGB/sYCC, IEC 61966-2-4, SMPTE RP 177:1993 Annex B
	ColorPrimariesUnspecified trimmer.ColorPrimaries = "unspecified" // 2: Unspecified
	ColorPrimariesReserved    trimmer.ColorPrimaries = "reserved"    // 3: Reserved
	ColorPrimariesBT470M      trimmer.ColorPrimaries = "bt470m"      // 4: BT.470M
	ColorPrimariesBT470BG     trimmer.ColorPrimaries = "bt470bg"     // 5: BT.470BG, BT.601 625, BT.1700 PAL/SECAM
	ColorPrimariesSMPTE170M   trimmer.ColorPrimaries = "smpte170m"   // 6: BT.601 525, BT.1700 NTSC, SMPTE 170M:2004
	ColorPrimariesSMPTE240M   trimmer.ColorPrimaries = "smpte240m"   // 7: SMPTE 240M:1999
	ColorPrimariesFilm        trimmer.ColorPrimaries = "film"        // 8: Generic Film (Illuminant C)
	ColorPrimariesBT2020      trimmer.ColorPrimaries = "bt2020"      // 9: BT.2020-2, BT.[HDR-TV]
	ColorPrimariesSMPTE428    trimmer.ColorPrimaries = "smpte428"    // 10: SMPTE 428-1 (CIE 1931 XYZ)
	ColorPrimariesP3DCI       trimmer.ColorPrimaries = "p3dci"       // 11: SMPTE RP 431-2:2011 (P3 Theatrical, DCI Whitepoint)
	ColorPrimariesP3D65       trimmer.ColorPrimaries = "p3d65"       // 12: SMPTE EG 432-1:2010 (P3 D65)
	ColorPrimariesP3D60       trimmer.ColorPrimaries = "p3d60"       // ??: SMPTE EG 432-1:2010 (P3 D60)
	ColorPrimariesXYZ         trimmer.ColorPrimaries = "xyz"
	ColorPrimariesJedec22     trimmer.ColorPrimaries = "jedec-p22" // EBU 3213-E (1975), JEDEC P22 Phosphors

	// ACES Primaries
	ColorPrimariesAcesAP0 trimmer.ColorPrimaries = "aces-ap0" // ACES full
	ColorPrimariesAcesAP1 trimmer.ColorPrimaries = "aces-ap1" // ACES limited

	// vendor-specific color primaries
	ColorPrimariesNative          trimmer.ColorPrimaries = "native"          // sensor-native, use when unknown vendor
	ColorPrimariesAdobeRGB        trimmer.ColorPrimaries = "adobergb"        // Adobe
	ColorPrimariesAlexaWideGamut  trimmer.ColorPrimaries = "logcwgam"        // Arri
	ColorPrimariesVgamut          trimmer.ColorPrimaries = "vgamut"          // Panasonic
	ColorPrimariesSgamut          trimmer.ColorPrimaries = "sgamut"          // Sony
	ColorPrimariesSgamut2         trimmer.ColorPrimaries = "sgamut2"         // Sony
	ColorPrimariesSgamut3         trimmer.ColorPrimaries = "sgamut3"         // Sony
	ColorPrimariesSgamut3cine     trimmer.ColorPrimaries = "sgamut3cine"     // Sony
	ColorPrimariesREDcolor        trimmer.ColorPrimaries = "redcolor"        // RED
	ColorPrimariesREDcolor2       trimmer.ColorPrimaries = "redcolor2"       // RED
	ColorPrimariesREDcolor3       trimmer.ColorPrimaries = "redcolor3"       // RED
	ColorPrimariesREDcolor4       trimmer.ColorPrimaries = "redcolor4"       // RED
	ColorPrimariesDragonColor     trimmer.ColorPrimaries = "dragoncolor"     // RED
	ColorPrimariesDragonColor2    trimmer.ColorPrimaries = "dragoncolor2"    // RED
	ColorPrimariesREDWideGamutRGB trimmer.ColorPrimaries = "redwidegamutrgb" // RED
)

const (
	ColorTransferUndefined    trimmer.ColorTransfer = ""
	ColorTransferReserved0    trimmer.ColorTransfer = "reserved0"    // 0: reserved
	ColorTransferBT709        trimmer.ColorTransfer = "bt709"        // 1: BT.709, BT.1316-0; (== 6, 14, 15; default)
	ColorTransferUnspecified  trimmer.ColorTransfer = "unspecified"  // 2: Unspecified
	ColorTransferReserved     trimmer.ColorTransfer = "reserved"     // 3: Reserved
	ColorTransferBT470M       trimmer.ColorTransfer = "bt470m"       // 4: BT.470M, BT.1700 PAL/SECAM (display gamma 2.2)
	ColorTransferBT470BG      trimmer.ColorTransfer = "bt470bg"      // 5: BT.470BG (display gamma 2.8)
	ColorTransferSMPTE170M    trimmer.ColorTransfer = "smpte170m"    // 6: BT.601 525, BT.1700 NTSC, SMPTE 170M:2004 (== _1_, 14, 15)
	ColorTransferSMPTE240M    trimmer.ColorTransfer = "smpte240m"    // 7: SMPTE 240M:1999
	ColorTransferLINEAR       trimmer.ColorTransfer = "linear"       // 8: Linear
	ColorTransferLOG100       trimmer.ColorTransfer = "log100"       // 9: Log
	ColorTransferLOG316       trimmer.ColorTransfer = "log316"       // 10: Log SQRT
	ColorTransferIEC61966_24  trimmer.ColorTransfer = "iec61966-24"  // 11: IEC 61966-2-4 xvYCC
	ColorTransferBT1361       trimmer.ColorTransfer = "bt1361"       // 12: BT.1361-0 ITU-R BT1361 Extended Colour Gamut
	ColorTransferIEC61966_21  trimmer.ColorTransfer = "iec61966-21"  // 13: IEC 61966-2-1 sRGB
	ColorTransferBT2020_10    trimmer.ColorTransfer = "bt2020-10"    // 14: BT.2020 10bit (== _1_, 6, 15)
	ColorTransferBT2020_12    trimmer.ColorTransfer = "bt2020-12"    // 15: BT.2020 12bit (== _1_, 6, 14)
	ColorTransferSMPTE2084    trimmer.ColorTransfer = "smpte2084"    // 16: SMPTE ST2084:2014 (10/12/14/16 bit), BT.2100 [HDR PQ]
	ColorTransferSMPTE428     trimmer.ColorTransfer = "smpte428"     // 17: SMPTE ST 428-1, DCI
	ColorTransferARIB_STD_B67 trimmer.ColorTransfer = "arib-std-b67" // 18: ARIB STD-B67, HDR HLG

	// ACES transfer characteristics (Note: ACES is an RGB color space)
	ColorTransferACES2065_1 trimmer.ColorTransfer = "aces_2065_1" // ST2065-1 and TB-2014-004
	ColorTransferACEScc     trimmer.ColorTransfer = "aces_cc"     // S-2014-003
	ColorTransferACEScct    trimmer.ColorTransfer = "aces_cct"    // S-2016-001
	ColorTransferACESproxy  trimmer.ColorTransfer = "aces_proxy"  // S-2013-001
	ColorTransferACEScg     trimmer.ColorTransfer = "aces_cg"     // S-2014-004

	// Film stock
	ColorTransferFilm trimmer.ColorTransfer = "film" // generic film stock, details in other metadata
	ColorTransferADX  trimmer.ColorTransfer = "adx"  // SMPTE ST 2065-3 (ADX, Academy Density Exchange Encoding)

	// vendor-specific color transfers
	ColorTransferAdobeRGB   trimmer.ColorTransfer = "adobergb"   // Adobe RGB
	ColorTransferLogC       trimmer.ColorTransfer = "log-c"      // Arri LogC
	ColorTransferSlog       trimmer.ColorTransfer = "slog"       // Sony
	ColorTransferSlog2      trimmer.ColorTransfer = "slog2"      // Sony
	ColorTransferSlog3      trimmer.ColorTransfer = "slog3"      // Sony
	ColorTransferCineonLog  trimmer.ColorTransfer = "cineon-log" // DPX
	ColorTransferCLog       trimmer.ColorTransfer = "c-log"      // Canon
	ColorTransferCLog2      trimmer.ColorTransfer = "c-log2"     // Canon
	ColorTransferCLog3      trimmer.ColorTransfer = "c-log3"     // Canon
	ColorTransferVLog       trimmer.ColorTransfer = "v-log"      // Panasonic (Varicam)
	ColorTransferREDlog     trimmer.ColorTransfer = "redlog"     // RED
	ColorTransferREDlogFilm trimmer.ColorTransfer = "redlogfilm" // RED
	ColorTransferREDLog3G10 trimmer.ColorTransfer = "redlog3g10" // RED
	ColorTransferREDgamma   trimmer.ColorTransfer = "redgamma"   // RED
	ColorTransferREDgamma2  trimmer.ColorTransfer = "redgamma2"  // RED
	ColorTransferREDgamma3  trimmer.ColorTransfer = "redgamma3"  // RED
	ColorTransferREDgamma4  trimmer.ColorTransfer = "redgamma4"  // RED
)

const (
	ColorRangeUndefined   trimmer.ColorRange = ""
	ColorRangeUnspecified trimmer.ColorRange = "unspecified"
	ColorRangeLegal       trimmer.ColorRange = "legal"
	ColorRangeFull        trimmer.ColorRange = "full"
)

const (
	ColorLocationUndefined   trimmer.ColorLocation = ""
	ColorLocationUnspecified trimmer.ColorLocation = "unspecified"
	ColorLocationLeft        trimmer.ColorLocation = "left"
	ColorLocationCenter      trimmer.ColorLocation = "center"
	ColorLocationTop         trimmer.ColorLocation = "top"
	ColorLocationTopLeft     trimmer.ColorLocation = "topleft"
	ColorLocationBottom      trimmer.ColorLocation = "bottom"
	ColorLocationBottomLeft  trimmer.ColorLocation = "bottomleft"
)

func IsValidColorModel(s trimmer.ColorModel) bool {
	switch s {
	case ColorModelUndefined, ColorModelUnspecified, ColorModelReserved:
		return false
	default:
		return true
	}
}

func IsValidColorPrimaries(p trimmer.ColorPrimaries) bool {
	switch p {
	case ColorPrimariesUndefined, ColorPrimariesUnspecified, ColorPrimariesReserved, ColorPrimariesReserved0:
		return false
	default:
		return true
	}
}

func IsValidColorTransfer(s trimmer.ColorTransfer) bool {
	switch s {
	case ColorTransferUndefined, ColorTransferUnspecified, ColorTransferReserved, ColorTransferReserved0:
		return false
	default:
		return true
	}
}

func IsValidColorRange(s trimmer.ColorRange) bool {
	switch s {
	case ColorRangeUndefined, ColorRangeUnspecified:
		return false
	default:
		return true
	}
}

func IsValidColorLocation(s trimmer.ColorLocation) bool {
	switch s {
	case ColorLocationUndefined, ColorLocationUnspecified:
		return false
	default:
		return true
	}
}

func ParseColorModel(s string) trimmer.ColorModel {
	switch strings.ToLower(s) {
	case "rgb":
		return ColorModelRGB
	case "rec-709", "rec709", "bt709":
		return ColorModelBT709
	case "unspecified", "unknown":
		return ColorModelUnspecified
	case "reserved":
		return ColorModelReserved
	case "fcc":
		return ColorModelFCC
	case "bt470bg":
		return ColorModelBT470BG
	case "smpte170m":
		return ColorModelSMPTE170M
	case "smpte240m":
		return ColorModelSMPTE240M
	case "ycgco":
		return ColorModelYCGCO
	case "bt2020nc", "bt2020ncl":
		return ColorModelBT2020NCL
	case "smpte2085", "bt2020cl":
		return ColorModelBT2020CL
	case "ydzyx":
		return ColorModelYDZYX
	case "derived_ncl":
		return ColorModelDERIVED_NCL
	case "derived_cl":
		return ColorModelDERIVED_CL
	case "ictcp":
		return ColorModelICTCP
	case "ycbcr":
		return ColorModelYCBCR
	case "xvycc":
		return ColorModelXVYCC
	case "srgb":
		return ColorModelSRGB
	case "adobergb":
		return ColorModelAdobeRGB
	case "hsv":
		return ColorModelHSV
	case "hls":
		return ColorModelHLS
	case "cie_xyz", "xyz":
		return ColorModelXYZ
	default:
		return ColorModelUndefined
	}
}

func ParseColorPrimaries(s string) trimmer.ColorPrimaries {
	switch strings.ToLower(s) {
	case "reserved0":
		return ColorPrimariesReserved0
	case "rec-709", "rec709", "bt709":
		return ColorPrimariesBT709
	case "unspecified", "unknown":
		return ColorPrimariesUnspecified
	case "reserved":
		return ColorPrimariesReserved
	case "bt470m":
		return ColorPrimariesBT470M
	case "bt470bg":
		return ColorPrimariesBT470BG
	case "smpte170m":
		return ColorPrimariesSMPTE170M
	case "smpte240m":
		return ColorPrimariesSMPTE240M
	case "film":
		return ColorPrimariesFilm
	case "bt2020":
		return ColorPrimariesBT2020
	case "smpte428":
		return ColorPrimariesSMPTE428
	case "p3dci", "smpte431":
		return ColorPrimariesP3DCI
	case "p3d65", "smpte432":
		return ColorPrimariesP3D65
	case "p3d60":
		return ColorPrimariesP3D60
	case "jedec-p22":
		return ColorPrimariesJedec22
	case "xyz":
		return ColorPrimariesXYZ
	case "aces-ap0":
		return ColorPrimariesAcesAP0
	case "aces-ap1":
		return ColorPrimariesAcesAP1
	case "native":
		return ColorPrimariesNative
	case "adobergb":
		return ColorPrimariesAdobeRGB
	case "logcwgam":
		return ColorPrimariesAlexaWideGamut
	case "vgamut":
		return ColorPrimariesVgamut
	case "sgamut":
		return ColorPrimariesSgamut
	case "sgamut2":
		return ColorPrimariesSgamut2
	case "sgamut3":
		return ColorPrimariesSgamut3
	case "sgamut3cine":
		return ColorPrimariesSgamut3cine
	case "redcolor":
		return ColorPrimariesREDcolor
	case "redcolor2":
		return ColorPrimariesREDcolor2
	case "redcolor3":
		return ColorPrimariesREDcolor3
	case "redcolor4":
		return ColorPrimariesREDcolor4
	case "dragoncolor":
		return ColorPrimariesDragonColor
	case "dragoncolor2":
		return ColorPrimariesDragonColor2
	case "redwidegamutrgb":
		return ColorPrimariesREDWideGamutRGB
	default:
		return ColorPrimariesUndefined
	}
}

func ParseColorTransfer(s string) trimmer.ColorTransfer {
	switch strings.ToLower(s) {
	case "reserved0":
		return ColorTransferReserved0
	case "rec-709", "rec709", "bt709":
		return ColorTransferBT709
	case "unspecified", "unknown":
		return ColorTransferUnspecified
	case "reserved":
		return ColorTransferReserved
	case "bt470m":
		return ColorTransferBT470M
	case "bt470bg":
		return ColorTransferBT470BG
	case "smpte170m":
		return ColorTransferSMPTE170M
	case "smpte240m":
		return ColorTransferSMPTE240M
	case "linear":
		return ColorTransferLINEAR
	case "log100":
		return ColorTransferLOG100
	case "log316":
		return ColorTransferLOG316
	case "iec61966-21", "iec61966_21", "iec61966_2_1", "iec61966-2-1":
		return ColorTransferIEC61966_21
	case "bt1361", "bt1361e":
		return ColorTransferBT1361
	case "iec61966-24", "iec61966_24", "iec61966_2_4", "iec61966-2-4":
		return ColorTransferIEC61966_24
	case "bt2020-10", "bt2020_10", "bt2020_10bit", "bt2020-10bit":
		return ColorTransferBT2020_10
	case "bt2020-12", "bt2020_12", "bt2020_12bit", "bt2020-12bit":
		return ColorTransferBT2020_12
	case "smpte2084":
		return ColorTransferSMPTE2084
	case "smpte428":
		return ColorTransferSMPTE428
	case "arib-std-b67", "arib_std_b67":
		return ColorTransferARIB_STD_B67
	case "aces_2065_1":
		return ColorTransferACES2065_1
	case "aces_cc":
		return ColorTransferACEScc
	case "aces_cct":
		return ColorTransferACEScct
	case "aces_proxy":
		return ColorTransferACESproxy
	case "aces_cg":
		return ColorTransferACEScg
	case "film":
		return ColorTransferFilm
	case "adx":
		return ColorTransferADX
	case "adobergb":
		return ColorTransferAdobeRGB
	case "log-c":
		return ColorTransferLogC
	case "slog":
		return ColorTransferSlog
	case "slog2":
		return ColorTransferSlog2
	case "slog3":
		return ColorTransferSlog3
	case "cineon-log":
		return ColorTransferCineonLog
	case "c-log":
		return ColorTransferCLog
	case "c-log2":
		return ColorTransferCLog2
	case "c-log3":
		return ColorTransferCLog3
	case "v-log":
		return ColorTransferVLog
	case "redlog":
		return ColorTransferREDlog
	case "redlogfilm":
		return ColorTransferREDlogFilm
	case "redlog3g10":
		return ColorTransferREDLog3G10
	case "redgamma":
		return ColorTransferREDgamma
	case "redgamma2":
		return ColorTransferREDgamma2
	case "redgamma3":
		return ColorTransferREDgamma3
	case "redgamma4":
		return ColorTransferREDgamma4
	default:
		return ColorTransferUndefined
	}
}

func ParseColorRange(s string) trimmer.ColorRange {
	switch strings.ToLower(s) {
	case "unspecified", "unknown":
		return ColorRangeUnspecified
	case "tv", "mpeg", "legal", "limited":
		return ColorRangeLegal
	case "pc", "jpeg", "full":
		return ColorRangeFull
	default:
		return ColorRangeUndefined
	}
}

func ParseColorLocation(s string) trimmer.ColorLocation {
	switch strings.ToLower(s) {
	case "unspecified", "unknown":
		return ColorLocationUnspecified
	case "left":
		return ColorLocationLeft
	case "center":
		return ColorLocationCenter
	case "top":
		return ColorLocationTop
	case "topleft":
		return ColorLocationTopLeft
	case "bottom":
		return ColorLocationBottom
	case "bottomleft":
		return ColorLocationBottomLeft
	default:
		return ColorLocationUndefined
	}
}

const (
	PixelFormatUndefined     trimmer.PixelFormat = ""
	PixelFormatUnspecified   trimmer.PixelFormat = "unspecified"
	PixelFormatNone          trimmer.PixelFormat = "none"
	PixelFormatYUV420P       trimmer.PixelFormat = "yuv420p"
	PixelFormatYUYV422       trimmer.PixelFormat = "yuyv422"
	PixelFormatRGB24         trimmer.PixelFormat = "rgb24"
	PixelFormatBGR24         trimmer.PixelFormat = "bgr24"
	PixelFormatYUV422P       trimmer.PixelFormat = "yuv422p"
	PixelFormatYUV444P       trimmer.PixelFormat = "yuv444p"
	PixelFormatHalf          trimmer.PixelFormat = "half"
	PixelFormatFloat         trimmer.PixelFormat = "float"
	PixelFormatUint          trimmer.PixelFormat = "uint"
	PixelFormatYUV410P       trimmer.PixelFormat = "yuv410p"
	PixelFormatYUV411P       trimmer.PixelFormat = "yuv411p"
	PixelFormatYUVJ411P      trimmer.PixelFormat = "yuvj411p"
	PixelFormatPAL8          trimmer.PixelFormat = "pal8"
	PixelFormatYUVJ420P      trimmer.PixelFormat = "yuvj420p"
	PixelFormatYUVJ422P      trimmer.PixelFormat = "yuvj422p"
	PixelFormatYUVJ444P      trimmer.PixelFormat = "yuvj444p"
	PixelFormatUYVY422       trimmer.PixelFormat = "uyvy422"
	PixelFormatUYYVYY411     trimmer.PixelFormat = "uyyvyy411"
	PixelFormatBGR8          trimmer.PixelFormat = "bgr8"
	PixelFormatRGB8          trimmer.PixelFormat = "rgb8"
	PixelFormatNV12          trimmer.PixelFormat = "nv12"
	PixelFormatNV21          trimmer.PixelFormat = "nv21"
	PixelFormatARGB          trimmer.PixelFormat = "argb"
	PixelFormatRGBA          trimmer.PixelFormat = "rgba"
	PixelFormatABGR          trimmer.PixelFormat = "abgr"
	PixelFormatBGRA          trimmer.PixelFormat = "bgra"
	PixelFormat0RGB          trimmer.PixelFormat = "0rgb"
	PixelFormatRGB0          trimmer.PixelFormat = "rgb0"
	PixelFormat0BGR          trimmer.PixelFormat = "0bgr"
	PixelFormatBGR0          trimmer.PixelFormat = "bgr0"
	PixelFormatGRAY16BE      trimmer.PixelFormat = "gray16be"
	PixelFormatGRAY16LE      trimmer.PixelFormat = "gray16le"
	PixelFormatYUV440P       trimmer.PixelFormat = "yuv440p"
	PixelFormatYUVJ440P      trimmer.PixelFormat = "yuvj440p"
	PixelFormatYUVA420P      trimmer.PixelFormat = "yuva420p"
	PixelFormatYUVA422P      trimmer.PixelFormat = "yuva422p"
	PixelFormatYUVA444P      trimmer.PixelFormat = "yuva444p"
	PixelFormatYUVA420P9BE   trimmer.PixelFormat = "yuva420p9be"
	PixelFormatYUVA420P9LE   trimmer.PixelFormat = "yuva420p9le"
	PixelFormatYUVA422P9BE   trimmer.PixelFormat = "yuva422p9be"
	PixelFormatYUVA422P9LE   trimmer.PixelFormat = "yuva422p9le"
	PixelFormatYUVA444P9BE   trimmer.PixelFormat = "yuva444p9be"
	PixelFormatYUVA444P9LE   trimmer.PixelFormat = "yuva444p9le"
	PixelFormatYUVA420P10BE  trimmer.PixelFormat = "yuva420p10be"
	PixelFormatYUVA420P10LE  trimmer.PixelFormat = "yuva420p10le"
	PixelFormatYUVA422P10BE  trimmer.PixelFormat = "yuva422p10be"
	PixelFormatYUVA422P10LE  trimmer.PixelFormat = "yuva422p10le"
	PixelFormatYUVA444P10BE  trimmer.PixelFormat = "yuva444p10be"
	PixelFormatYUVA444P10LE  trimmer.PixelFormat = "yuva444p10le"
	PixelFormatYUVA420P16BE  trimmer.PixelFormat = "yuva420p16be"
	PixelFormatYUVA420P16LE  trimmer.PixelFormat = "yuva420p16le"
	PixelFormatYUVA422P16BE  trimmer.PixelFormat = "yuva422p16be"
	PixelFormatYUVA422P16LE  trimmer.PixelFormat = "yuva422p16le"
	PixelFormatYUVA444P16BE  trimmer.PixelFormat = "yuva444p16be"
	PixelFormatYUVA444P16LE  trimmer.PixelFormat = "yuva444p16le"
	PixelFormatRGB48BE       trimmer.PixelFormat = "rgb48be"
	PixelFormatRGB48LE       trimmer.PixelFormat = "rgb48le"
	PixelFormatRGBA64BE      trimmer.PixelFormat = "rgba64be"
	PixelFormatRGBA64LE      trimmer.PixelFormat = "rgba64le"
	PixelFormatRGB565BE      trimmer.PixelFormat = "rgb565be"
	PixelFormatRGB565LE      trimmer.PixelFormat = "rgb565le"
	PixelFormatRGB555BE      trimmer.PixelFormat = "rgb555be"
	PixelFormatRGB555LE      trimmer.PixelFormat = "rgb555le"
	PixelFormatRGB444BE      trimmer.PixelFormat = "rgb444be"
	PixelFormatRGB444LE      trimmer.PixelFormat = "rgb444le"
	PixelFormatBGR48BE       trimmer.PixelFormat = "bgr48be"
	PixelFormatBGR48LE       trimmer.PixelFormat = "bgr48le"
	PixelFormatBGRA64BE      trimmer.PixelFormat = "bgra64be"
	PixelFormatBGRA64LE      trimmer.PixelFormat = "bgra64le"
	PixelFormatBGR565BE      trimmer.PixelFormat = "bgr565be"
	PixelFormatBGR565LE      trimmer.PixelFormat = "bgr565le"
	PixelFormatBGR555BE      trimmer.PixelFormat = "bgr555be"
	PixelFormatBGR555LE      trimmer.PixelFormat = "bgr555le"
	PixelFormatBGR444BE      trimmer.PixelFormat = "bgr444be"
	PixelFormatBGR444LE      trimmer.PixelFormat = "bgr444le"
	PixelFormatYUV420P9LE    trimmer.PixelFormat = "yuv420p9le"
	PixelFormatYUV420P9BE    trimmer.PixelFormat = "yuv420p9be"
	PixelFormatYUV420P10LE   trimmer.PixelFormat = "yuv420p10le"
	PixelFormatYUV420P10BE   trimmer.PixelFormat = "yuv420p10be"
	PixelFormatYUV420P12LE   trimmer.PixelFormat = "yuv420p12le"
	PixelFormatYUV420P12BE   trimmer.PixelFormat = "yuv420p12be"
	PixelFormatYUV420P14LE   trimmer.PixelFormat = "yuv420p14le"
	PixelFormatYUV420P14BE   trimmer.PixelFormat = "yuv420p14be"
	PixelFormatYUV420P16LE   trimmer.PixelFormat = "yuv420p16le"
	PixelFormatYUV420P16BE   trimmer.PixelFormat = "yuv420p16be"
	PixelFormatYUV422P9LE    trimmer.PixelFormat = "yuv422p9le"
	PixelFormatYUV422P9BE    trimmer.PixelFormat = "yuv422p9be"
	PixelFormatYUV422P10LE   trimmer.PixelFormat = "yuv422p10le"
	PixelFormatYUV422P10BE   trimmer.PixelFormat = "yuv422p10be"
	PixelFormatYUV422P12LE   trimmer.PixelFormat = "yuv422p12le"
	PixelFormatYUV422P12BE   trimmer.PixelFormat = "yuv422p12be"
	PixelFormatYUV422P14LE   trimmer.PixelFormat = "yuv422p14le"
	PixelFormatYUV422P14BE   trimmer.PixelFormat = "yuv422p14be"
	PixelFormatYUV422P16LE   trimmer.PixelFormat = "yuv422p16le"
	PixelFormatYUV422P16BE   trimmer.PixelFormat = "yuv422p16be"
	PixelFormatYUV444P16LE   trimmer.PixelFormat = "yuv444p16le"
	PixelFormatYUV444P16BE   trimmer.PixelFormat = "yuv444p16be"
	PixelFormatYUV444P10LE   trimmer.PixelFormat = "yuv444p10le"
	PixelFormatYUV444P10BE   trimmer.PixelFormat = "yuv444p10be"
	PixelFormatYUV444P9LE    trimmer.PixelFormat = "yuv444p9le"
	PixelFormatYUV444P9BE    trimmer.PixelFormat = "yuv444p9be"
	PixelFormatYUV444P12LE   trimmer.PixelFormat = "yuv444p12le"
	PixelFormatYUV444P12BE   trimmer.PixelFormat = "yuv444p12be"
	PixelFormatYUV444P14LE   trimmer.PixelFormat = "yuv444p14le"
	PixelFormatYUV444P14BE   trimmer.PixelFormat = "yuv444p14be"
	PixelFormatGRAY8         trimmer.PixelFormat = "gray8"
	PixelFormatGRAY8A        trimmer.PixelFormat = "gray8a"
	PixelFormatGBRP          trimmer.PixelFormat = "gbrp"
	PixelFormatGBRP9LE       trimmer.PixelFormat = "gbrp9le"
	PixelFormatGBRP9BE       trimmer.PixelFormat = "gbrp9be"
	PixelFormatGBRP10LE      trimmer.PixelFormat = "gbrp10le"
	PixelFormatGBRP10BE      trimmer.PixelFormat = "gbrp10be"
	PixelFormatGBRP12LE      trimmer.PixelFormat = "gbrp12le"
	PixelFormatGBRP12BE      trimmer.PixelFormat = "gbrp12be"
	PixelFormatGBRP14LE      trimmer.PixelFormat = "gbrp14le"
	PixelFormatGBRP14BE      trimmer.PixelFormat = "gbrp14be"
	PixelFormatGBRP16LE      trimmer.PixelFormat = "gbrp16le"
	PixelFormatGBRP16BE      trimmer.PixelFormat = "gbrp16be"
	PixelFormatGBRAP         trimmer.PixelFormat = "gbrap"
	PixelFormatGBRAP16LE     trimmer.PixelFormat = "gbrap16le"
	PixelFormatGBRAP16BE     trimmer.PixelFormat = "gbrap16be"
	PixelFormatXYZ12LE       trimmer.PixelFormat = "xyz12le"
	PixelFormatXYZ12BE       trimmer.PixelFormat = "xyz12be"
	PixelFormatBayerBGGR8    trimmer.PixelFormat = "bayer_bggr8"
	PixelFormatBayerRGGB8    trimmer.PixelFormat = "bayer_rggb8"
	PixelFormatBayerGBRG8    trimmer.PixelFormat = "bayer_gbrg8"
	PixelFormatBayerGRBG8    trimmer.PixelFormat = "bayer_grbg8"
	PixelFormatBayerBGGR16LE trimmer.PixelFormat = "bayer_bggr16le"
	PixelFormatBayerBGGR16BE trimmer.PixelFormat = "bayer_bggr16be"
	PixelFormatBayerRGGB16LE trimmer.PixelFormat = "bayer_rggb16le"
	PixelFormatBayerRGGB16BE trimmer.PixelFormat = "bayer_rggb16be"
	PixelFormatBayerGBRG16LE trimmer.PixelFormat = "bayer_gbrg16le"
	PixelFormatBayerGBRG16BE trimmer.PixelFormat = "bayer_gbrg16be"
	PixelFormatBayerGRBG16LE trimmer.PixelFormat = "bayer_grbg16le"
	PixelFormatBayerGRBG16BE trimmer.PixelFormat = "bayer_grbg16be"
)

func ParsePixelFormat(s string) trimmer.PixelFormat {
	switch strings.ToLower(s) {
	case "unspecified":
		return PixelFormatUnspecified
	case "none":
		return PixelFormatNone
	case "rgb", "rgb24":
		return PixelFormatRGB24
	case "bgr", "bgr24":
		return PixelFormatBGR24
	case "yuv420p":
		return PixelFormatYUV420P
	case "yuv422p":
		return PixelFormatYUV422P
	case "yuv444p":
		return PixelFormatYUV444P
	case "half":
		return PixelFormatHalf
	case "float":
		return PixelFormatFloat
	case "uint":
		return PixelFormatUint
	case "yuv410p":
		return PixelFormatYUV410P
	case "yuv411p":
		return PixelFormatYUV411P
	case "yuvj411p":
		return PixelFormatYUVJ411P
	case "yuvj420p":
		return PixelFormatYUVJ420P
	case "yuvj422p":
		return PixelFormatYUVJ422P
	case "yuvj444p":
		return PixelFormatYUVJ444P
	case "yuyv422":
		return PixelFormatYUYV422
	case "uyvy422":
		return PixelFormatUYVY422
	case "uyyvyy411":
		return PixelFormatUYYVYY411
	case "pal8":
		return PixelFormatPAL8
	case "bgr8":
		return PixelFormatBGR8
	case "rgb8":
		return PixelFormatRGB8
	case "nv12":
		return PixelFormatNV12
	case "nv21":
		return PixelFormatNV21
	case "argb":
		return PixelFormatARGB
	case "rgba":
		return PixelFormatRGBA
	case "abgr":
		return PixelFormatABGR
	case "bgra":
		return PixelFormatBGRA
	case "0rgb":
		return PixelFormat0RGB
	case "rgb0":
		return PixelFormatRGB0
	case "0bgr":
		return PixelFormat0BGR
	case "bgr0":
		return PixelFormatBGR0
	case "gray16be":
		return PixelFormatGRAY16BE
	case "gray16le":
		return PixelFormatGRAY16LE
	case "yuv440p":
		return PixelFormatYUV440P
	case "yuvj440p":
		return PixelFormatYUVJ440P
	case "yuva420p":
		return PixelFormatYUVA420P
	case "yuva422p":
		return PixelFormatYUVA422P
	case "yuva444p":
		return PixelFormatYUVA444P
	case "yuva420p9be":
		return PixelFormatYUVA420P9BE
	case "yuva420p9le":
		return PixelFormatYUVA420P9LE
	case "yuva422p9be":
		return PixelFormatYUVA422P9BE
	case "yuva422p9le":
		return PixelFormatYUVA422P9LE
	case "yuva444p9be":
		return PixelFormatYUVA444P9BE
	case "yuva444p9le":
		return PixelFormatYUVA444P9LE
	case "yuva420p10be":
		return PixelFormatYUVA420P10BE
	case "yuva420p10le":
		return PixelFormatYUVA420P10LE
	case "yuva422p10be":
		return PixelFormatYUVA422P10BE
	case "yuva422p10le":
		return PixelFormatYUVA422P10LE
	case "yuva444p10be":
		return PixelFormatYUVA444P10BE
	case "yuva444p10le":
		return PixelFormatYUVA444P10LE
	case "yuva420p16be":
		return PixelFormatYUVA420P16BE
	case "yuva420p16le":
		return PixelFormatYUVA420P16LE
	case "yuva422p16be":
		return PixelFormatYUVA422P16BE
	case "yuva422p16le":
		return PixelFormatYUVA422P16LE
	case "yuva444p16be":
		return PixelFormatYUVA444P16BE
	case "yuva444p16le":
		return PixelFormatYUVA444P16LE
	case "rgb48be":
		return PixelFormatRGB48BE
	case "rgb48le":
		return PixelFormatRGB48LE
	case "rgba64be":
		return PixelFormatRGBA64BE
	case "rgba64le":
		return PixelFormatRGBA64LE
	case "rgb565be":
		return PixelFormatRGB565BE
	case "rgb565le":
		return PixelFormatRGB565LE
	case "rgb555be":
		return PixelFormatRGB555BE
	case "rgb555le":
		return PixelFormatRGB555LE
	case "rgb444be":
		return PixelFormatRGB444BE
	case "rgb444le":
		return PixelFormatRGB444LE
	case "bgr48be":
		return PixelFormatBGR48BE
	case "bgr48le":
		return PixelFormatBGR48LE
	case "bgra64be":
		return PixelFormatBGRA64BE
	case "bgra64le":
		return PixelFormatBGRA64LE
	case "bgr565be":
		return PixelFormatBGR565BE
	case "bgr565le":
		return PixelFormatBGR565LE
	case "bgr555be":
		return PixelFormatBGR555BE
	case "bgr555le":
		return PixelFormatBGR555LE
	case "bgr444be":
		return PixelFormatBGR444BE
	case "bgr444le":
		return PixelFormatBGR444LE
	case "yuv420p9le":
		return PixelFormatYUV420P9LE
	case "yuv420p9be":
		return PixelFormatYUV420P9BE
	case "yuv420p10le":
		return PixelFormatYUV420P10LE
	case "yuv420p10be":
		return PixelFormatYUV420P10BE
	case "yuv420p12le":
		return PixelFormatYUV420P12LE
	case "yuv420p12be":
		return PixelFormatYUV420P12BE
	case "yuv420p14le":
		return PixelFormatYUV420P14LE
	case "yuv420p14be":
		return PixelFormatYUV420P14BE
	case "yuv420p16le":
		return PixelFormatYUV420P16LE
	case "yuv420p16be":
		return PixelFormatYUV420P16BE
	case "yuv422p9le":
		return PixelFormatYUV422P9LE
	case "yuv422p9be":
		return PixelFormatYUV422P9BE
	case "yuv422p10le":
		return PixelFormatYUV422P10LE
	case "yuv422p10be":
		return PixelFormatYUV422P10BE
	case "yuv422p12le":
		return PixelFormatYUV422P12LE
	case "yuv422p12be":
		return PixelFormatYUV422P12BE
	case "yuv422p14le":
		return PixelFormatYUV422P14LE
	case "yuv422p14be":
		return PixelFormatYUV422P14BE
	case "yuv422p16le":
		return PixelFormatYUV422P16LE
	case "yuv422p16be":
		return PixelFormatYUV422P16BE
	case "yuv444p16le":
		return PixelFormatYUV444P16LE
	case "yuv444p16be":
		return PixelFormatYUV444P16BE
	case "yuv444p10le":
		return PixelFormatYUV444P10LE
	case "yuv444p10be":
		return PixelFormatYUV444P10BE
	case "yuv444p9le":
		return PixelFormatYUV444P9LE
	case "yuv444p9be":
		return PixelFormatYUV444P9BE
	case "yuv444p12le":
		return PixelFormatYUV444P12LE
	case "yuv444p12be":
		return PixelFormatYUV444P12BE
	case "yuv444p14le":
		return PixelFormatYUV444P14LE
	case "yuv444p14be":
		return PixelFormatYUV444P14BE
	case "gray8":
		return PixelFormatGRAY8
	case "gray8a":
		return PixelFormatGRAY8A
	case "gbrp":
		return PixelFormatGBRP
	case "gbrp9le":
		return PixelFormatGBRP9LE
	case "gbrp9be":
		return PixelFormatGBRP9BE
	case "gbrp10le":
		return PixelFormatGBRP10LE
	case "gbrp10be":
		return PixelFormatGBRP10BE
	case "gbrp12le":
		return PixelFormatGBRP12LE
	case "gbrp12be":
		return PixelFormatGBRP12BE
	case "gbrp14le":
		return PixelFormatGBRP14LE
	case "gbrp14be":
		return PixelFormatGBRP14BE
	case "gbrp16le":
		return PixelFormatGBRP16LE
	case "gbrp16be":
		return PixelFormatGBRP16BE
	case "gbrap":
		return PixelFormatGBRAP
	case "gbrap16le":
		return PixelFormatGBRAP16LE
	case "gbrap16be":
		return PixelFormatGBRAP16BE
	case "xyz12le":
		return PixelFormatXYZ12LE
	case "xyz12be":
		return PixelFormatXYZ12BE
	case "bayer_bggr8":
		return PixelFormatBayerBGGR8
	case "bayer_rggb8":
		return PixelFormatBayerRGGB8
	case "bayer_gbrg8":
		return PixelFormatBayerGBRG8
	case "bayer_grbg8":
		return PixelFormatBayerGRBG8
	case "bayer_bggr16le":
		return PixelFormatBayerBGGR16LE
	case "bayer_bggr16be":
		return PixelFormatBayerBGGR16BE
	case "bayer_rggb16le":
		return PixelFormatBayerRGGB16LE
	case "bayer_rggb16be":
		return PixelFormatBayerRGGB16BE
	case "bayer_gbrg16le":
		return PixelFormatBayerGBRG16LE
	case "bayer_gbrg16be":
		return PixelFormatBayerGBRG16BE
	case "bayer_grbg16le":
		return PixelFormatBayerGRBG16LE
	case "bayer_grbg16be":
		return PixelFormatBayerGRBG16BE
	default:
		return PixelFormatUndefined
	}
}

func PixelDepth(f trimmer.PixelFormat) int {
	switch f {
	case PixelFormatYUVA420P9BE,
		PixelFormatYUVA420P9LE,
		PixelFormatYUVA422P9BE,
		PixelFormatYUVA422P9LE,
		PixelFormatYUVA444P9BE,
		PixelFormatYUVA444P9LE,
		PixelFormatYUV420P9LE,
		PixelFormatYUV420P9BE,
		PixelFormatYUV422P9LE,
		PixelFormatYUV422P9BE,
		PixelFormatYUV444P9LE,
		PixelFormatYUV444P9BE,
		PixelFormatGBRP9LE,
		PixelFormatGBRP9BE:
		return 9

	case PixelFormatYUVA420P10BE,
		PixelFormatYUVA420P10LE,
		PixelFormatYUVA422P10BE,
		PixelFormatYUVA422P10LE,
		PixelFormatYUVA444P10BE,
		PixelFormatYUVA444P10LE,
		PixelFormatYUV420P10LE,
		PixelFormatYUV420P10BE,
		PixelFormatYUV422P10LE,
		PixelFormatYUV422P10BE,
		PixelFormatYUV444P10LE,
		PixelFormatYUV444P10BE,
		PixelFormatGBRP10LE,
		PixelFormatGBRP10BE:
		return 10

	case PixelFormatXYZ12LE,
		PixelFormatXYZ12BE,
		PixelFormatYUV420P12LE,
		PixelFormatYUV420P12BE,
		PixelFormatYUV422P12LE,
		PixelFormatYUV422P12BE,
		PixelFormatYUV444P12LE,
		PixelFormatYUV444P12BE,
		PixelFormatGBRP12LE,
		PixelFormatGBRP12BE:
		return 12

	case
		PixelFormatYUV420P14LE,
		PixelFormatYUV420P14BE,
		PixelFormatYUV422P14LE,
		PixelFormatYUV422P14BE,
		PixelFormatYUV444P14LE,
		PixelFormatYUV444P14BE,
		PixelFormatGBRP14LE,
		PixelFormatGBRP14BE:
		return 14

	case
		PixelFormatHalf,
		PixelFormatYUV420P16LE,
		PixelFormatYUV420P16BE,
		PixelFormatYUV422P16LE,
		PixelFormatYUV422P16BE,
		PixelFormatYUV444P16LE,
		PixelFormatYUV444P16BE,
		PixelFormatGRAY16BE,
		PixelFormatGRAY16LE,
		PixelFormatYUVA420P16BE,
		PixelFormatYUVA420P16LE,
		PixelFormatYUVA422P16BE,
		PixelFormatYUVA422P16LE,
		PixelFormatYUVA444P16BE,
		PixelFormatYUVA444P16LE,
		PixelFormatRGB48BE,
		PixelFormatRGB48LE,
		PixelFormatRGBA64BE,
		PixelFormatRGBA64LE,
		PixelFormatRGB565BE,
		PixelFormatRGB565LE,
		PixelFormatRGB555BE,
		PixelFormatRGB555LE,
		PixelFormatRGB444BE,
		PixelFormatRGB444LE,
		PixelFormatBGR48BE,
		PixelFormatBGR48LE,
		PixelFormatBGRA64BE,
		PixelFormatBGRA64LE,
		PixelFormatBGR565BE,
		PixelFormatBGR565LE,
		PixelFormatBGR555BE,
		PixelFormatBGR555LE,
		PixelFormatBGR444BE,
		PixelFormatBGR444LE,
		PixelFormatGBRP16LE,
		PixelFormatGBRP16BE,
		PixelFormatGBRAP16LE,
		PixelFormatGBRAP16BE,
		PixelFormatBayerBGGR16LE,
		PixelFormatBayerBGGR16BE,
		PixelFormatBayerRGGB16LE,
		PixelFormatBayerRGGB16BE,
		PixelFormatBayerGBRG16LE,
		PixelFormatBayerGBRG16BE,
		PixelFormatBayerGRBG16LE,
		PixelFormatBayerGRBG16BE:
		return 16

	case PixelFormatFloat, PixelFormatUint:
		return 32

	default:
		return 8
	}
}

const (
	VideoCodecUndefined      trimmer.VideoCodec = ""
	VideoCodecNone           trimmer.VideoCodec = "none"
	VideoCodecAVC            trimmer.VideoCodec = "avc"
	VideoCodecAVCIntra       trimmer.VideoCodec = "avcintra"
	VideoCodecAVCProxy       trimmer.VideoCodec = "avcproxy"
	VideoCodecAVCLongG       trimmer.VideoCodec = "avclongg"
	VideoCodecHEVC           trimmer.VideoCodec = "hevc"
	VideoCodecVP8            trimmer.VideoCodec = "vp8"
	VideoCodecVP9            trimmer.VideoCodec = "vp9"
	VideoCodecMPEG2          trimmer.VideoCodec = "mpeg2"
	VideoCodecMPEG4          trimmer.VideoCodec = "mpeg4"
	VideoCodecProres422Proxy trimmer.VideoCodec = "prores422proxy"
	VideoCodecProres422Lt    trimmer.VideoCodec = "prores422lt"
	VideoCodecProres422      trimmer.VideoCodec = "prores422"
	VideoCodecProres422hq    trimmer.VideoCodec = "prores422hq"
	VideoCodecProres4444     trimmer.VideoCodec = "prores4444"
	VideoCodecProres4444xq   trimmer.VideoCodec = "prores4444xq"
	VideoCodecDNxHD          trimmer.VideoCodec = "dnxhd"
	VideoCodecDNxHR          trimmer.VideoCodec = "dnxhr"
	VideoCodecArriraw        trimmer.VideoCodec = "arriraw"
	VideoCodecRedCode        trimmer.VideoCodec = "redcode"
	VideoCodecJPEG           trimmer.VideoCodec = "jpeg"
	VideoCodecJPEG2000       trimmer.VideoCodec = "jpeg2000"
	VideoCodecDNG            trimmer.VideoCodec = "dng"
	VideoCodecDPX            trimmer.VideoCodec = "dpx"
	VideoCodecTIFF           trimmer.VideoCodec = "tiff"
	VideoCodecEXR            trimmer.VideoCodec = "exr"
	VideoCodecPNG            trimmer.VideoCodec = "png"
)

func ParseVideoCodec(s string) trimmer.VideoCodec {
	switch s {
	case "none":
		return VideoCodecNone
	case "avc":
		return VideoCodecAVC
	case "avcintra":
		return VideoCodecAVCIntra
	case "avcproxy":
		return VideoCodecAVCProxy
	case "avclongg":
		return VideoCodecAVCLongG
	case "hevc":
		return VideoCodecHEVC
	case "vp8":
		return VideoCodecVP8
	case "vp9":
		return VideoCodecVP9
	case "mpeg2":
		return VideoCodecMPEG2
	case "mpeg4":
		return VideoCodecMPEG4
	case "prores422proxy":
		return VideoCodecProres422Proxy
	case "prores422lt":
		return VideoCodecProres422Lt
	case "prores422":
		return VideoCodecProres422
	case "prores422hq":
		return VideoCodecProres422hq
	case "prores4444":
		return VideoCodecProres4444
	case "prores4444xq":
		return VideoCodecProres4444xq
	case "dnxhd":
		return VideoCodecDNxHD
	case "dnxhr":
		return VideoCodecDNxHR
	case "arriraw":
		return VideoCodecArriraw
	case "redcode":
		return VideoCodecRedCode
	case "jpeg":
		return VideoCodecJPEG
	case "jpeg2000":
		return VideoCodecJPEG2000
	case "dng":
		return VideoCodecDNG
	case "dpx":
		return VideoCodecDPX
	case "tiff":
		return VideoCodecTIFF
	case "exr":
		return VideoCodecEXR
	case "png":
		return VideoCodecPNG
	default:
		return VideoCodecUndefined
	}
}

const (
	AudioLayoutUndefined     trimmer.AudioLayout = ""
	AudioLayoutNone          trimmer.AudioLayout = "none"
	AudioLayoutMono          trimmer.AudioLayout = "mono"
	AudioLayoutStereo        trimmer.AudioLayout = "stereo"
	AudioLayout2_1           trimmer.AudioLayout = "2.1"      // FL+FR+LFE
	AudioLayout3_0           trimmer.AudioLayout = "3.0"      // FL+FR+FC
	AudioLayout3_0Back       trimmer.AudioLayout = "3.0b"     // FL+FR+BC
	AudioLayout4_0           trimmer.AudioLayout = "4.0"      // FL+FR+FC+BC
	AudioLayoutQuad          trimmer.AudioLayout = "quad"     // FL+FR+BL+BR
	AudioLayoutQuadSide      trimmer.AudioLayout = "quads"    // FL+FR+SL+SR
	AudioLayout3_1           trimmer.AudioLayout = "3.1"      // FL+FR+FC+LFE
	AudioLayout5_0           trimmer.AudioLayout = "5.0"      // FL+FR+FC+BL+BR
	AudioLayout5_0Side       trimmer.AudioLayout = "5.0s"     // FL+FR+FC+SL+SR
	AudioLayout4_1           trimmer.AudioLayout = "4.1"      // FL+FR+FC+LFE+BC
	AudioLayout5_1           trimmer.AudioLayout = "5.1"      // FL+FR+FC+LFE+BL+BR
	AudioLayout5_1Side       trimmer.AudioLayout = "5.1s"     // FL+FR+FC+LFE+SL+SR
	AudioLayout6_0           trimmer.AudioLayout = "6.0"      // FL+FR+FC+BC+SL+SR
	AudioLayout6_0Front      trimmer.AudioLayout = "6.0f"     // FL+FR+FLC+FRC+SL+SR
	AudioLayoutHexagonal     trimmer.AudioLayout = "hexa"     // FL+FR+FC+BL+BR+BC
	AudioLayout6_1           trimmer.AudioLayout = "6.1"      // FL+FR+FC+LFE+BC+SL+SR
	AudioLayout6_1Back       trimmer.AudioLayout = "6.1b"     // FL+FR+FC+LFE+BL+BR+BC
	AudioLayout6_1Front      trimmer.AudioLayout = "6.1f"     // FL+FR+LFE+FLC+FRC+SL+SR
	AudioLayout7_0           trimmer.AudioLayout = "7.0"      // FL+FR+FC+BL+BR+SL+SR
	AudioLayout7_0Front      trimmer.AudioLayout = "7.0f"     // FL+FR+FC+FLC+FRC+SL+SR
	AudioLayout7_1           trimmer.AudioLayout = "7.1"      // FL+FR+FC+LFE+BL+BR+SL+SR
	AudioLayout7_1Wide       trimmer.AudioLayout = "7.1w"     // FL+FR+FC+LFE+BL+BR+FLC+FRC
	AudioLayout7_1WideSide   trimmer.AudioLayout = "7.1ws"    // FL+FR+FC+LFE+FLC+FRC+SL+SR
	AudioLayoutOctagonal     trimmer.AudioLayout = "octa"     // FL+FR+FC+BL+BR+BC+SL+SR
	AudioLayoutHexadecagonal trimmer.AudioLayout = "hexadeca" // FL+FR+FC+BL+BR+BC+SL+SR+TFL+TFC+TFR+TBL+TBC+TBR+WL+WR
	AudioLayoutDownmix       trimmer.AudioLayout = "downmix"  // DL+DR
)

func ParseAudioLayout(s string) trimmer.AudioLayout {
	switch s {
	case "none":
		return AudioLayoutNone
	case "mono":
		return AudioLayoutMono
	case "stereo":
		return AudioLayoutStereo
	case "2.1":
		return AudioLayout2_1
	case "3.0":
		return AudioLayout3_0
	case "3.0b":
		return AudioLayout3_0Back
	case "4.0":
		return AudioLayout4_0
	case "quad":
		return AudioLayoutQuad
	case "quads":
		return AudioLayoutQuadSide
	case "3.1":
		return AudioLayout3_1
	case "5.0":
		return AudioLayout5_0
	case "5.0s":
		return AudioLayout5_0Side
	case "4.1":
		return AudioLayout4_1
	case "5.1":
		return AudioLayout5_1
	case "5.1s":
		return AudioLayout5_1Side
	case "6.0":
		return AudioLayout6_0
	case "6.0f":
		return AudioLayout6_0Front
	case "hexa":
		return AudioLayoutHexagonal
	case "6.1":
		return AudioLayout6_1
	case "6.1b":
		return AudioLayout6_1Back
	case "6.1f":
		return AudioLayout6_1Front
	case "7.0":
		return AudioLayout7_0
	case "7.0f":
		return AudioLayout7_0Front
	case "7.1":
		return AudioLayout7_1
	case "7.1w":
		return AudioLayout7_1Wide
	case "7.1ws":
		return AudioLayout7_1WideSide
	case "octa":
		return AudioLayoutOctagonal
	case "hexadeca":
		return AudioLayoutHexadecagonal
	case "downmix":
		return AudioLayoutDownmix
	default:
		return AudioLayoutUndefined
	}
}

const (
	AudioCodecUndefined trimmer.AudioCodec = ""
	AudioCodecNone      trimmer.AudioCodec = "none"
	AudioCodecPCM       trimmer.AudioCodec = "pcm"
	AudioCodecAAC       trimmer.AudioCodec = "aac"
	AudioCodecAC3       trimmer.AudioCodec = "ac3"
	AudioCodecMP3       trimmer.AudioCodec = "mp3"
	AudioCodecFLAC      trimmer.AudioCodec = "flac"
	AudioCodecVorbis    trimmer.AudioCodec = "vorbis"
	AudioCodecOpus      trimmer.AudioCodec = "opus"
)

func ParseAudioCodec(s string) trimmer.AudioCodec {
	switch s {
	case "none":
		return AudioCodecNone
	case "pcm":
		return AudioCodecPCM
	case "aac":
		return AudioCodecAAC
	case "ac3":
		return AudioCodecAC3
	case "mp3":
		return AudioCodecMP3
	case "flac":
		return AudioCodecFLAC
	case "vorbis":
		return AudioCodecVorbis
	case "opus":
		return AudioCodecOpus
	default:
		return AudioCodecUndefined
	}
}

const (
	SubtitleCodecUndefined    trimmer.SubtitleCodec = ""
	SubtitleCodecNone         trimmer.SubtitleCodec = "none"
	SubtitleCodec3GPP         trimmer.SubtitleCodec = "3gpp"
	SubtitleCodecBluRay       trimmer.SubtitleCodec = "bluray"
	SubtitleCodecDCI          trimmer.SubtitleCodec = "dci"
	SubtitleCodecDVB          trimmer.SubtitleCodec = "dvb"
	SubtitleCodecDVD          trimmer.SubtitleCodec = "dvd"
	SubtitleCodecSMPTE20521TT trimmer.SubtitleCodec = "smpte_2052_1_tt"
	SubtitleCodecSRT          trimmer.SubtitleCodec = "srt"
	SubtitleCodecTTML         trimmer.SubtitleCodec = "ttml"
	SubtitleCodecWebVTT       trimmer.SubtitleCodec = "webvtt"
	SubtitleCodecEBUTT        trimmer.SubtitleCodec = "ebu_tt"
)

func ParseSubtitleCodec(s string) trimmer.SubtitleCodec {
	switch s {
	case "none":
		return SubtitleCodecNone
	case "3gpp":
		return SubtitleCodec3GPP
	case "bluray":
		return SubtitleCodecBluRay
	case "dci":
		return SubtitleCodecDCI
	case "dvb":
		return SubtitleCodecDVB
	case "dvd":
		return SubtitleCodecDVD
	case "smpte_2052_1_tt":
		return SubtitleCodecSMPTE20521TT
	case "srt":
		return SubtitleCodecSRT
	case "ttml":
		return SubtitleCodecTTML
	case "webvtt":
		return SubtitleCodecWebVTT
	case "ebu_tt":
		return SubtitleCodecEBUTT
	default:
		return SubtitleCodecUndefined
	}
}

const (
	CDLModeUndefined trimmer.CDLMode = ""
	CDLModeSource    trimmer.CDLMode = "source" // Arri: "CDL LogC"
	CDLModeTarget    trimmer.CDLMode = "target" // Arri: "CDL Video"
)

const (
	CropNone         trimmer.CropMode = ""
	CropStored       trimmer.CropMode = "stored"
	CropSampled      trimmer.CropMode = "sampled"
	CropDisplay      trimmer.CropMode = "display"
	CropCenterCut43  trimmer.CropMode = "centercut43"  // SMPTE ST 2067-2:2013 Annex G
	CropCenterCut169 trimmer.CropMode = "centercut169" // SMPTE ST 2067-2:2013 Annex G
)

const (
	FlipNone       trimmer.FlipMode = ""
	FlipHorizontal trimmer.FlipMode = "horizontal"
	FlipVertical   trimmer.FlipMode = "vertical"
	FlipBoth       trimmer.FlipMode = "both"
)

const (
	RotateNone                  trimmer.RotateMode = ""
	RotateClockPortrait         trimmer.RotateMode = "portrait_clock"   // rotate clock-wise into portrait mode if necessary
	RotateClockLandscape        trimmer.RotateMode = "landscape_clock"  // rotate clock-wise into landscape mode if necessary
	RotateCounterClockPortrait  trimmer.RotateMode = "portrait_cclock"  // rotate counter-clock-wise into portrait mode if necessary
	RotateCounterClockLandscape trimmer.RotateMode = "landscape_cclock" // rotate counter-clock-wise into landscape mode if necessary
	RotateCounterClock          trimmer.RotateMode = "cclock"           // 90 deg counter-clock-wise
	RotateClock                 trimmer.RotateMode = "clock"            // 90 deg clock-wise
	RotateCounterClockAndFlip   trimmer.RotateMode = "cclock_flip"      // 90 deg counter-clock-wise + vflip
	RotateClockAndFlip          trimmer.RotateMode = "clock_flip"       // 90 deg clock-wise + vflip
)

const (
	WatermarkModeUndefined trimmer.WatermarkMode = ""
	WatermarkModeLogo      trimmer.WatermarkMode = "logo"
	WatermarkModeTimecode  trimmer.WatermarkMode = "timecode"
	WatermarkModeText      trimmer.WatermarkMode = "text"
	WatermarkModeMetadata  trimmer.WatermarkMode = "meta"
)
