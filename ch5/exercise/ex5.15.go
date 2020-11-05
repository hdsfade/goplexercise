//@author: hdsfade
//@date: 2020-11-05-13:37
package main

import "fmt"

func max(num ...int) int {
	maxv := num[0]
	for _, v := range num[1:]{
		if v > maxv {
			maxv = v
		}
	}
	return maxv
}

func min(num ...int) int {
	minv := num[0]
	for _,v := range num[1:] {
		if v <minv {
			minv = v
		}
	}
	return minv
}

func main() {
	fmt.Println(max(1,2,3,4,5))
	fmt.Println(min(1,2,3,4,5))
}