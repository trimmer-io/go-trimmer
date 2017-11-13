// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package trimmer

import (
	"database/sql/driver"
	"strings"
)

type ApiEmbedFlags int64

const (
	API_EMBED_UNDEFINED = 0
	API_EMBED_MEDIA     = 1 << (iota - 1)
	API_EMBED_STATS
	API_EMBED_WORKSPACE
	API_EMBED_OWNER
	API_EMBED_AUTHOR
	API_EMBED_ORIGIN
	API_EMBED_META
	API_EMBED_PERMS
	API_EMBED_DETAILS
	API_EMBED_URLS
	API_EMBED_ALL = 0x3FF
)

func ParseApiEmbedFlag(s string) ApiEmbedFlags {
	switch s {
	case "media":
		return API_EMBED_MEDIA
	case "stats":
		return API_EMBED_STATS
	case "workspace":
		return API_EMBED_WORKSPACE
	case "owner":
		return API_EMBED_OWNER
	case "author":
		return API_EMBED_AUTHOR
	case "origin":
		return API_EMBED_ORIGIN
	case "meta":
		return API_EMBED_META
	case "perms":
		return API_EMBED_PERMS
	case "details":
		return API_EMBED_DETAILS
	case "urls":
		return API_EMBED_URLS
	case "all":
		return API_EMBED_ALL
	default:
		return API_EMBED_UNDEFINED
	}
}

func PrintApiEmbedFlag(f ApiEmbedFlags) string {
	switch f {
	case API_EMBED_MEDIA:
		return "media"
	case API_EMBED_STATS:
		return "stats"
	case API_EMBED_WORKSPACE:
		return "workspace"
	case API_EMBED_OWNER:
		return "owner"
	case API_EMBED_AUTHOR:
		return "author"
	case API_EMBED_ORIGIN:
		return "origin"
	case API_EMBED_META:
		return "meta"
	case API_EMBED_PERMS:
		return "perms"
	case API_EMBED_DETAILS:
		return "details"
	case API_EMBED_URLS:
		return "urls"
	case API_EMBED_ALL:
		return "all"
	default:
		return ""
	}
}

func (f ApiEmbedFlags) IsValid() bool {
	return f != API_EMBED_UNDEFINED
}

func (f ApiEmbedFlags) Contains(b ApiEmbedFlags) bool {
	return f&b == b
}

func (f ApiEmbedFlags) Complement(b ApiEmbedFlags) bool {
	return f&^b > 0
}

func ParseApiEmbedFlags(s string) (ApiEmbedFlags, error) {
	var flags ApiEmbedFlags
	for _, v := range strings.Split(s, ",") {
		f := ParseApiEmbedFlag(strings.ToLower(v))
		if f == API_EMBED_UNDEFINED {
			return 0, EParamInvalid
		}
		flags |= f
	}
	return flags, nil
}

func (f ApiEmbedFlags) String() string {
	s := make([]string, 0)
	for i := 1; i < API_EMBED_ALL; i = i << 1 {
		flag := ApiEmbedFlags(i)
		if f.Contains(flag) {
			s = append(s, PrintApiEmbedFlag(flag))
		}
	}
	return strings.Join(s, ",")
}

func (f ApiEmbedFlags) MarshalText() ([]byte, error) {
	if f == 0 {
		return []byte{}, nil
	}
	return []byte(f.String()), nil
}

func (f *ApiEmbedFlags) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		*f = 0
		return nil
	}
	if flags, err := ParseApiEmbedFlags(string(data)); err != nil {
		return err
	} else {
		*f = flags
	}
	return nil
}

// SQL conversion
func (f *ApiEmbedFlags) Scan(value interface{}) error {
	switch v := value.(type) {
	case int, int64:
		*f = ApiEmbedFlags(v.(int64))
	case string:
		if flags, err := ParseApiEmbedFlags(v); err != nil {
			return err
		} else {
			*f = flags
		}

	case []byte:
		if flags, err := ParseApiEmbedFlags(string(v)); err != nil {
			return err
		} else {
			*f = flags
		}
	}
	return nil
}

func (f ApiEmbedFlags) Value() (driver.Value, error) {
	return int64(f), nil
}
