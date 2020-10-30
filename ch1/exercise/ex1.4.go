//@author: hdsfade
//@date: 2020-10-29-22:54
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filenames := make(map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		counts := make(map[string]int)
		countlines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			counts := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex1.4: %v", err)
				continue
			}
			countlines(f, counts)
			for _, count := range counts {
				if count > 1 {
					filenames[arg] = true
				}
			}
		}
	}

	for filename, _ := range filenames {
		fmt.Println(filename)
	}
}

func countlines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
