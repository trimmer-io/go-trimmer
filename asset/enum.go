// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package asset

import (
	trimmer "trimmer.io/go-trimmer"
)

const (
	AssetStateUndefined   trimmer.AssetState = ""
	AssetStateNomedia     trimmer.AssetState = "nomedia"
	AssetStateUploading   trimmer.AssetState = "uploading"
	AssetStateAnalyzing   trimmer.AssetState = "analyzing"
	AssetStateTranscoding trimmer.AssetState = "transcoding"
	AssetStateAttention   trimmer.AssetState = "attention"
	AssetStateReady       trimmer.AssetState = "ready"
	AssetStateBlocked     trimmer.AssetState = "blocked"
	AssetStateBanned      trimmer.AssetState = "banned"
	AssetStateDeleting    trimmer.AssetState = "deleting"
	AssetStateCleaning    trimmer.AssetState = "cleaning"
	AssetStateDeleted     trimmer.AssetState = "deleted"
)

func ParseAssetState(s string) trimmer.AssetState {
	switch s {
	case "nomedia":
		return AssetStateNomedia
	case "uploading":
		return AssetStateUploading
	case "analyzing":
		return AssetStateAnalyzing
	case "transcoding":
		return AssetStateTranscoding
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
