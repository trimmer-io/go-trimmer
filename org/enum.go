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
