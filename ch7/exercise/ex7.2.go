//@author: hdsfade
//@date: 2020-11-06-08:51
package exercise

import "io"

//纸质版的GO语言圣经对此题的翻译不正确，应该是：
//传入一个io.Writer，返回一个封装了传入io.Writer和int64类型指针的io.Writer
//int64类型指针对应的值表示新io.Writer写入的字节数
type WriterCounter struct {
	Writer io.Writer
	Count  *int64
}

func (w WriterCounter) Write(p []byte) (int, error) {
	n, err := w.Writer.Write(p)
	*w.Count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &WriterCounter{w, 0}
	return c, c.Count
}
