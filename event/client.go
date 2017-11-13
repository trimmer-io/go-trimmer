// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

// Package event provides utilities for audit events
package event

import (
	trimmer "trimmer.io/go-trimmer"
)

// Iter is an iterator for lists of Events.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*trimmer.Iter
}

// Event returns the most recent Event visited by a call to Next.
func (i *Iter) Event() *trimmer.Event {
	return i.Current().(*trimmer.Event)
}
