//@author: hdsfade
//@date: 2020-11-06-10:27
package exercise

import "io"

type limitReader struct {
	r        io.Reader
	n, limit int
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	n, err = l.r.Read(p[:l.limit-l.n])
	l.n += n
	if l.n >= l.limit {
		return n, io.EOF
	}
	return
}
func LimitReader(r io.Reader, n int) io.Reader {
	return &limitReader{r: r, limit: n}
}
