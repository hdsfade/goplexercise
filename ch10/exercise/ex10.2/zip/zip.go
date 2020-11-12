//@author: hdsfade
//@date: 2020-11-11-19:53

//注册zip格式
package zip

import (
	"archive/zip"
	"fmt"
	arprint "github.com/torbiak/gopl/ex10.2"
	"io"
	"os"
)

type reader struct {
	zipReader *zip.Reader
	fileLeft  []*zip.File   //zip内文件
	r         io.ReadCloser //当前文件
	toWrite   string
}

//实现io.Reader接口
func (r *reader) Read(b []byte) (int, error) {
	if r.r == nil && len(r.fileLeft) == 0 {
		return 0, io.EOF
	}
	if r.r == nil {
		f := r.fileLeft[0]
		r.fileLeft = r.fileLeft[1:]
		var err error
		r.r, err = f.Open()
		if err != nil {
			return 0, fmt.Errorf("read zip: %s", err)
		}
		//f不是目录
		if f.Mode()&os.ModeDir == 0 {
			r.toWrite = f.Name + ":\n"
		}
		written := 0
		if len(r.toWrite) > 0 {
			n := copy(b, r.toWrite)
			b = b[n:]
			r.toWrite = r.toWrite[n:]
			written += n
		}
		//将文件内容读到b中
		n, err := r.r.Read(b)
		written += n
		if err != nil {
			r.r.Close()
			r.r = nil
			if err == io.EOF {
				return written, nil
			}
		}
		return written, nil
	}
}

func NewReader(f *os.File) (io.Reader, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("new zip reader: %s", err)
	}
	r, err := zip.NewReader(f, stat.Size())
	if err != nil {
		return nil, fmt.Errorf("new zip reader: %s", err)
	}
	return &reader{r, r.File, nil, ""}, nil
}

func init() {
	arprint.RegisterFormat("zip", "PK", 0, NewReader)
}
