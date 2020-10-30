//@author: hdsfade
//@date: 2020-10-30-19:39
//windows powershell: go test -bench="."
//to test performance
package main

import "testing"

var pc [256]byte

const testNum = 155

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountLoop(x uint64) int {
	result := 0
	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func BenchmarkPC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(testNum)
	}
}

func BenchmarkPCL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(testNum)
	}
}
