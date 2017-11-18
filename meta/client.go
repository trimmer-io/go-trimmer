// Trimmer SDK
//
// Copyright (c) 2017 Alexander Eichhorn
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
