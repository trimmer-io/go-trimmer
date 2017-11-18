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

package trimmer

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// MetaUpdateParams is the set of parameters that can be used to update asset metadata.
//
type MetaUpdateParams struct {
	Actions MetaValueList `json:"actions"`
	Version string        `json:"version"`
	Embed   ApiEmbedFlags `json:"embed,omitempty"`
}

// MetaQueryParams is the set of parameters that can be used to query asset metadata.
//
type MetaQueryParams struct {
	Filter  string        `json:"-"`
	Version string        `json:"-"`
	Embed   ApiEmbedFlags `json:"embed,omitempty"`
}

// MetaDiffParams is the set of parameters that can be used to generate a
// unified diff between two versions.
//
type MetaDiffParams struct {
	V1 string
	V2 string
}

// MetaListParams is the set of parameters that can be used when listing assets.
type MetaListParams struct {
	ListParams
	AuthorId string        `json:"authorId,omitempty"`
	Embed    ApiEmbedFlags `json:"embed,omitempty"`
}

type MetaValue struct {
	Path  MetaPath  `json:"path"`
	Value string    `json:"value"`
	Flags MetaFlags `json:"flags"`
}

type MetaValueList []MetaValue

type MetaDocument struct {
	Namespaces map[string]string      `json:"namespaces"`
	Models     map[string]interface{} `json:"models"`
}

type MetaVersion struct {
	VersionId   string        `json:"versionId"`
	Hash        string        `json:"hash"`
	AssetId     string        `json:"assetId"`
	AuthorId    string        `json:"authorId"`
	WorkspaceId string        `json:"workspaceId"`
	AccessClass AccessClass   `json:"access"`
	CreatedAt   time.Time     `json:"createdAt"`
	IsFragment  bool          `json:"isFragment"`
	Comment     string        `json:"comment"`
	Action      string        `json:"action"`
	Changed     string        `json:"changed"`
	Metadata    *MetaDocument `json:"metadata"`
	Workspace   *Workspace    `json:"workspace"`
	Author      *User         `json:"author"`
}

type MetaVersionList []*MetaVersion

func (x *MetaUpdateParams) Add(path MetaPath, value string) {
	x.Actions = append(x.Actions, MetaValue{
		Path:  path,
		Value: value,
	})
}

func (x *MetaUpdateParams) AddWithFlags(path MetaPath, value string, flags MetaFlags) {
	x.Actions = append(x.Actions, MetaValue{
		Path:  path,
		Value: value,
	})
}

func (d *MetaDocument) GetPath(path MetaPath) (string, error) {
	if !path.IsValid() {
		return "", NewUsageError(fmt.Sprintf("invalid meta path '%s'", string(path)), nil)
	}
	ns := path.Namespace()
	if _, ok := d.Namespaces[ns]; !ok {
		return "", NewUsageError(fmt.Sprintf("undefined meta namespace '%s' in document", string(path)), nil)
	}
	v, ok := d.Models[ns]
	if !ok {
		return "", NewUsageError(fmt.Sprintf("missing meta model '%s' in document", string(path)), nil)
	}
	val, ok := v.(map[string]interface{})
	for _, fname := range path.Fields() {
		// split lang or array index from name, be safe with slice indexes
		var lang string
		var idx int
		if k := strings.Index(fname, "["); k > 0 && len(fname) > k+2 {
			s := strings.TrimSuffix(fname[k+1:], "]")
			if j, err := strconv.Atoi(s); err == nil {
				idx = j
			} else {
				lang = s
			}
			fname = fname[:k]
		}
		fname = ns + ":" + fname
		v, ok = val[fname]
		if !ok {
			return "", NewUsageError(fmt.Sprintf("missing field '%s' in document", string(path)), nil)
		}
		switch x := v.(type) {
		case map[string]interface{}:
			val = x
		// case bool:
		// case float64:
		case string:
			return x, nil
		case []interface{}:
			// there's two types of lists in XMP: normal and alternative
			if isAltArrayType(x) {
				for _, av := range x {
					alt := av.(map[string]interface{})
					if lang != "" && alt["lang"].(string) == lang {
						return alt["value"].(string), nil
					}
					if lang == "" && alt["isDefault"].(bool) == true {
						return alt["value"].(string), nil
					}
				}
				return "", NewUsageError(fmt.Sprintf("missing field '%s' in document: no such language", string(path)), nil)
			} else {
				if len(x) >= idx {
					return "", NewUsageError(fmt.Sprintf("missing field '%s' in document: index out of range", string(path)), nil)
				}
				v = x[idx]
				switch x := v.(type) {
				case map[string]interface{}:
					val = x
				case string:
					return x, nil
				case []interface{}:
					NewUsageError(fmt.Sprintf("unsupported nested array at '%s' in document", string(path)), nil)
				}
			}
		default:
			NewUsageError(fmt.Sprintf("unsupported JSON/XMP content type at '%s' in document", string(path)), nil)
		}
	}
	return "", NewUsageError(fmt.Sprintf("missing field '%s' in document", string(path)), nil)
}

