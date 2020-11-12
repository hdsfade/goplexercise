//@author: hdsfade
//@date: 2020-11-11-15:56
package ex9_4

import "testing"

func Benchmarkpipeline(b *testing.B) {
	//创建100000个goroutine
	in, out := pipeline(1000000)
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}
