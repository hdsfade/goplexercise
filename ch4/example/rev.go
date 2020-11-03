//@author: hdsfade
//@date: 2020-11-01-16:46
package main

import "fmt"

func reverse1(s []int) {
	for i, j:=0,len(s)-1;i < j;i,j=i+1,j-1{
		s[i], s[j]=s[j],s[i]
	}
}

func reverse2(s []int) {
	for i:=0; i <= (len(s)-1)/2;i++{
		s[i],s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func main() {
	s1 := []int{1,2,3,4,5}
	s2 := []int{1,2,3,4}
	reverse1(s1)
	reverse2(s2)
	fmt.Println(s1,s2)
	reverse1(s2)
	reverse2(s1)
	fmt.Println(s1,s2)
}