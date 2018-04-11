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

package workflow

import (
	trimmer "trimmer.io/go-trimmer"
)

const (
	ActionUndefined         trimmer.WorkflowAction = ""
	ActionNoop              trimmer.WorkflowAction = "noop"
	ActionMediaAnalyze      trimmer.WorkflowAction = "media.analyze"
	ActionMediaTranscode    trimmer.WorkflowAction = "media.transcode"
	ActionMediaRender       trimmer.WorkflowAction = "media.render"
	ActionMediaDemux        trimmer.WorkflowAction = "media.demux"
	ActionMediaMux          trimmer.WorkflowAction = "media.mux"
	ActionMediaCopy         trimmer.WorkflowAction = "media.copy"
	ActionMediaUnlink       trimmer.WorkflowAction = "media.unlink"
	ActionMediaSnapshot     trimmer.WorkflowAction = "media.snapshot"
	ActionMetadataSet       trimmer.WorkflowAction = "metadata.set"
	ActionMetadataOverwrite trimmer.WorkflowAction = "metadata.overwrite"
	ActionMetadataClear     trimmer.WorkflowAction = "metadata.clear"
	ActionVolumeScan        trimmer.WorkflowAction = "volume.scan"
	ActionVolumeWipe        trimmer.WorkflowAction = "volume.wipe"
	ActionVolumeWatch       trimmer.WorkflowAction = "volume.watch"
	ActionVolumeUnwatch     trimmer.WorkflowAction = "volume.unwatch"
	ActionAccessGrant       trimmer.WorkflowAction = "access.grant"
	ActionAccessRevoke      trimmer.WorkflowAction = "access.revoke"
	ActionNotifyEmail       trimmer.WorkflowAction = "notify.email"
	ActionNotifyWebhook     trimmer.WorkflowAction = "notify.webhook"
	ActionNotifySms         trimmer.WorkflowAction = "notify.sms"
	ActionAssetCheck        trimmer.WorkflowAction = "asset.check"
	ActionAssetPublish      trimmer.WorkflowAction = "asset.publish"
	ActionAssetMerge        trimmer.WorkflowAction = "asset.merge"
)

func ParseWorkflowAction(s string) WorkflowAction {
	switch s {
	case "noop":
		return WorkflowActionNoop
	case "media.analyze":
		return WorkflowActionMediaAnalyze
	case "media.transcode":
		return WorkflowActionMediaTranscode
	case "media.render":
		return WorkflowActionMediaRender
	case "media.demux":
		return WorkflowActionMediaDemux
	case "media.mux":
		return WorkflowActionMediaMux
	case "media.copy":
		return WorkflowActionMediaCopy
	case "media.unlink":
		return WorkflowActionMediaUnlink
	case "media.snapshot":
		return WorkflowActionMediaSnapshot
	case "metadata.set":
		return WorkflowActionMetadataSet
	case "metadata.overwrite":
		return WorkflowActionMetadataOverwrite
	case "metadata.clear":
		return WorkflowActionMetadataClear
	case "volume.scan":
		return WorkflowActionVolumeScan
	case "volume.wipe":
		return WorkflowActionVolumeWipe
	case "volume.watch":
		return WorkflowActionVolumeWatch
	case "volume.unwatch":
		return WorkflowActionVolumeUnwatch
	case "access.grant":
		return WorkflowActionAccessGrant
	case "access.revoke":
		return WorkflowActionAccessRevoke
	case "notify.email":
		return WorkflowActionNotifyEmail
	case "notify.webhook":
		return WorkflowActionNotifyWebhook
	case "notify.sms":
		return WorkflowActionNotifySms
	case "asset.check":
		return WorkflowActionAssetCheck
	case "asset.publish":
		return WorkflowActionAssetPublish
	case "asset.merge":
		return WorkflowActionAssetMerge
	default:
		return WorkflowActionUndefined
	}
}
