//@author: hdsfade
//@date: 2020-11-03-20:51
package printtext

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

func printtext(r io.Reader, w io.Writer) error {
	z := html.NewTokenizer(os.Stdin)
	stack := make([]string, 20)
	for {
		switch z.Next() {
		case html.ErrorToken:
			if z.Err() != io.EOF {
				return z.Err()
			} else {
				return nil
			}
		//遇到开始标签，则压入栈，遇到结束标签，则将对应开始标签弹出栈
		//遇到文本，若其标签不是script和style，则输出对应标签和内容
		case html.StartTagToken:
			b, _ := z.TagName()
			stack = append(stack, string(b))
		case html.EndTagToken:
			stack = stack[:len(stack)-1]
		case html.TextToken:
			cur := stack[len(stack)-1]
			if cur != "script" && cur != "style" {
				text := z.Text()
				if len(strings.TrimSpace(string(text))) == 0{
					continue
				}
				w.Write([]byte(fmt.Sprintf("<%s>",cur)))
				w.Write(text)
				if text[len(text)-1] != '\n' {
					io.WriteString(w, "\n")
				}
			}
		}
	}
}

func main() {

}