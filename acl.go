// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

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
