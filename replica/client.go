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
