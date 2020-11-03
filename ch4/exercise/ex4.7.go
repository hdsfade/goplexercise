//@author: hdsfade
//@date: 2020-11-01-21:15
package revrune

import (
	"unicode/utf8"
)

func rev(b []byte) {
	blen := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[blen-1-i] = b[blen-1-i], b[i]
	}
}

func revrune(s []byte) []byte {
	for i := 0; i < len(s); {
		_, size := utf8.DecodeRune(s[i:])
		rev(s[i : i+size])
		i += size
	}
	rev(s)
	return s
}
