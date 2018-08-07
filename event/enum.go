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

package event

import (
	trimmer "trimmer.io/go-trimmer"
)

const (
	EventTypeUndefined    trimmer.EventType = ""
	EventTypeAccount      trimmer.EventType = "account"
	EventTypeAcl          trimmer.EventType = "acl"
	EventTypeApplication  trimmer.EventType = "application"
	EventTypeAsset        trimmer.EventType = "asset"
	EventTypeAction       trimmer.EventType = "action"
	EventTypeAuth         trimmer.EventType = "auth"
	EventTypeBilling      trimmer.EventType = "billing"
	EventTypeMedia        trimmer.EventType = "media"
	EventTypeInvite       trimmer.EventType = "invite"
	EventTypeJob          trimmer.EventType = "job"
	EventTypeStash        trimmer.EventType = "stash"
	EventTypeTag          trimmer.EventType = "tag"
	EventTypeWorkspace    trimmer.EventType = "workspace"
	EventTypeTeam         trimmer.EventType = "team"
	EventTypeVolume       trimmer.EventType = "volume"
	EventTypeOrganization trimmer.EventType = "organization"
)

func ParseEventType(s string) trimmer.EventType {
	switch s {
	case "account":
		return EventTypeAccount
	case "acl":
		return EventTypeAcl
	case "application":
		return EventTypeApplication
	case "asset":
		return EventTypeAsset
	case "auth":
		return EventTypeAuth
	case "action":
		return EventTypeAction
	case "billing":
		return EventTypeBilling
	case "media":
		return EventTypeMedia
	case "invite":
		return EventTypeInvite
	case "job":
		return EventTypeJob
	case "stash":
		return EventTypeStash
	case "tag":
		return EventTypeTag
	case "workspace":
		return EventTypeWorkspace
	case "team":
		return EventTypeTeam
	case "volume":
		return EventTypeVolume
	case "organization":
		return EventTypeOrganization
	default:
		return EventTypeUndefined
	}
}

