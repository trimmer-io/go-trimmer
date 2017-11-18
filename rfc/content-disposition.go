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

// Content-Disposition header helpers
//
// RFC2231 - https://tools.ietf.org/html/rfc2231
// RFC2183 - https://tools.ietf.org/html/rfc2183
// RFC6266 - https://tools.ietf.org/html/rfc6266
// RFC5987 - https://tools.ietf.org/html/rfc5987

package rfc

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type ContentDisposition struct {
	Type    string            // inline, attachement, form-data
	Lang    string            // ISO language
	Charset string            // utf-8, ..
	Tokens  map[string]string // name=xx, filename=xx, filename*=xx
}

func NewContentDisposition(t string) *ContentDisposition {
	return &ContentDisposition{
		Type:   t,
		Lang:   "",
		Tokens: make(map[string]string),
	}
}

func ParseContentDisposition(h string) *ContentDisposition {
	c := NewContentDisposition("")
	h = strings.NewReplacer("\r\n", "", "\n", "", "\r", "").Replace(h)
	for _, f := range strings.Split(h, ";") {
		f = strings.TrimSpace(f)
		tokens := strings.SplitN(f, "=", 2)
		switch len(tokens) {
		case 1:
			c.Type = strings.TrimSpace(strings.ToLower(f))
		case 2:
			token := strings.TrimSpace(strings.ToLower(tokens[0]))
			if strings.HasSuffix(token, "*") {
				token = strings.TrimRight(token, "*")
				ext := strings.SplitN(tokens[1], "'", 3)
				var lang, charset, content string
				// handle illegal extended field values
				switch len(ext) {
				case 1:
					content = ext[0]
				case 2:
					charset, content = ext[0], ext[1]
				case 3:
					charset, lang, content = ext[0], ext[1], ext[2]
				}
				c.Lang = strings.ToLower(lang)
				c.Charset = strings.ToLower(charset)
				switch c.Charset {
				case "utf8", "utf-8":
					c.Tokens[token] = unquoteUTF8(content)
				default:
					c.Tokens[token] = content
				}
			} else {
				c.Tokens[token] = fromIso8859_1(unquote(tokens[1]))
			}
		}
	}
	return c
}

func (c *ContentDisposition) Get(token string) string {
	if v, ok := c.Tokens[token]; ok {
		return v
	}
	return ""
}

func (c *ContentDisposition) Set(token, value string) {
	c.Tokens[token] = value
}

func (c *ContentDisposition) Encode() string {
	s := make([]string, 0, 1+len(c.Tokens))
	if c.Type != "" {
		s = append(s, c.Type)
	}

	for n, v := range c.Tokens {
		// check if token name is valid ascii
		name := strings.Map(func(r rune) rune {
			if r > unicode.MaxASCII {
				return '-'
			}
			return r
		}, n)

		// always encode fields for legacy software
		s = append(s, strings.Join([]string{name, quote(toIso8859_1(v))}, "="))

		// check if token value contains UTF-8 and also encode with * synatx
		switch strings.ToLower(c.Charset) {
		case "", "utf-8", "utf8":
			if c.Charset != "" && utf8.RuneCountInString(v) != len(v) {
				s = append(s, strings.Join([]string{
					name,
					"*=utf-8'",
					c.Lang,
					"'",
					quoteUTF8(v),
				}, ""))
			}
		}

	}
	return strings.Join(s, "; ")
}

// ISO-8859-1 contains UTF-8 runes between 0x00 and 0xFF identity mapped
func toIso8859_1(utf8String string) string {
	buf := make([]byte, utf8.RuneCountInString(utf8String))
	pos := 0
	for _, v := range utf8String {
		if int(v) <= 0xFF {
			buf[pos] = uint8(v)
		} else {
			buf[pos] = '_'
		}
		pos++
	}
	return string(buf)
}

func fromIso8859_1(iso8859_1_buf string) string {
	buf := make([]rune, len(iso8859_1_buf))
	for i, b := range []byte(iso8859_1_buf) {
		buf[i] = rune(b)
	}
	return string(buf)
}

// RFC 3986, 2.1 - https://tools.ietf.org/html/rfc3986#section-2.1
func percentEncode(b byte) []byte {
	t := make([]byte, 3)
	t[0] = '%'
	t[1] = "0123456789abcdef"[b>>4]
	t[2] = "0123456789abcdef"[b&15]
	return t
}

func quote(s string) string {
	return strings.Join([]string{"\"", s, "\""}, "")
}

func unquote(s string) string {
	return strings.Trim(s, "\"")
}

func decodeHex(b byte) byte {
	switch {
	case b >= '0' && b <= '9':
		return b - '0'
	case b >= 'a' && b <= 'f':
		return b - 'a' + 10
	case b >= 'A' && b <= 'F':
		return b - 'A' + 10
	default:
		return 0
	}
}

func percentDecode(b string) byte {
	if len(b) != 3 {
		return 0
	}
	return decodeHex(b[1])<<4 | decodeHex(b[2])
}

// RFC5987 https://tools.ietf.org/html/rfc5987
func quoteUTF8(s string) string {

	chars := make([]string, 0)

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		if width > 1 {
			// percent-hex-quote UTF-8 codepoints outside ASCII set
			buf := make([]byte, 3)
			var t []byte
			n := utf8.EncodeRune(buf, runeValue)
			for i := 0; i < n; i++ {
				t = append(t, percentEncode(buf[i])...)
			}
			chars = append(chars, string(t))
		} else {
			// percent-hex-quote reserved UTF-8 codepoints in ASCII set
			switch s[i] {
			// https://tools.ietf.org/html/rfc3986#section-2.2 + space (%20)
			case ' ', ':', '/', '?', '#', '[', ']', '@', '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=':
				chars = append(chars, string(percentEncode(s[i])))
			default:
				chars = append(chars, string(s[i]))
			}
		}
		w = width
	}

	return strings.Join(chars, "")
}

func unquoteUTF8(s string) string {

	chars := make([]string, 0)
	runeBuf := make([]byte, 0)
	for i, w := 0, 0; i < len(s); i += w {

		// percent-decode if we find a percent
		if s[i] == '%' {
			runeBuf = append(runeBuf, percentDecode(s[i:i+3]))
			w = 3

			// opportunistically try decoding a UTF-8 rune from transcoded bytes
			if r, size := utf8.DecodeRune(runeBuf); r != utf8.RuneError {
				chars = append(chars, string(r))
				runeBuf = runeBuf[size:]
			}

		} else {
			chars = append(chars, string(s[i]))
			w = 1
		}

	}
	return strings.Join(chars, "")
}
