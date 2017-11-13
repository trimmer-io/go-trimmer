// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

// Package replica provides helpers for media replicas
package replica

import (
	trimmer "trimmer.io/go-trimmer"
)

// Iter is an iterator for lists of Replicas.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// Replica returns the most recent Replica visited by a call to Next.
func (i *Iter) Replica() *trimmer.Replica {
	return i.Current().(*trimmer.Replica)
}
