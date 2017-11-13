// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package tag

import (
	trimmer "trimmer.io/go-trimmer"
)

const (
	TagLabelUndefined    trimmer.TagLabel = ""
	TagLabelBlack        trimmer.TagLabel = "black"
	TagLabelChapter      trimmer.TagLabel = "chapter"
	TagLabelComment      trimmer.TagLabel = "comment"
	TagLabelEvent        trimmer.TagLabel = "event"
	TagLabelFavorite     trimmer.TagLabel = "favorite"
	TagLabelFrameData    trimmer.TagLabel = "framedata"
	TagLabelKeyword      trimmer.TagLabel = "keyword"
	TagLabelLds          trimmer.TagLabel = "lds"
	TagLabelNote         trimmer.TagLabel = "note"
	TagLabelOrganization trimmer.TagLabel = "organization"
	TagLabelPerson       trimmer.TagLabel = "person"
	TagLabelPlace        trimmer.TagLabel = "place"
	TagLabelProduct      trimmer.TagLabel = "product"
	TagLabelQuote        trimmer.TagLabel = "quote"
	TagLabelRating       trimmer.TagLabel = "rating"
	TagLabelScenecut     trimmer.TagLabel = "scenecut"
	TagLabelSilence      trimmer.TagLabel = "silence"
	TagLabelTask         trimmer.TagLabel = "task"
	TagLabelText         trimmer.TagLabel = "text"
)

func ParseTagLabel(s string) trimmer.TagLabel {
	switch s {
	case "black":
		return TagLabelBlack
	case "chapter":
		return TagLabelChapter
	case "comment":
		return TagLabelComment
	case "event":
		return TagLabelEvent
	case "favorite":
		return TagLabelFavorite
	case "framedata":
		return TagLabelFrameData
	case "keyword":
		return TagLabelKeyword
	case "lds":
		return TagLabelLds
	case "note":
		return TagLabelNote
	case "organization":
		return TagLabelOrganization
	case "person":
		return TagLabelPerson
	case "place":
		return TagLabelPlace
	case "product":
		return TagLabelProduct
	case "quote":
		return TagLabelQuote
	case "rating":
		return TagLabelRating
	case "scenecut":
		return TagLabelScenecut
	case "silence":
		return TagLabelSilence
	case "task":
		return TagLabelTask
	case "text":
		return TagLabelText
	default:
		return TagLabelUndefined
	}
}
