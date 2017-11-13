// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

// Package link provides utility functions for asset/stash links
package link

import (
	trimmer "trimmer.io/go-trimmer"
)

// Iter is an iterator for lists of Links.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// Link returns the most recent User visited by a call to Next.
func (i *Iter) Link() *trimmer.Link {
	return i.Current().(*trimmer.Link)
}
