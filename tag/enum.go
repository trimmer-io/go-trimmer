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
