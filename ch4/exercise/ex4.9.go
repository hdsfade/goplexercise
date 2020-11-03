//@author: hdsfade
//@date: 2020-11-02-10:18
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan(){
		count[input.Text()]++
	}
	for k, v := range count{
		fmt.Printf("%s\t%d\n", k, v)
	}
}
