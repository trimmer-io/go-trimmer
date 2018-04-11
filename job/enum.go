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

package job

import (
	trimmer "trimmer.io/go-trimmer"
)

const (
	JobStateUndefined trimmer.JobState = ""
	JobStateActive    trimmer.JobState = "active" // used for listing states only
	JobStateDone      trimmer.JobState = "done"   // used for listing states only
	JobStateCreated   trimmer.JobState = "created"
	JobStateQueued    trimmer.JobState = "queued"
	JobStateRunning   trimmer.JobState = "running"
	JobStateComplete  trimmer.JobState = "complete"
	JobStateFailed    trimmer.JobState = "failed"
	JobStateAborted   trimmer.JobState = "aborted"
)

func ParseJobState(s string) trimmer.JobState {
	switch s {
	case "active":
		return JobStateActive
	case "done":
		return JobStateDone
	case "created":
		return JobStateCreated
	case "queued":
		return JobStateQueued
	case "running":
		return JobStateRunning
	case "complete":
		return JobStateComplete
	case "failed":
		return JobStateFailed
	case "aborted":
		return JobStateAborted
	default:
		return JobStateUndefined
	}
}

const (
	JobTypeUndefined trimmer.JobType = ""
	JobTypeUpload    trimmer.JobType = "upload"    // client-side upload
	JobTypeCopy      trimmer.JobType = "copy"      // copy media bewteen volumes
	JobTypeScan      trimmer.JobType = "scan"      // volume scan
	JobTypeWipe      trimmer.JobType = "wipe"      // volume wipe
	JobTypeWatch     trimmer.JobType = "watch"     // volume watch
	JobTypeUnwatch   trimmer.JobType = "unwatch"   // volume unwatch
	JobTypeAnalyze   trimmer.JobType = "analyze"   // transcoder analyze media
	JobTypeTranscode trimmer.JobType = "transcode" // transcode media
	JobTypeRender    trimmer.JobType = "render"    // render editlist into media
	JobTypeEMail     trimmer.JobType = "email"     // send email(s)
	JobTypeInvoice   trimmer.JobType = "invoice"   // create invoice PDF
	JobTypeReceipt   trimmer.JobType = "receipt"   // create receipt PDF
	JobTypePublish   trimmer.JobType = "publish"   // publishing job
)

func ParseJobType(s string) trimmer.JobType {
	switch s {
	case "upload":
		return JobTypeUpload
	case "copy":
		return JobTypeCopy
	case "scan":
		return JobTypeScan
	case "wipe":
		return JobTypeWipe
	case "watch":
		return JobTypeWatch
	case "unwatch":
		return JobTypeUnwatch
	case "analyze":
		return JobTypeAnalyze
	case "transcode":
		return JobTypeTranscode
	case "render":
		return JobTypeRender
	case "email":
		return JobTypeEMail
	case "invoice":
		return JobTypeInvoice
	case "receipt":
		return JobTypeReceipt
	case "publish":
		return JobTypePublish
	default:
		return JobTypeUndefined
	}
}
