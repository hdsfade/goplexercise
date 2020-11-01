//@author: hdsfade
//@date: 2020-11-01-09:39
package comma

import (
	"bytes"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	if s[0] == '-' || s[0] == '+' {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	ilen := strings.Index(s, ".") //整数部分长度
	i := ilen

	//确定第一个插入逗号的位置
	for i > 3 {
		i -= 3
	}
	buf.WriteString(s[:i])
	for ; i < ilen; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	buf.WriteString(s[ilen:])

	return buf.String()

}
