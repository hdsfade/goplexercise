//@author: hdsfade
//@date: 2020-10-30-19:33
package main

import (
	"fmt"
	"sync"
)

var pc [256]byte
var pcOnce sync.Once

func pcinit() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	pcOnce.Do(pcinit)
	result := 0
	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

func main() {
	fmt.Println(PopCount(1), PopCount(3))
}
