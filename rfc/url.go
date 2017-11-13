// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//
package rfc

import (
	"net/url"
	"path"
	"strings"
)

func Basename(s string) string {
	if u, err := url.Parse(s); err == nil {
		return path.Base(u.Path)
	}
	f := strings.Split(strings.Split(strings.Split(s, "?")[0], "#")[0], "/")
	return f[len(f)-1]
}
