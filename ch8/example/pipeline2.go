//@author: hdsfade
//@date: 2020-11-08-18:25
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x
		}
		close(squares)
	}()

	go func() {
		for x := range squares {
			fmt.Println(x)
		}
	}()
}
