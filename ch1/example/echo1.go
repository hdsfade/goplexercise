//@author: hdsfade
//@date: 2020-10-29-21:56
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s = s + sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
