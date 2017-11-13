// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

// Package member provides utilities for managing members
package member

import (
	trimmer "trimmer.io/go-trimmer"
)

// Iter is an iterator for lists of Members.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// Member returns the most recent User visited by a call to Next.
func (i *Iter) Member() *trimmer.Member {
	return i.Current().(*trimmer.Member)
}
