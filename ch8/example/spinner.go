//@author: hdsfade
//@date: 2020-11-08-14:36
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100*time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n",n, fibN)
}

func spinner(d time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c",r)
			time.Sleep(d)
		}
	}
}

func fib(n int) int{
	if n < 2{
		return n
	}else {
		return fib(n-1)+fib(n-2)
	}
}