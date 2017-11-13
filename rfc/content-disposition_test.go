// Trimmer SDK
//
// Copyright (c) 2016-2017 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package rfc

import (
	"strings"
	"testing"
)

type Testcase struct {
	Decoded string
	Encoded string
}

type CDTestcase struct {
	Encoded string
	Type    string
	Lang    string
	Charset string
	Tokens  []string
}

var (
	Iso8859_1Testcases = []Testcase{
		Testcase{"normal.txt", "normal.txt"},
		Testcase{`åæø£`, "\xe5\xe6\xf8\xa3"},
		Testcase{`鏄庝集`, "___"},
		Testcase{`日本語`, "___"},
	}

	Utf8Testcases = []Testcase{
		Testcase{"normal.txt", "normal.txt"},
		Testcase{`åæø£`, `%c3%a5%c3%a6%c3%b8%c2%a3`},
		Testcase{`鏄庝集`, `%e9%8f%84%e5%ba%9d%e9%9b%86`},
		Testcase{`日本語`, `%e6%97%a5%e6%9c%ac%e8%aa%9e`},
	}

	ContentDispositionTestcases = []CDTestcase{
		CDTestcase{
			Encoded: "form-data; name=\"sample\"; filename=\"sample\"",
			Type:    "form-data",
			Lang:    "",
			Charset: "",
			Tokens:  []string{"name", "sample", "filename", "sample"},
		},
		CDTestcase{
			Encoded: "attachement; filename=\"\xe5\xe6\xf8\xa3\"; filename*=utf-8''%c3%a5%c3%a6%c3%b8%c2%a3",
			Type:    "attachement",
			Lang:    "",
			Charset: "utf-8",
			Tokens:  []string{"filename", "åæø£"},
		},
		CDTestcase{
			Encoded: "attachement; filename=\"___\"; filename*=utf-8''%e6%97%a5%e6%9c%ac%e8%aa%9e",
			Type:    "attachement",
			Lang:    "",
			Charset: "utf-8",
			Tokens:  []string{"filename", "日本語"},
		},
		CDTestcase{
			Encoded: "attachement; filename=\"___\"; filename*=utf-8'jp'%e6%97%a5%e6%9c%ac%e8%aa%9e",
			Type:    "attachement",
			Lang:    "jp",
			Charset: "utf-8",
			Tokens:  []string{"filename", "日本語"},
		},
	}
)

func TestIso8859_1Encode(t *testing.T) {
	for i, v := range Iso8859_1Testcases {
		enc := toIso8859_1(v.Decoded)
		if v.Encoded != enc {
			t.Errorf("Case %d unexpected ISO-8850-1 encoding: %s -> %s (%x), expected %s (%x)", i, v.Decoded, enc, enc, v.Encoded, v.Encoded)
		}
	}
}

func TestIso8859_1Decode(t *testing.T) {
	for i, v := range Iso8859_1Testcases {
		// ignore test cases with known non-existing reverse mapping
		if v.Encoded[0] == '_' {
			continue
		}
		dec := fromIso8859_1(v.Encoded)
		if v.Decoded != dec {
			t.Errorf("Case %d unexpected ISO-8850-1 decoding: %s -> %s (%x), expected %s (%x)", i, v.Encoded, dec, dec, v.Decoded, v.Decoded)
		}
	}
}

func TestUtf8Encode(t *testing.T) {
	for i, v := range Utf8Testcases {
		enc := strings.ToLower(quoteUTF8(v.Decoded))
		if v.Encoded != enc {
			t.Errorf("Case %d unexpected UTF-8 quote encoding: %s -> %s, expected %s", i, v.Decoded, enc, v.Encoded)
		}
	}
}

func TestUtf8Decode(t *testing.T) {
	for i, v := range Utf8Testcases {
		dec := unquoteUTF8(v.Encoded)
		if v.Decoded != dec {
			t.Errorf("Case %d unexpected UTF-8 quote decoding: %s -> %s, expected %s", i, v.Encoded, dec, v.Decoded)
		}
	}
}

func TestContentDispositionEncode(t *testing.T) {
	for i, v := range ContentDispositionTestcases {
		cd := ContentDisposition{
			Type:    v.Type,
			Lang:    v.Lang,
			Charset: v.Charset,
			Tokens:  make(map[string]string),
		}
		for j, l := 0, len(v.Tokens)/2; j < l; j++ {
			cd.Tokens[v.Tokens[j*2]] = v.Tokens[j*2+1]
		}
		enc := cd.Encode()
		if v.Encoded != enc {
			t.Errorf("Case %d unexpected Content-Disposition encode:\n  got: %s\n  expected: %s", i, enc, v.Encoded)
		}
	}
}

func TestContentDispositionDecode(t *testing.T) {
	for i, v := range ContentDispositionTestcases {
		cd := ContentDisposition{
			Type:    v.Type,
			Lang:    v.Lang,
			Charset: v.Charset,
			Tokens:  make(map[string]string),
		}
		for j, l := 0, len(v.Tokens)/2; j < l; j++ {
			cd.Tokens[v.Tokens[j*2]] = v.Tokens[j*2+1]
		}
		testcd := ParseContentDisposition(v.Encoded)
		if testcd.Type != v.Type {
			t.Errorf("Case %d unexpected Content-Disposition type:\n  got: %s\n  expected: %s", i, testcd.Type, v.Type)
		}
		for j, l := 0, len(v.Tokens)/2; j < l; j++ {
			name := v.Tokens[j*2]
			value := v.Tokens[j*2+1]
			testvalue := testcd.Get(name)
			if testvalue != value {
				t.Errorf("Case %d unexpected Content-Disposition value for %s:\n  got: %s\n  expected: %s", i, name, testvalue, value)
			}
		}
	}
}