const (
	EventKeyUndefined trimmer.EventKey = ""

	EventKeyAccountCreated           trimmer.EventKey = "account.created"
	EventKeyAccountActivated         trimmer.EventKey = "account.activated"
	EventKeyAccountDeactivated       trimmer.EventKey = "account.deactivated"
	EventKeyAccountBlocked           trimmer.EventKey = "account.blocked"
	EventKeyAccountBanned            trimmer.EventKey = "account.banned"
	EventKeyAccountDeletionInitiated trimmer.EventKey = "account.deletion.initiated"
	EventKeyAccountCleanupCompleted  trimmer.EventKey = "account.cleanup.completed"
	EventKeyAccountRejected          trimmer.EventKey = "account.rejected"
	EventKeyAccountEmailChanged      trimmer.EventKey = "account.email.changed"
	EventKeyAccountEmailProbed       trimmer.EventKey = "account.email.probed"
	EventKeyAccountEmailVerified     trimmer.EventKey = "account.email.verified"
	EventKeyAccountPlanChanged       trimmer.EventKey = "account.plan.changed"
	EventKeyAccountPlanActivated     trimmer.EventKey = "account.plan.activated"
	EventKeyAccountPlanExpired       trimmer.EventKey = "account.plan.expired"

	EventKeyAuthLoginSuccess        trimmer.EventKey = "login.success"
	EventKeyAuthLoginFailed         trimmer.EventKey = "login.failed"
	EventKeyAuthLogoutSuccess       trimmer.EventKey = "logout.success"
	EventKeyAuthLoginExpired        trimmer.EventKey = "login.expired"
	EventKeyAuthLoginRefresh        trimmer.EventKey = "login.refresh"
	EventKeyAuthPasswordChanged     trimmer.EventKey = "password.changed"
	EventKeyAuthPasswordResetInit   trimmer.EventKey = "password.reset.init"
	EventKeyAuthPasswordResetAction trimmer.EventKey = "password.reset.action"

	EventKeyAclCreated trimmer.EventKey = "acl.created"
	EventKeyAclUpdated trimmer.EventKey = "acl.updated"
	EventKeyAclDeleted trimmer.EventKey = "acl.deleted"

	EventKeyActionCreated trimmer.EventKey = "action.created"
	EventKeyActionUpdated trimmer.EventKey = "action.updated"
	EventKeyActionDeleted trimmer.EventKey = "action.deleted"
	EventKeyActionRun     trimmer.EventKey = "action.run"

	EventKeyApplicationCreated      trimmer.EventKey = "app.created"
	EventKeyApplicationUpdated      trimmer.EventKey = "app.updated"
	EventKeyApplicationRekeyed      trimmer.EventKey = "app.rekeyed"
	EventKeyApplicationDeleted      trimmer.EventKey = "app.deleted"
	EventKeyApplicationTokenCreated trimmer.EventKey = "app.token.created"
	EventKeyApplicationTokenRevoked trimmer.EventKey = "app.token.revoked"

	EventKeyBillingCreated trimmer.EventKey = "billing.created"
	EventKeyBillingUpdated trimmer.EventKey = "billing.updated"

	EventKeyInvoiceCreated   trimmer.EventKey = "invoice.created"
	EventKeyInvoiceSent      trimmer.EventKey = "invoice.sent"
	EventKeyInvoicePayed     trimmer.EventKey = "invoice.payed"
	EventKeyInvoiceFailed    trimmer.EventKey = "invoice.failed"
	EventKeyInvoiceCancelled trimmer.EventKey = "invoice.cancelled"
	EventKeyInvoiceRefunded  trimmer.EventKey = "invoice.refunded"

	EventKeyOrganizationCreated       trimmer.EventKey = "org.created"
	EventKeyOrganizationUpdated       trimmer.EventKey = "org.updated"
	EventKeyOrganizationDeleted       trimmer.EventKey = "org.deleted"
	EventKeyOrganizationMemberAdded   trimmer.EventKey = "org.member.added"
	EventKeyOrganizationMemberUpdated trimmer.EventKey = "org.member.udated"
	EventKeyOrganizationMemberRemoved trimmer.EventKey = "org.member.removed"

	EventKeyWorkspaceCreated       trimmer.EventKey = "workspace.created"
	EventKeyWorkspaceUpdated       trimmer.EventKey = "workspace.updated"
	EventKeyWorkspaceDeleted       trimmer.EventKey = "workspace.deleted"
	EventKeyWorkspaceMemberAdded   trimmer.EventKey = "workspace.member.added"
	EventKeyWorkspaceMemberUpdated trimmer.EventKey = "workspace.member.udated"
	EventKeyWorkspaceMemberRemoved trimmer.EventKey = "workspace.member.removed"

	EventKeyInviteSent     trimmer.EventKey = "invite.sent"
	EventKeyInviteResent   trimmer.EventKey = "invite.resent"
	EventKeyInviteVisited  trimmer.EventKey = "invite.visited"
	EventKeyInviteAccepted trimmer.EventKey = "invite.accepted"
	EventKeyInviteDeclined trimmer.EventKey = "invite.declined"
	EventKeyInviteExpired  trimmer.EventKey = "invite.expired"
	EventKeyInviteRecalled trimmer.EventKey = "invite.recalled"

	EventKeyStashCreated    trimmer.EventKey = "stash.created"
	EventKeyStashDeleted    trimmer.EventKey = "stash.deleted"
	EventKeyStashUpdated    trimmer.EventKey = "stash.updated"
	EventKeyStashTransfered trimmer.EventKey = "stash.transfered"

	EventKeyAssetCreated     trimmer.EventKey = "asset.created"
	EventKeyAssetUpdated     trimmer.EventKey = "asset.updated"
	EventKeyAssetForked      trimmer.EventKey = "asset.forked"
	EventKeyAssetDeleted     trimmer.EventKey = "asset.deleted"
	EventKeyAssetLinked      trimmer.EventKey = "asset.linked"
	EventKeyAssetLinkUpdated trimmer.EventKey = "asset.linkupdated"
	EventKeyAssetUnlinked    trimmer.EventKey = "asset.unlinked"
	EventKeyAssetReviewing   trimmer.EventKey = "asset.reviewing"
	EventKeyAssetAccepted    trimmer.EventKey = "asset.accepted"
	EventKeyAssetRejected    trimmer.EventKey = "asset.rejected"
	EventKeyAssetPublishing  trimmer.EventKey = "asset.publishing"
	EventKeyAssetPublished   trimmer.EventKey = "asset.published"

	EventKeyTagCreated trimmer.EventKey = "tag.created"
	EventKeyTagUpdated trimmer.EventKey = "tag.updated"
	EventKeyTagReplied trimmer.EventKey = "tag.replied"
	EventKeyTagDeleted trimmer.EventKey = "tag.deleted"

	EventKeyTeamCreated trimmer.EventKey = "team.created"
	EventKeyTeamUpdated trimmer.EventKey = "team.updated"
	EventKeyTeamDeleted trimmer.EventKey = "team.deleted"
	EventKeyTeamJoined  trimmer.EventKey = "team.joined"
	EventKeyTeamChanged trimmer.EventKey = "team.changed"
	EventKeyTeamLeft    trimmer.EventKey = "team.left"
	EventKeyTeamGranted trimmer.EventKey = "team.granted"
	EventKeyTeamRevoked trimmer.EventKey = "team.revoked"

	EventKeyVolumeCreated      trimmer.EventKey = "volume.created"
	EventKeyVolumeUpdated      trimmer.EventKey = "volume.updated"
	EventKeyVolumeDeleted      trimmer.EventKey = "volume.deleted"
	EventKeyVolumeMounted      trimmer.EventKey = "volume.mounted"
	EventKeyVolumeUnmounted    trimmer.EventKey = "volume.unmounted"
	EventKeyVolumeScanned      trimmer.EventKey = "volume.scanned"
	EventKeyVolumeCleared      trimmer.EventKey = "volume.cleared"
	EventKeyVolumeWiped        trimmer.EventKey = "volume.wiped"
	EventKeyVolumeWatchStarted trimmer.EventKey = "volume.watch.started"
	EventKeyVolumeWatchStopped trimmer.EventKey = "volume.watch.stopped"

	EventKeyMediaCreated        trimmer.EventKey = "media.created"
	EventKeyMediaUpdated        trimmer.EventKey = "media.updated"
	EventKeyMediaDeleted        trimmer.EventKey = "media.deleted"
	EventKeyMediaLinked         trimmer.EventKey = "media.linked"
	EventKeyMediaUnlinked       trimmer.EventKey = "media.unlinked"
	EventKeyMediaCopied         trimmer.EventKey = "media.copied"
	EventKeyMediaProcessed      trimmer.EventKey = "media.processed"
	EventKeyMediaUploaded       trimmer.EventKey = "media.uploaded"
	EventKeyMediaReplicaCreated trimmer.EventKey = "media.replica.created"
	EventKeyMediaReplicaDeleted trimmer.EventKey = "media.replica.deleted"

	EventKeyJobCreated  trimmer.EventKey = "job.created"
	EventKeyJobQueued   trimmer.EventKey = "job.queued"
	EventKeyJobStarted  trimmer.EventKey = "job.started"
	EventKeyJobRetried  trimmer.EventKey = "job.retried"
	EventKeyJobFailed   trimmer.EventKey = "job.failed"
	EventKeyJobAborted  trimmer.EventKey = "job.aborted"
	EventKeyJobFinished trimmer.EventKey = "job.finished"
)

