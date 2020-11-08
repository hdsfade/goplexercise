//@author: hdsfade
//@date: 2020-11-06-12:02
package main

import (
	"flag"
	"fmt"
	""
)

var temp = CelsiusFalg("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
