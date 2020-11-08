//@author: hdsfade
//@date: 2020-11-07-07:56
package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	s1 := []string{"ab", "cd", "cd", "ab"}
	s2 := []string{"ab", "dd"}
	fmt.Println(IsPalindrome(sort.StringSlice(s1)))
	fmt.Println(IsPalindrome(sort.StringSlice(s2)))
}
