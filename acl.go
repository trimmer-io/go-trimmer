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

package trimmer

const (
	PERM_NONE    = 0
	PERM_READ    = 0x1
	PERM_COPY    = 0x2
	PERM_WRITE   = 0x4
	PERM_CREATE  = 0x8
	PERM_DELETE  = 0x10
	PERM_ADMIN   = 0x20
	PERM_RELEASE = 0x40
	PERM_CHOWN   = 0x80
	PERM_MAX     = 0xFF
	// mask for user-settable permissions
	PERM_MASK_USER   = 0x1F
	PERM_MASK_SYSTEM = 0xE0
)

const (
	ACCESS_INVALID  AccessClass = ""
	ACCESS_PUBLIC   AccessClass = "public"
	ACCESS_PRIVATE  AccessClass = "private"
	ACCESS_PERSONAL AccessClass = "personal"
)

type AccessClass string

// permission bitmask
type Permission int

// permission struct carrying multiple permission bitmasks
type Permissions struct {
	Self    Permission `json:"self"`    // caller permissions
	Private Permission `json:"private"` // group permissions
	Public  Permission `json:"public"`  // public permissions
}

type AclObjectType string
type AclSubjectType string
