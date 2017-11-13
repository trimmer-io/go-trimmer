// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package trimmer

import (
	"strings"
	"time"
	"trimmer.io/go-trimmer/hash"
)

// VolumeState is the list of allowed values for states. Allowed states are
// "init", "ready", "failed", "scanning", "loading", "offloading", "wiping",
// "stopping", "transit", "lost", "archived", "retired".
type VolumeState string
type VolumeStateList []VolumeState

func (l VolumeStateList) Contains(s VolumeState) bool {
	if len(l) == 0 {
		return false
	}
	for _, v := range l {
		if v == s {
			return true
		}
	}
	return false
}

func (l *VolumeStateList) Add(s VolumeState) {
	for !l.Contains(s) {
		*l = append(*l, s)
	}
}

func (l *VolumeStateList) Del(s VolumeState) {
	i := -1
	for j, v := range *l {
		if v == s {
			i = j
		}
	}
	if i > -1 {
		(*l)[i] = (*l)[len(*l)-1]
		*l = (*l)[:len(*l)-1]
	}
}

func (l VolumeStateList) String() string {
	if len(l) == 0 {
		return ""
	}
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = string(v)
	}
	return strings.Join(s, ",")
}

// VolumeType is the list of allowed values for types of volumes.
// Allowed values are "client", "cloud", "shuttle", "nas", "san", "tape".
type VolumeType string
type VolumeTypeList []VolumeType

func (l VolumeTypeList) Contains(s VolumeType) bool {
	if len(l) == 0 {
		return false
	}
	for _, v := range l {
		if v == s {
			return true
		}
	}
	return false
}

func (l *VolumeTypeList) Add(s VolumeType) {
	for !l.Contains(s) {
		*l = append(*l, s)
	}
}

func (l *VolumeTypeList) Del(s VolumeType) {
	i := -1
	for j, v := range *l {
		if v == s {
			i = j
		}
	}
	if i > -1 {
		(*l)[i] = (*l)[len(*l)-1]
		*l = (*l)[:len(*l)-1]
	}
}

func (l VolumeTypeList) String() string {
	if len(l) == 0 {
		return ""
	}
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = string(v)
	}
	return strings.Join(s, ",")
}

// VolumeProvider is the type of a cloud storage provider supported by Trimmer.
type VolumeProvider string

// VolumeNamespace is the list of allowed values for filename conventions.
// Allowed values are "sha256", "plain", "plain-sha256", "uuid", "uuid-sha256",
// "custom".
type VolumeNamespace string

// VolumeAuthScope is the list of allowed values for operations that require
// authentication. Allowed values are "read", "create", "delete".
type VolumeAuthScope string

// VolumeAuthType is the list of allowed values for authentication schemes.
// Allowed values are "none", "signature", "token", "aws", "oauth2", "basic".
type VolumeAuthType string

// VolumeAuthType is the list of allowed values for readonly status.
// Allowed values are "on", "off".
type VolumeReadonlyState string

// VolumeAutomountState is the list of allowed values for automount status.
// Allowed values are "on", "off".
type VolumeAutomountState string

// VolumeOnlineState is the list of allowed values for online status.
// Allowed values are "on", "off".
type VolumeOnlineState string

// VolumeListParams is the set of parameters that can be used when listing volumes.
type VolumeListParams struct {
	ListParams
	AccessClass AccessClass          `json:"access,omitempty"`
	States      VolumeStateList      `json:"state,omitempty"`
	Name        string               `json:"name,omitempty"`
	UUID        string               `json:"uuid,omitempty"`
	Serial      string               `json:"serial,omitempty"`
	Types       VolumeTypeList       `json:"type,omitempty"`
	Provider    VolumeProvider       `json:"provider,omitempty"`
	Region      string               `json:"region,omitempty"`
	Brand       string               `json:"brand,omitempty"`
	Readonly    VolumeReadonlyState  `json:"readonly,omitempty"`
	Automount   VolumeAutomountState `json:"automount,omitempty"`
	Online      VolumeOnlineState    `json:"online,omitempty"`
	Embed       ApiEmbedFlags        `json:"embed,omitempty"`
}

