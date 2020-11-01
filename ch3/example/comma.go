//@author: hdsfade
//@date: 2020-11-01-08:56
package comma

import (
	"fmt"
	"strings"
)

func comma1(s string) string {
	for i := len(s) - 3; i > 0;i -= 3{
		s = s[:i] +","+s[i:]
	}
	return s
}

//减治法
func comma2(s string) string {
	n := len(s)
	if n <= 3{
		return s
	}
	return comma2(s[:n-3]) + "," + s[n-3:]
}
