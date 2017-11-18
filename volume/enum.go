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

package volume

import (
	trimmer "trimmer.io/go-trimmer"
)

const (
	VolumeReadonlyStateUndefined trimmer.VolumeReadonlyState = ""
	VolumeReadonlyStateOn        trimmer.VolumeReadonlyState = "on"
	VolumeReadonlyStateOff       trimmer.VolumeReadonlyState = "off"
)

const (
	VolumeOnlineStateUndefined trimmer.VolumeOnlineState = ""
	VolumeOnlineStateOn        trimmer.VolumeOnlineState = "on"
	VolumeOnlineStateOff       trimmer.VolumeOnlineState = "off"
)

const (
	VolumeAutomountStateUndefined trimmer.VolumeAutomountState = ""
	VolumeAutomountStateOn        trimmer.VolumeAutomountState = "on"
	VolumeAutomountStateOff       trimmer.VolumeAutomountState = "off"
)

const (
	VolumeStateReady      trimmer.VolumeState = "ready"
	VolumeStateFailed     trimmer.VolumeState = "failed"
	VolumeStateScanning   trimmer.VolumeState = "scanning"
	VolumeStateLoading    trimmer.VolumeState = "loading"
	VolumeStateOffloading trimmer.VolumeState = "offloading"
	VolumeStateWiping     trimmer.VolumeState = "wiping"
	VolumeStateWatching   trimmer.VolumeState = "watching"
	VolumeStateStopping   trimmer.VolumeState = "stopping"
	VolumeStateTransit    trimmer.VolumeState = "transit"
	VolumeStateLost       trimmer.VolumeState = "lost"
	VolumeStateArchived   trimmer.VolumeState = "archived"
	VolumeStateRetired    trimmer.VolumeState = "retired"
)

const (
	VolumeTypeClient  trimmer.VolumeType = "client"  // all volumes on local discs at clients
	VolumeTypeCloud   trimmer.VolumeType = "cloud"   // cloud storage, object stores
	VolumeTypeShuttle trimmer.VolumeType = "shuttle" // removable shuttle drives
	VolumeTypeNAS     trimmer.VolumeType = "nas"     // local network servers
	VolumeTypeSAN     trimmer.VolumeType = "san"     // local storage libraries
	VolumeTypeTape    trimmer.VolumeType = "tape"    // removable tapes
	VolumeTypeCDN     trimmer.VolumeType = "cdn"     // content-delivery network
)

const (
	VolumeProviderNone           trimmer.VolumeProvider = ""
	VolumeProviderTrimmerCloud   trimmer.VolumeProvider = "tcloud"
	VolumeProviderTrimmerGateway trimmer.VolumeProvider = "tgate"
	VolumeProviderAWS            trimmer.VolumeProvider = "s3"
)

const (
	VolumeAuthTypeNone      trimmer.VolumeAuthType = "none"      // no auth
	VolumeAuthTypeSignature trimmer.VolumeAuthType = "signature" // Trimmer url signature
	VolumeAuthTypeToken     trimmer.VolumeAuthType = "token"     // Trimmer signed token
	VolumeAuthTypeAWS       trimmer.VolumeAuthType = "aws"       // Amazon IAM user
	VolumeAuthTypeOAuth2    trimmer.VolumeAuthType = "oauth2"    // OAuth2 authorization code grant
	VolumeAuthTypeBasic     trimmer.VolumeAuthType = "basic"     // user:password-based auth
)

const (
	VolumeAuthScopeNone   trimmer.VolumeAuthScope = "none"
	VolumeAuthScopeRead   trimmer.VolumeAuthScope = "read"
	VolumeAuthScopeCreate trimmer.VolumeAuthScope = "create"
	VolumeAuthScopeDelete trimmer.VolumeAuthScope = "delete"
)
