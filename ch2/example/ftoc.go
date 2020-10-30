//@author: hdsfade
//@date: 2020-10-30-16:03
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("freezing point = %g°F or %g°C\n", freezingF, ftoc(freezingF))
	fmt.Printf("boiling point = %g°F or %g°C\n", boilingF, ftoc(boilingF))
}

func ftoc(f float64) (c float64) {
	c = (f - 32) * 5 / 9
	return
}
