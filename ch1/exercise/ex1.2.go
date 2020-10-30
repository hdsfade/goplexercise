//@author: hdsfade
//@date: 2020-10-29-22:37
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("%d\t%s\n", i, arg)
	}
}