func ParseEventKey(s string) trimmer.EventKey {
	switch s {
	case "account.created":
		return EventKeyAccountCreated
	case "account.activated":
		return EventKeyAccountActivated
	case "account.deactivated":
		return EventKeyAccountDeactivated
	case "account.blocked":
		return EventKeyAccountBlocked
	case "account.banned":
		return EventKeyAccountBanned
	case "account.deletion.initiated":
		return EventKeyAccountDeletionInitiated
	case "account.cleanup.completed":
		return EventKeyAccountCleanupCompleted
	case "account.rejected":
		return EventKeyAccountRejected
	case "account.email.changed":
		return EventKeyAccountEmailChanged
	case "account.email.probed":
		return EventKeyAccountEmailProbed
	case "account.email.verified":
		return EventKeyAccountEmailVerified
	case "account.plan.changed":
		return EventKeyAccountPlanChanged
	case "account.plan.activated":
		return EventKeyAccountPlanActivated
	case "account.plan.expired":
		return EventKeyAccountPlanExpired

	case "login.success":
		return EventKeyAuthLoginSuccess
	case "login.failed":
		return EventKeyAuthLoginFailed
	case "logout.success":
		return EventKeyAuthLogoutSuccess
	case "login.expired":
		return EventKeyAuthLoginExpired
	case "login.refresh":
		return EventKeyAuthLoginRefresh
	case "password.changed":
		return EventKeyAuthPasswordChanged
	case "password.reset.init":
		return EventKeyAuthPasswordResetInit
	case "password.reset.action":
		return EventKeyAuthPasswordResetAction

	case "app.created":
		return EventKeyApplicationCreated
	case "app.updated":
		return EventKeyApplicationUpdated
	case "app.rekeyed":
		return EventKeyApplicationRekeyed
	case "app.deleted":
		return EventKeyApplicationDeleted
	case "app.token.created":
		return EventKeyApplicationTokenCreated
	case "app.token.revoked":
		return EventKeyApplicationTokenRevoked

	case "billing.created":
		return EventKeyBillingCreated
	case "billing.updated":
		return EventKeyBillingUpdated
	case "invoice.created":
		return EventKeyInvoiceCreated
	case "invoice.sent":
		return EventKeyInvoiceSent
	case "invoice.payed":
		return EventKeyInvoicePayed
	case "invoice.failed":
		return EventKeyInvoiceFailed
	case "invoice.cancelled":
		return EventKeyInvoiceCancelled
	case "invoice.refunded":
		return EventKeyInvoiceRefunded

	case "acl.created":
		return EventKeyAclCreated
	case "acl.updated":
		return EventKeyAclUpdated
	case "acl.deleted":
		return EventKeyAclDeleted

	case "action.created":
		return EventKeyActionCreated
	case "action.updated":
		return EventKeyActionUpdated
	case "action.deleted":
		return EventKeyActionDeleted
	case "action.run":
		return EventKeyActionRun

	case "org.created":
		return EventKeyOrganizationCreated
	case "org.updated":
		return EventKeyOrganizationUpdated
	case "org.deleted":
		return EventKeyOrganizationDeleted
	case "org.member.added":
		return EventKeyOrganizationMemberAdded
	case "org.member.updated":
		return EventKeyOrganizationMemberUpdated
	case "org.member.removed":
		return EventKeyOrganizationMemberRemoved

	case "workspace.created":
		return EventKeyWorkspaceCreated
	case "workspace.updated":
		return EventKeyWorkspaceUpdated
	case "workspace.deleted":
		return EventKeyWorkspaceDeleted
	case "workspace.member.added":
		return EventKeyWorkspaceMemberAdded
	case "workspace.member.updated":
		return EventKeyWorkspaceMemberUpdated
	case "workspace.member.removed":
		return EventKeyWorkspaceMemberRemoved

	case "invite.sent":
		return EventKeyInviteSent
	case "invite.resent":
		return EventKeyInviteResent
	case "invite.visited":
		return EventKeyInviteVisited
	case "invite.accepted":
		return EventKeyInviteAccepted
	case "invite.declined":
		return EventKeyInviteDeclined
	case "invite.expired":
		return EventKeyInviteExpired
	case "invite.recalled":
		return EventKeyInviteRecalled

	case "stash.created":
		return EventKeyStashCreated
	case "stash.deleted":
		return EventKeyStashDeleted
	case "stash.updated":
		return EventKeyStashUpdated
	case "stash.transfered":
		return EventKeyStashTransfered

	case "asset.created":
		return EventKeyAssetCreated
	case "asset.updated":
		return EventKeyAssetUpdated
	case "asset.forked":
		return EventKeyAssetForked
	case "asset.deleted":
		return EventKeyAssetDeleted
	case "asset.linked":
		return EventKeyAssetLinked
	case "asset.unlinked":
		return EventKeyAssetUnlinked
	case "asset.linkupdated":
		return EventKeyAssetLinkUpdated
	case "asset.reviewing":
		return EventKeyAssetReviewing
	case "asset.accepted":
		return EventKeyAssetAccepted
	case "asset.rejected":
		return EventKeyAssetRejected
	case "asset.publishing":
		return EventKeyAssetPublishing
	case "asset.published":
		return EventKeyAssetPublished

	case "tag.created":
		return EventKeyTagCreated
	case "tag.updated":
		return EventKeyTagUpdated
	case "tag.replied":
		return EventKeyTagReplied
	case "tag.deleted":
		return EventKeyTagDeleted

	case "team.created":
		return EventKeyTeamCreated
	case "team.updated":
		return EventKeyTeamUpdated
	case "team.deleted":
		return EventKeyTeamDeleted
	case "team.joined":
		return EventKeyTeamJoined
	case "team.changed":
		return EventKeyTeamChanged
	case "team.left":
		return EventKeyTeamLeft
	case "team.granted":
		return EventKeyTeamGranted
	case "team.revoked":
		return EventKeyTeamRevoked

	case "volume.created":
		return EventKeyVolumeCreated
	case "volume.updated":
		return EventKeyVolumeUpdated
	case "volume.deleted":
		return EventKeyVolumeDeleted
	case "volume.mounted":
		return EventKeyVolumeMounted
	case "volume.unmounted":
		return EventKeyVolumeUnmounted
	case "volume.scanned":
		return EventKeyVolumeScanned
	case "volume.cleared":
		return EventKeyVolumeCleared
	case "volume.wiped":
		return EventKeyVolumeWiped
	case "volume.watch.started":
		return EventKeyVolumeWatchStarted
	case "volume.watch.stopped":
		return EventKeyVolumeWatchStopped

	case "media.created":
		return EventKeyMediaCreated
	case "media.updated":
		return EventKeyMediaUpdated
	case "media.deleted":
		return EventKeyMediaDeleted
	case "media.linked":
		return EventKeyMediaLinked
	case "media.unlinked":
		return EventKeyMediaUnlinked
	case "media.copied":
		return EventKeyMediaCopied
	case "media.processed":
		return EventKeyMediaProcessed
	case "media.uploaded":
		return EventKeyMediaUploaded
	case "media.replica.created":
		return EventKeyMediaReplicaCreated
	case "media.replica.deleted":
		return EventKeyMediaReplicaDeleted

	case "job.created":
		return EventKeyJobCreated
	case "job.queued":
		return EventKeyJobQueued
	case "job.started":
		return EventKeyJobStarted
	case "job.retried":
		return EventKeyJobRetried
	case "job.failed":
		return EventKeyJobFailed
	case "job.aborted":
		return EventKeyJobAborted
	case "job.finished":
		return EventKeyJobFinished

	default:
		return EventKeyUndefined
	}
}
