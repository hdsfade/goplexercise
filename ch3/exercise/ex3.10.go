//@author: hdsfade
//@date: 2020-11-01-09:26
package comma

import (
	"bytes"
)

func comma(s string) string {
	var buf bytes.Buffer
	i := len(s)
	//确定第一个要插入逗号的位置
	for i > 3 {
		i -= 3
	}

	buf.WriteString(s[:i])
	for ; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
