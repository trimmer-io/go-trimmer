// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

// Package hash provides helper functions for file hashing
package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"database/sql/driver"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"hash"
	"io"
	"strings"

	tiger "trimmer.io/go-trimmer/hash/go-tiger"
	"trimmer.io/go-trimmer/hash/xxhash"
)

type HashType string
type HashTypeList []HashType
type HashFlags int

type HashBlock struct { // = 36[prefixes:] + 328[sum(hash digets)] + 5[limit;] = 369
	Md5    string `json:"md5,omitempty"`    // 128bit, 32 digits + 4
	Sha1   string `json:"sha1,omitempty"`   // 160bit, 40 digits + 5
	Sha256 string `json:"sha256,omitempty"` // 256bit, 64 digits + 7
	Sha512 string `json:"sha512,omitempty"` // 512bit, 128 digits + 7
	XXHash string `json:"xxhash,omitempty"` // 64bit, 16 digits + 7
	Tiger  string `json:"tiger,omitempty"`  // 192bit, 48 digits + 6

	md5    hash.Hash
	sha1   hash.Hash
	sha256 hash.Hash
	sha512 hash.Hash
	xxhash hash.Hash64
	tiger  hash.Hash
}

const (
	HASH_TYPE_INVALID = 0
	HASH_TYPE_MD5     = 1 << iota
	HASH_TYPE_SHA1
	HASH_TYPE_SHA256
	HASH_TYPE_SHA512
	HASH_TYPE_XXHASH
	HASH_TYPE_TIGER
)

var (
	HashTypeInvalid HashType = ""
	HashTypeMd5     HashType = "md5"
	HashTypeSha1    HashType = "sha1"
	HashTypeSha256  HashType = "sha256"
	HashTypeSha512  HashType = "sha512"
	HashTypeXxhash  HashType = "xxhash"
	HashTypeTiger   HashType = "tiger"

	DefaultHash = HashTypeSha256

	HashTypesAll HashTypeList = HashTypeList{
		HashTypeMd5,
		HashTypeSha1,
		HashTypeSha256,
		HashTypeSha512,
		HashTypeXxhash,
		HashTypeTiger,
	}

	EInvalidHash = errors.New("checksums do not match")
)

func (f HashFlags) Contains(b HashFlags) bool {
	return f&b == b
}

func (t HashType) Flag() HashFlags {
	switch t {
	case HashTypeMd5:
		return HASH_TYPE_MD5
	case HashTypeSha1:
		return HASH_TYPE_SHA1
	case HashTypeSha256:
		return HASH_TYPE_SHA256
	case HashTypeSha512:
		return HASH_TYPE_SHA512
	case HashTypeXxhash:
		return HASH_TYPE_XXHASH
	case HashTypeTiger:
		return HASH_TYPE_TIGER
	default:
		return HASH_TYPE_INVALID
	}
}

func (l HashTypeList) String() string {
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = string(v)
	}
	return strings.Join(s, ",")
}

func (l HashTypeList) Flags() HashFlags {
	var f HashFlags
	for _, v := range l {
		f |= v.Flag()
	}
	return f
}

func ParseTypeList(s string) HashTypeList {
	l := make(HashTypeList, 0)
	for _, v := range strings.Split(s, ",") {
		t := HashType(v)
		if t.Flag() > 0 {
			l = append(l, t)
		}
	}
	return l
}

// Text/JSON conversion
func (r HashTypeList) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *HashTypeList) UnmarshalText(data []byte) error {
	rr := ParseTypeList(string(data))
	*r = rr
	return nil
}

// SQL conversion
func (r *HashTypeList) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		*r = ParseTypeList(v)
	case []byte:
		*r = ParseTypeList(string(v))
	}
	return nil
}

func (r HashTypeList) Value() (driver.Value, error) {
	return r.String(), nil
}

func (h *HashBlock) IsZero() bool {
	return h.Md5 == "" &&
		h.Sha1 == "" &&
		h.Sha256 == "" &&
		h.Sha512 == "" &&
		h.XXHash == "" &&
		h.Tiger == ""
}

func (h *HashBlock) Contains(l ...HashType) bool {
	hf := h.Flags()
	ll := HashTypeList(l)
	lf := ll.Flags()
	return hf&lf == lf
}

func (h *HashBlock) Flags() HashFlags {
	var a HashFlags
	if h.Md5 != "" {
		a |= HASH_TYPE_MD5
	}
	if h.Sha1 != "" {
		a |= HASH_TYPE_SHA1
	}
	if h.Sha256 != "" {
		a |= HASH_TYPE_SHA256
	}
	if h.Sha512 != "" {
		a |= HASH_TYPE_SHA512
	}
	if h.XXHash != "" {
		a |= HASH_TYPE_XXHASH
	}
	if h.Tiger != "" {
		a |= HASH_TYPE_TIGER
	}
	return a
}