// VolumeParams is the set of parameters that can be used to create and
// update a volume.
//
type VolumeParams struct {
	Name            string               `json:"name"`            // required
	Type            VolumeType           `json:"type"`            // required
	Namespace       VolumeNamespace      `json:"namespace"`       // required
	DisplayName     string               `json:"displayName"`     // default: = name
	State           VolumeState          `json:"state"`           // update-only
	UUID            string               `json:"uuid"`            // default: read from manifest, create
	SerialNo        string               `json:"serial"`          // default: empty
	Provider        string               `json:"provider"`        // default: trimmer
	Brand           string               `json:"brand"`           // default: empty
	Region          string               `json:"region"`          // default: by cloud provider
	Url             string               `json:"url"`             // default: empty, reqired for type cloud
	Template        string               `json:"template"`        // default: empty
	Readonly        VolumeReadonlyState  `json:"readonly"`        // default: true
	Online          VolumeOnlineState    `json:"online"`          // default: true
	Automount       VolumeAutomountState `json:"automount"`       // default: false
	HashTypes       hash.HashTypeList    `json:"hashTypes"`       // default: md5,sha256
	AuthType        VolumeAuthType       `json:"authType"`        // default: none
	AuthScope       VolumeAuthScope      `json:"authScope"`       // default: write
	AuthCredentials string               `json:"authCredentials"` // default: empty
	RoleMatch       MediaRoleMatch       `json:"roleMatch"`       // default: empty
	AccessClass     AccessClass          `json:"access"`          // default: public
	CacheControl    string               `json:"cacheControl"`    // default: public
	Capacity        int64                `json:"capacity"`        // default: 0 = unlimited
	Embed           ApiEmbedFlags        `json:"embed"`
}

// VolumeScanParams is the set of parameters that can be used when scanning
// a volume.
//
type VolumeScanParams struct {
}

// VolumeClearParams is the set of parameters that can be used when clearing
// a volume.
//
type VolumeClearParams struct {
	Wipe bool `json:"wipe,omitempty"`
}

// Volume is the resource representing a Trimmer volume.
type Volume struct {
	ID           string            `json:"volumeId"`
	UUID         string            `json:"uuid"`
	AccountId    string            `json:"accountId"`
	AuthorId     string            `json:"authorId"`
	State        VolumeState       `json:"state"`
	Name         string            `json:"name"`
	DisplayName  string            `json:"displayName"`
	Type         VolumeType        `json:"type"`
	Brand        string            `json:"brand"`
	Provider     string            `json:"provider"`
	Region       string            `json:"region"`
	Url          string            `json:"url"`
	SerialNo     string            `json:"serial"`
	Namespace    VolumeNamespace   `json:"namespace"`
	Template     string            `json:"template"`
	Online       bool              `json:"online"`
	Readonly     bool              `json:"readonly"`
	Automount    bool              `json:"automount"`
	Revision     int               `json:"revision"`
	HashTypes    hash.HashTypeList `json:"hashTypes"`
	AuthType     VolumeAuthType    `json:"authType"`
	AuthScope    VolumeAuthScope   `json:"authScope"`
	RoleMatch    MediaRoleMatch    `json:"roleMatch"`
	AccessClass  AccessClass       `json:"access"`
	CacheControl string            `json:"cacheControl"`
	Capacity     int64             `json:"capacity"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
	Error        *TrimmerError     `json:"error"`
	Limits       *VolumeLimits     `json:"limits"`
	Statistics   *VolumeStatistics `json:"stats"`
	Permissions  *Permissions      `json:"perms"`
	User         *User             `json:"user"`
	Org          *Org              `json:"org"`
	Author       *User             `json:"author"`
}

type VolumeList []*Volume

func (l VolumeList) SearchId(id string) (int, *Volume) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}

func (l VolumeList) SearchUUID(u string) (int, *Volume) {
	for i, v := range l {
		if v.UUID == u {
			return i, v
		}
	}
	return len(l), nil
}

type VolumeManifest struct {
	UUID      string            `json:"uuid"`
	Name      string            `json:"name"`
	UrlBase   string            `json:"url"`
	UrlPrefix string            `json:"prefix"`
	Namespace VolumeNamespace   `json:"namespace"`
	Readonly  bool              `json:"readonly"`
	Auth      VolumeAuth        `json:"auth"`
	Limits    *VolumeLimits     `json:"limits,omitempty"`
	HashTypes hash.HashTypeList `json:"hashes"`
}

type VolumeAuth struct {
	Type  VolumeAuthType  `json:"type"`
	Scope VolumeAuthScope `json:"scope"`
}

type VolumeLimits struct {
	PartSizeMin   int64 `json:"minPartSize"`
	PartSizeMax   int64 `json:"maxPartSize"`
	PartsMax      int64 `json:"maxParts"`
	FileSizeMax   int64 `json:"maxFileSize"`
	SinglePartMax int64 `json:"maxSinglePart"`
}

type VolumeStatistics struct {
	Totals     *VolumeCounts            `json:"totals"`
	Accounts   map[string]*VolumeCounts `json:"accounts"`
	Workspaces map[string]*VolumeCounts `json:"workspaces"`
	Types      map[string]*VolumeCounts `json:"mediaTypes"`
	Relations  map[string]*VolumeCounts `json:"mediaRelations"`
	Families   map[string]*VolumeCounts `json:"mediaFamilies"`
	Roles      map[string]*VolumeCounts `json:"mediaRoles"`
}

type VolumeCounts struct {
	NumFiles      int64  `json:"numFiles"`
	NumBytes      int64  `json:"usedBytes"`
	NumUploads    int64  `json:"activeUploads"`
	FreeBytes     *int64 `json:"freeBytes"`
	CapacityBytes *int64 `json:"capacityBytes"`
}
