//@author: hdsfade
//@date: 2020-11-11-16:05
package main

import (
	"fmt"
	"time"
)

var ch1 = make(chan int)
var ch2 = make(chan int)
var count int

func person(in, out chan int) {
	for {
		<-in
		count++
		out <- 1
	}
}

func main() {
	//两个goroutine不可能同时修改count
	go person(ch1, ch2)
	go person(ch2, ch1)
	ch1 <- 1
	<-time.After(10 * time.Second)
	fmt.Println(count)
}
