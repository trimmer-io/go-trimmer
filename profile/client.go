// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

// Package profile provides helpers for transcode profiles
package profile

import (
	trimmer "trimmer.io/go-trimmer"
)

// Iter is an iterator for lists of Profiles.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// Profile returns the most recent Profile visited by a call to Next.
func (i *Iter) Profile() *trimmer.Profile {
	return i.Current().(*trimmer.Profile)
}
