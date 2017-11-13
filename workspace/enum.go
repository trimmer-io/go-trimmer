// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

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
