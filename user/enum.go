// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

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
