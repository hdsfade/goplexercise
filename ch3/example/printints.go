//@author: hdsfade
//@date: 2020-11-01-09:18
package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteString("]")
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
}
