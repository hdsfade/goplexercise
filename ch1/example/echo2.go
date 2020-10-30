//@author: hdsfade
//@date: 2020-10-29-22:02
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s = s + sep + arg
		sep = " "
	}
	fmt.Println(s)
}
