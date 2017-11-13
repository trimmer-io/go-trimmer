// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

// Package mount provides helpers for the workspace/volumes mounting API
package mount

import (
	trimmer "trimmer.io/go-trimmer"
)

// Iter is an iterator for lists of Mounts.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// Mount returns the most recent Mount visited by a call to Next.
func (i *Iter) Mount() *trimmer.Mount {
	return i.Current().(*trimmer.Mount)
}
