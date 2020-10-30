//@author: hdsfade
//@date: 2020-10-29-22:39
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const testNum = 1000

func main() {
	var s1, s2, sep string

	start := time.Now().Nanosecond()
	for i := 0; i < testNum; i++ {
		s1 = ""
		for _, arg := range os.Args[1:] {
			s1 = s1 + sep + arg
			sep = " "
		}
	}
	fmt.Println(s1)
	end := time.Now().Nanosecond()
	fmt.Printf("method 1: %d nanosecond\n", end-start)

	start = time.Now().Nanosecond()
	for i := 0; i < testNum; i++ {
		s2 = strings.Join(os.Args[1:], " ")
	}
	fmt.Println(s2)
	end = time.Now().Nanosecond()
	fmt.Printf("method 2: %d nanosecond\n", end-start)
}
