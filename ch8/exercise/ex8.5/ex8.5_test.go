//@author: hdsfade
//@date: 2020-11-08-19:46
package ex8_5

import (
	"testing"
)

func BenchmarkSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sequence()
	}
}

func BenchmarkSimultanous(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simultaneous()
	}
}
