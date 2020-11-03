//@author: hdsfade
//@date: 2020-11-01-16:00
package popcount

var pc [256]int

func init() {
	for i := 0; i < 256; i++ {
		pc[i] = pc[i/2] + i&1
	}
}

func popcount(sha [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		count += pc[sha[i]]
	}
}
