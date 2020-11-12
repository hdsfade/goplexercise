//@author: hdsfade
//@date: 2020-11-11-19:41
package arprint

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//压缩格式信息
type format struct {
	name, magic string
	magicOffset int
	reader NewReader
}

type NewReader func(*os.File) (error)

//全局压缩格式信息
var formats []format

//注册新的压缩格式
func RegisterFormat(name, magic string, magicOffset int,f NewReader) {
	formats = append(formats, format{name,magic,magicOffset,f})
}

//打开已注册格式的压缩文件
func Open(file *os.File) (error) {
	var found *format
	r := bufio.NewReader(file)

	//识别文件压缩格式
	for _, f := range formats {
		p, err := r.Peek(f.magicOffset+len(f.magic))
		if err != nil {
			continue
		}
		if string(p[f.magicOffset:]) == f.magic{
			found = &f
			break
		}
	}

	if found == nil{
		return fmt.Errorf("open archive: can't determine format")
	}

	//返回文件开头处
	_,err := file.Seek(0, os.SEEK_SET)
	if err != nil{
		return fmt.Errorf("open archive: %s\n", err)
		//返回对应压缩格式解读的io.Reader
		return found.reader(file)
	}
}