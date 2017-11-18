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

package workspace

import (
	trimmer "trimmer.io/go-trimmer"
)

const (
	WorkspaceStateUndefined trimmer.WorkspaceState = ""
	WorkspaceStateInactive  trimmer.WorkspaceState = "inactive"
	WorkspaceStateActive    trimmer.WorkspaceState = "active"
	WorkspaceStateBlocked   trimmer.WorkspaceState = "blocked"
	WorkspaceStateBanned    trimmer.WorkspaceState = "banned"
	WorkspaceStateDeleting  trimmer.WorkspaceState = "deleting"
	WorkspaceStateCleaning  trimmer.WorkspaceState = "cleaning"
	WorkspaceStateDeleted   trimmer.WorkspaceState = "deleted"
)

func ParseWorkspaceState(s string) trimmer.WorkspaceState {
	switch s {
	case "inactive":
		return WorkspaceStateInactive
	case "active":
		return WorkspaceStateActive
	case "blocked":
		return WorkspaceStateBlocked
	case "banned":
		return WorkspaceStateBanned
	case "deleting":
		return WorkspaceStateDeleting
	case "cleaning":
		return WorkspaceStateCleaning
	case "deleted":
		return WorkspaceStateDeleted
	default:
		return WorkspaceStateUndefined
	}
}
