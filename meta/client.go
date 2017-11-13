// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

// Package meta provides the /assets/:id/meta APIs
package meta

import (
	trimmer "trimmer.io/go-trimmer"
)

// Iter is an iterator for lists of Mmetadata versions.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// Version returns the most recent metadata version visited by a call to Next.
func (i *Iter) Version() *trimmer.MetaVersion {
	return i.Current().(*trimmer.MetaVersion)
}
