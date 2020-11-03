//@author: hdsfade
//@date: 2020-11-01-21:00
package main

import (
	"fmt"
	"unicode"
)

func oneempty(s []rune) []rune {
	i := 0
	for _, v := range s[1:] {
		if unicode.IsSpace(v) && unicode.IsSpace(s[i-1]) {
			continue
		} else if unicode.IsSpace(v) {
			s[i] = ' '
			continue
		}
		s[i] = v
		i++
	}
	return s[:i]
}

func main() {
	s := []rune{' ', '是', ' ', ' ', ' ', '介'}
	fmt.Printf("%s", string(s))
	s = oneempty(s)
	fmt.Printf("%s", string(s))
}
