//@author: hdsfade
//@date: 2020-11-05-17:47
package main

import "fmt"

func a() (num int) {
	defer func() {
		if p := recover(); p != 0 {
			num = 1
		}
	}()
	panic("bad")
}

func main(){
	fmt.Println(a())
}