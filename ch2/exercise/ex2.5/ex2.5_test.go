//@author: hdsfade
//@date: 2020-10-30-20:04
package main

import "testing"

const testNum = 325

func PopCount(x uint64) {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
}

func BenchmarkPC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(testNum)
	}
}
