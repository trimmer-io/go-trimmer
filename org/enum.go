// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

// Package org provides the /orgs APIs
package org

import (
	trimmer "trimmer.io/go-trimmer"
)

const (
	OrgStateParked   trimmer.OrgState = "parked"
	OrgStateCreated  trimmer.OrgState = "created"
	OrgStateActive   trimmer.OrgState = "active"
	OrgStateInactive trimmer.OrgState = "inactive"
	OrgStateExpired  trimmer.OrgState = "expired"
	OrgStateBlocked  trimmer.OrgState = "blocked"
	OrgStateBanned   trimmer.OrgState = "banned"
	OrgStateDeleting trimmer.OrgState = "deleting"
	OrgStateCleaning trimmer.OrgState = "cleaning"
	OrgStateDeleted  trimmer.OrgState = "deleted"
)
