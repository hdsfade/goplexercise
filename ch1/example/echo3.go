//@author: hdsfade
//@date: 2020-10-29-22:05
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
