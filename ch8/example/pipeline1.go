//@author: hdsfade
//@date: 2020-11-08-18:20
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squqres := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squqres <- x * x
		}
	}()

	for {
		fmt.Println(<-squqres)
	}
}
