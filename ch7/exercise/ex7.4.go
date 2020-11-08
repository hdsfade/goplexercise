//@author: hdsfade
//@date: 2020-11-06-10:05
package exercise

import "io"

type stringReader struct {
	s string
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0{
		return n, io.EOF
	}
	return
}

func NewReader(s string) io.Reader{
	return &stringReader{s}
}