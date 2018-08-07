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

package asset

import (
	trimmer "trimmer.io/go-trimmer"
)

const (
	AssetStateUndefined  trimmer.AssetState = ""
	AssetStateEmpty      trimmer.AssetState = "empty"
	AssetStateUploading  trimmer.AssetState = "uploading"
	AssetStateProcessing trimmer.AssetState = "processing"
	AssetStatePublishing trimmer.AssetState = "publishing"
	AssetStatePublished  trimmer.AssetState = "published"
	AssetStateReviewing  trimmer.AssetState = "reviewing"
	AssetStateApproved   trimmer.AssetState = "approved"
	AssetStateRejected   trimmer.AssetState = "rejected"
	AssetStateArchived   trimmer.AssetState = "archived"
	AssetStateAttention  trimmer.AssetState = "attention"
	AssetStateReady      trimmer.AssetState = "ready"
	AssetStateBlocked    trimmer.AssetState = "blocked"
	AssetStateBanned     trimmer.AssetState = "banned"
	AssetStateDeleting   trimmer.AssetState = "deleting"
	AssetStateCleaning   trimmer.AssetState = "cleaning"
	AssetStateDeleted    trimmer.AssetState = "deleted"
)

func ParseAssetState(s string) trimmer.AssetState {
	switch s {
	case "empty":
		return AssetStateEmpty
	case "uploading":
		return AssetStateUploading
	case "processing":
		return AssetStateProcessing
	case "publishing":
		return AssetStatePublishing
	case "published":
		return AssetStatePublished
	case "reviewing":
		return AssetStateReviewing
	case "approved":
		return AssetStateApproved
	case "rejected":
		return AssetStateRejected
	case "archived":
		return AssetStateArchived
	case "attention":
		return AssetStateAttention
	case "ready":
		return AssetStateReady
	case "blocked":
		return AssetStateBlocked
	case "banned":
		return AssetStateBanned
	case "deleting":
		return AssetStateDeleting
	case "cleaning":
		return AssetStateCleaning
	case "deleted":
		return AssetStateDeleted
	default:
		return AssetStateUndefined
	}
}

// AssetListEvent
const (
	AssetListEventUndefined trimmer.AssetListEvent = ""
	AssetListEventCreated   trimmer.AssetListEvent = "created"
	AssetListEventUpdated   trimmer.AssetListEvent = "updated"
)

func ParseAssetListEvent(s string) trimmer.AssetListEvent {
	switch s {
	case "created":
		return AssetListEventCreated
	case "updated":
		return AssetListEventUpdated
	default:
		return AssetListEventUndefined
	}
}
