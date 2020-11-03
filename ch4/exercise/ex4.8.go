//@author: hdsfade
//@date: 2020-11-02-10:03
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	count := make(map[string]int)
	count["letter"], count["number"] = 0, 0
	count["other"], count["valid"] = 0, 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount %v\n", err)
			os.Exit(1)
		}
		switch {
		case unicode.ReplacementChar == r && n == 1:
			count["invalid"]++
		case unicode.IsLetter(r):
			count["letter"]++
		case unicode.IsNumber(r):
			count["number"]++
		default:
			count["other"]++
		}
	}

	fmt.Printf("letter\tnumber\tother\tinvalid\n")
	fmt.Printf("%d\t%d\t%d\t%d\n", count["letter"], count["number"],
		count["other"], count["invalid"])

}