func isObjectValue(v interface{}) bool {
	switch v.(type) {
	case map[string]interface{}:
		return true
	default:
		return false
	}
}

func isAltArrayItemType(v interface{}) bool {
	// alternative array item
	if isObjectValue(v) {
		val := v.(map[string]interface{})
		if len(val) > 3 {
			return false
		}
		if _, ok := val["value"]; !ok {
			return false
		}
		if _, ok := val["lang"]; !ok {
			return false
		}
		if _, ok := val["isDefault"]; !ok {
			return false
		}
		return true
	}
	return false
}

func isAltArrayType(v interface{}) bool {
	slice, ok := v.([]interface{})
	if !ok || len(slice) == 0 {
		return false
	}
	if isAltArrayItemType(slice[0]) {
		return true
	}
	return false
}

type MetaPath string
type MetaGroup string

func (x MetaPath) IsValid() bool {
	return strings.Index(string(x), ":") > -1
}

func (x MetaPath) Namespace() string {
	if i := strings.Index(string(x), ":"); i > -1 {
		return string(x[:i])
	}
	return string(x)
}

func (x MetaPath) MatchNamespace(ns string) bool {
	if ns == "" {
		return false
	}
	if i := strings.Index(string(x), ":"); i > -1 {
		return ns == string(x[:i])
	}
	return false
}

func (x MetaPath) Length() int {
	return strings.Count(string(x), "/") + 1
}

func (x MetaPath) Fields() []string {
	s := string(x)
	if i := strings.Index(s, ":"); i > -1 {
		return strings.Split(s[i+1:], "/")
	}
	return nil
}

type MetaFlags int

const (
	META_NOFLAG = 0
	META_CREATE = 1 << (iota - 1)
	META_REPLACE
	META_DELETE
	META_APPEND
	META_UNIQUE
	META_NOFAIL
	META_DEFAULT = META_CREATE | META_REPLACE | META_DELETE | META_UNIQUE
	META_MERGE   = META_CREATE | META_REPLACE | META_UNIQUE | META_NOFAIL
	META_ADD     = META_CREATE | META_UNIQUE | META_NOFAIL
)

func PrintMetaFlag(f MetaFlags) string {
	switch f {
	case META_CREATE:
		return "create"
	case META_REPLACE:
		return "replace"
	case META_DELETE:
		return "delete"
	case META_APPEND:
		return "append"
	case META_UNIQUE:
		return "unique"
	case META_NOFAIL:
		return "nofail"
	default:
		return ""
	}
}

func (f MetaFlags) IsValid() bool {
	return f != META_NOFLAG
}

func (f MetaFlags) Contains(b MetaFlags) bool {
	return f&b == b
}

func (f MetaFlags) Complement(b MetaFlags) bool {
	return f&^b > 0
}

func (f MetaFlags) String() string {
	s := make([]string, 0)
	for i := 1; i <= META_UNIQUE; i = i << 1 {
		flag := MetaFlags(i)
		if f.Contains(flag) {
			s = append(s, PrintMetaFlag(flag))
		}
	}
	return strings.Join(s, ",")
}

func (f MetaFlags) MarshalText() ([]byte, error) {
	if f == 0 {
		return []byte{}, nil
	}
	return []byte(f.String()), nil
}