func (h *HashBlock) AnyFlag() HashFlags {
	switch {
	case h.Sha256 != "":
		return HASH_TYPE_SHA256
	case h.Sha512 != "":
		return HASH_TYPE_SHA512
	case h.XXHash != "":
		return HASH_TYPE_XXHASH
	case h.Tiger != "":
		return HASH_TYPE_TIGER
	case h.Sha1 != "":
		return HASH_TYPE_SHA1
	case h.Md5 != "":
		return HASH_TYPE_MD5
	}
	return 0
}

func (h HashBlock) Clone(flags HashFlags) HashBlock {
	var n HashBlock
	if flags&HASH_TYPE_MD5 > 0 {
		n.Md5 = h.Md5
	}
	if flags&HASH_TYPE_SHA1 > 0 {
		n.Sha1 = h.Sha1
	}
	if flags&HASH_TYPE_SHA256 > 0 {
		n.Sha256 = h.Sha256
	}
	if flags&HASH_TYPE_SHA512 > 0 {
		n.Sha512 = h.Sha512
	}
	if flags&HASH_TYPE_XXHASH > 0 {
		n.XXHash = h.XXHash
	}
	if flags&HASH_TYPE_TIGER > 0 {
		n.Tiger = h.Tiger
	}
	return n
}

func (h *HashBlock) NewReader(r io.Reader, flags HashFlags) io.Reader {

	// always enable default hash
	flags |= DefaultHash.Flag()

	rr := r
	if (flags&HASH_TYPE_MD5 > 0) && h.md5 == nil {
		h.md5 = md5.New()
		rr = io.TeeReader(rr, h.md5)
	}
	if (flags&HASH_TYPE_SHA1 > 0) && h.sha1 == nil {
		h.sha1 = sha1.New()
		rr = io.TeeReader(rr, h.sha1)
	}
	if (flags&HASH_TYPE_SHA256 > 0) && h.sha256 == nil {
		h.sha256 = sha256.New()
		rr = io.TeeReader(rr, h.sha256)
	}
	if (flags&HASH_TYPE_SHA512 > 0) && h.sha512 == nil {
		h.sha512 = sha512.New()
		rr = io.TeeReader(rr, h.sha512)
	}
	if (flags&HASH_TYPE_XXHASH > 0) && h.xxhash == nil {
		h.xxhash = xxhash.New()
		rr = io.TeeReader(rr, h.xxhash)
	}
	if (flags&HASH_TYPE_TIGER > 0) && h.tiger == nil {
		h.tiger = tiger.NewTiger2()
		rr = io.TeeReader(rr, h.tiger)
	}
	return rr
}

func (h *HashBlock) NewWriter(w io.Writer, flags HashFlags) io.Writer {

	// always enable default hash
	flags |= DefaultHash.Flag()

	// each hash.Hash implements io.Writer
	wl := make([]io.Writer, 0, len(HashTypesAll)+1)
	wl = append(wl, w)

	if (flags&HASH_TYPE_MD5 > 0) && h.md5 == nil {
		h.md5 = md5.New()
		wl = append(wl, h.md5)
	}
	if (flags&HASH_TYPE_SHA1 > 0) && h.sha1 == nil {
		h.sha1 = sha1.New()
		wl = append(wl, h.sha1)
	}
	if (flags&HASH_TYPE_SHA256 > 0) && h.sha256 == nil {
		h.sha256 = sha256.New()
		wl = append(wl, h.sha256)
	}
	if (flags&HASH_TYPE_SHA512 > 0) && h.sha512 == nil {
		h.sha512 = sha512.New()
		wl = append(wl, h.sha512)
	}
	if (flags&HASH_TYPE_XXHASH > 0) && h.xxhash == nil {
		h.xxhash = xxhash.New()
		wl = append(wl, h.xxhash)
	}
	if (flags&HASH_TYPE_TIGER > 0) && h.tiger == nil {
		h.tiger = tiger.NewTiger2()
		wl = append(wl, h.tiger)
	}
	return io.MultiWriter(wl...)
}

