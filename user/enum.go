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

package user

import (
	trimmer "trimmer.io/go-trimmer"
)

const (
	UserStateParked   trimmer.UserState = "parked"
	UserStateCreated  trimmer.UserState = "created"
	UserStateInvited  trimmer.UserState = "invited"
	UserStateActive   trimmer.UserState = "active"
	UserStateInactive trimmer.UserState = "inactive"
	UserStateExpired  trimmer.UserState = "expired"
	UserStateBlocked  trimmer.UserState = "blocked"
	UserStateBanned   trimmer.UserState = "banned"
	UserStateDeleting trimmer.UserState = "deleting"
	UserStateCleaning trimmer.UserState = "cleaning"
	UserStateDeleted  trimmer.UserState = "deleted"
	UserStateRejected trimmer.UserState = "rejected"
)

const (
	UserSearchDefault     trimmer.UserSearchFields = "default"
	UserSearchAll         trimmer.UserSearchFields = "all"
	UserSearchName        trimmer.UserSearchFields = "name"
	UserSearchDisplayName trimmer.UserSearchFields = "displayName"
	UserSearchEmail       trimmer.UserSearchFields = "email"
	UserSearchLocation    trimmer.UserSearchFields = "location"
	UserSearchDescription trimmer.UserSearchFields = "description"
)
