// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package trimmer

import (
	"time"
)

// OrgState is the list of allowed values for the organization account status.
// Allowed values are "parked", "created", "active", "inactive", "expired"
// "blocked", "banned", "deleting", "cleaning", "deleted".
type OrgState string

// OrgListParams is the set of parameters that can be used when listing organizations.
type OrgListParams struct {
	ListParams
	Embed ApiEmbedFlags `json:"embed,omitempty"`
}

// OrgParams is the set of parameters that can be used when reading, creating or
// updating organizations.
type OrgParams struct {
	DisplayName string        `json:"displayName,omitempty"`
	ImageId     string        `json:"imageId,omitempty"`
	Homepage    string        `json:"homepage,omitempty"`
	Location    string        `json:"location,omitempty"`
	Description string        `json:"description,omitempty"`
	Embed       ApiEmbedFlags `json:"embed,omitempty"`
}

// Org is the resource representing a Trimmer organization account.
type Org struct {
	ID               string         `json:"orgId"`
	Name             string         `json:"name"`
	DisplayName      string         `json:"displayName"`
	State            OrgState       `json:"state"`
	Location         string         `json:"location"`
	Description      string         `json:"description"`
	Homepage         string         `json:"homepage"`
	ImageKey         string         `json:"imageKey"`
	Language         string         `json:"language"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	PaymentAt        time.Time      `json:"paymentAt"`
	AccountPlan      string         `json:"plan"`
	BillingId        string         `json:"billingId"`
	BillingAccountId string         `json:"billingAccountId"`
	IsAnnual         bool           `json:"isAnnual"`
	IsTrial          bool           `json:"isTrial"`
	Quantity         int            `json:"quantity"`
	Media            *MediaEmbed    `json:"media"`
	Membership       *Member        `json:"membership"`
	Permissions      *Permissions   `json:"perms"`
	Statistics       *OrgStatistics `json:"stats"`
}

type OrgList []*Org

func (l OrgList) SearchId(id string) (int, *Org) {
	for i, v := range l {
		if v.ID == id {
			return i, v
		}
	}
	return len(l), nil
}

// OrgStatistics is the sub-resource representing organization counters.
type OrgStatistics struct {
	Workspaces int64 `json:"workspaces"`
	Volumes    int64 `json:"volumes"`
	Members    int64 `json:"members"`
	Teams      int64 `json:"teams"`
	Assets     int64 `json:"assets"`
	Time       int64 `json:"time"`
	Media      int64 `json:"media"`
	Size       int64 `json:"size"`
	Files      int64 `json:"files"`
}