func (h *HashBlock) Sum() {
	if h.md5 != nil {
		h.Md5 = hex.EncodeToString(h.md5.Sum(nil))
	}
	if h.sha1 != nil {
		h.Sha1 = hex.EncodeToString(h.sha1.Sum(nil))
	}
	if h.sha256 != nil {
		h.Sha256 = hex.EncodeToString(h.sha256.Sum(nil))
	}
	if h.sha512 != nil {
		h.Sha512 = hex.EncodeToString(h.sha512.Sum(nil))
	}
	if h.xxhash != nil {
		h.XXHash = hex.EncodeToString(h.xxhash.Sum(nil))
	}
	if h.tiger != nil {
		h.Tiger = hex.EncodeToString(h.tiger.Sum(nil))
	}
}

func (h HashBlock) Check(h2 HashBlock, ignoreempty bool) error {
	for _, k := range HashTypesAll {
		v1 := h.Get(k)
		v2 := h2.Get(k)
		if ignoreempty && (v1 == "" || v2 == "") {
			continue
		}
		if v1 != v2 {
			return EInvalidHash
		}
	}
	return nil
}

func (h *HashBlock) Clear() {
	h.Reset()
	h.Md5 = ""
	h.Sha1 = ""
	h.Sha256 = ""
	h.Sha512 = ""
	h.XXHash = ""
	h.Tiger = ""
}

func (h *HashBlock) Reset() {
	h.md5 = nil
	h.sha1 = nil
	h.sha256 = nil
	h.sha512 = nil
	h.xxhash = nil
	h.tiger = nil
}

func (h *HashBlock) Set(key HashType, val string) {
	switch key {
	case HashTypeMd5:
		h.Md5 = val
	case HashTypeSha1:
		h.Sha1 = val
	case HashTypeSha256:
		h.Sha256 = val
	case HashTypeSha512:
		h.Sha512 = val
	case HashTypeXxhash:
		h.XXHash = val
	case HashTypeTiger:
		h.Tiger = val
	}
}

func (h *HashBlock) Get(key HashType) string {
	switch key {
	case HashTypeMd5:
		return h.Md5
	case HashTypeSha1:
		return h.Sha1
	case HashTypeSha256:
		return h.Sha256
	case HashTypeSha512:
		return h.Sha512
	case HashTypeXxhash:
		return h.XXHash
	case HashTypeTiger:
		return h.Tiger
	default:
		return ""
	}
}

func (h HashBlock) JsonString() string {
	b, _ := json.MarshalIndent(h, "", "  ")
	return string(b)
}

func (h *HashBlock) Parse(r io.Reader) error {
	jsonDecoder := json.NewDecoder(r)
	if err := jsonDecoder.Decode(h); err != nil {
		return err
	}
	return nil
}

func (h HashBlock) String() string {
	s := make([]string, 0, len(HashTypesAll))
	for _, k := range HashTypesAll {
		v := h.Get(k)
		if v != "" {
			s = append(s, strings.Join([]string{string(k), v}, ":"))
		}
	}
	return strings.Join(s, ";")
}

func (h *HashBlock) ParseEtag(etag string) {
	if etag == "" {
		return
	}

	// trim spaces, optional weak prefix and quotes
	etag = strings.Trim(strings.TrimPrefix(strings.TrimSpace(etag), "W/"), "\"")

	// before decoding
	if b, err := base64.StdEncoding.DecodeString(etag); err == nil {
		h.Md5 = hex.EncodeToString(b)
	}
}

func (h *HashBlock) Etag() string {
	if h.Md5 == "" {
		return ""
	}
	// etag is base 64 encoded and wrapped in quotes
	if m5, err := hex.DecodeString(h.Md5); err == nil {
		return base64.StdEncoding.EncodeToString(m5)
	}
	return ""
}

func (h *HashBlock) EtagHeader() string {
	if h.Md5 == "" {
		return "\"\""
	}
	return strings.Join([]string{
		"\"",
		h.Etag(),
		"\"",
	}, "")
}

func ParseString(s string) HashBlock {
	var h HashBlock
	if s == "" {
		return h
	}
	for _, v := range strings.Split(s, ";") {
		fields := strings.Split(v, ":")
		if len(fields) != 2 {
			continue
		}
		h.Set(HashType(fields[0]), fields[1])
	}
	return h
}

func Parse(r io.Reader) (HashBlock, error) {
	var h HashBlock
	err := h.Parse(r)
	return h, err
}

func ParseEtag(s string) HashBlock {
	var h HashBlock
	h.ParseEtag(s)
	return h
}

// Text/JSON conversion
func (r HashBlock) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *HashBlock) UnmarshalText(data []byte) error {
	rr := ParseString(string(data))
	*r = rr
	return nil
}

// SQL conversion
func (r *HashBlock) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		*r = ParseString(v)
	case []byte:
		*r = ParseString(string(v))
	}
	return nil
}

func (r HashBlock) Value() (driver.Value, error) {
	return r.String(), nil
}
