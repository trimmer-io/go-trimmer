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

package trimmer

import (
	"net/url"
	"reflect"
)

// Query is the function used to get a page listing.
type Query func(url.Values) ([]interface{}, ListMeta, error)

// Iter provides a convenient interface
// for iterating over the elements
// returned from paginated list API calls.
// Successive calls to the Next method
// will step through each item in the list,
// fetching pages of items as needed.
// Iterators are not thread-safe, so they should not be consumed
// across multiple goroutines.
type Iter struct {
	query  Query
	qs     url.Values
	values []interface{}
	meta   ListMeta
	params ListParams
	err    error
	cur    interface{}
}

// GetIter returns a new Iter for a given query and its options.
func GetIter(params *ListParams, qs *url.Values, query Query) *Iter {
	iter := &Iter{}
	iter.query = query

	p := params
	if p == nil {
		p = &ListParams{}
	}
	iter.params = *p

	q := qs
	if q == nil {
		q = &url.Values{}
	}
	iter.qs = *q

	iter.getPage()
	return iter
}

func GetIterErr(err error) *Iter {
	iter := &Iter{}
	iter.err = err
	return iter
}

func (it *Iter) getPage() {
	it.values, it.meta, it.err = it.query(it.qs)

	// when moving backwards strip off the first item because maxId is inclusive
	if it.params.MaxId != "" && len(it.values) > 0 {
		it.values = it.values[1:]
	}

	if len(it.values) > 0 {
		it.meta.More = true
	}

	// update total count
	it.meta.Total += len(it.values)

	// Reverse order when moving forward, because items arrive in backward order.
	// 	reverse(it.values)
	// }
}

// Next advances the Iter to the next item in the list,
// which will then be available when calling Current()
// Next() returns false when the iterator stops
// at the end of the list.
func (it *Iter) Next() bool {
	if len(it.values) == 0 && it.meta.More {
		// determine if we're moving forward or backwards in paging
		// if it.params.MinId != "" {
		// 	it.params.MinId = listItemID(it.cur)
		// 	it.qs.Set(maxId, it.params.MinId)
		// } else {
		// 	it.params.MaxId = listItemID(it.cur)
		// 	it.qs.Set(minId, it.params.MaxId)
		// }

		// backward order: set the new maxId from the current page's last (minId) item
		it.params.MaxId = it.meta.MinId
		it.qs.Set(maxId, it.params.MaxId)
		it.getPage()
	}
	if len(it.values) == 0 {
		return false
	}
	it.cur = it.values[0]

	// uhhh, this strips from the slice!
	it.values = it.values[1:]
	return true
}

// Current returns the most recent item
// visited by a call to Next.
func (it *Iter) Current() interface{} {
	return it.cur
}

// Err returns the error, if any,
// that caused the Iter to stop.
// It must be inspected
// after Next returns false.
func (it *Iter) Err() error {
	return it.err
}

// Meta returns the list metadata.
func (it *Iter) Meta() *ListMeta {
	return &it.meta
}

func listItemID(x interface{}) string {
	return reflect.ValueOf(x).Elem().FieldByName("ID").String()
}

func reverse(a []interface{}) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}
