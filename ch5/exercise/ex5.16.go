//@author: hdsfade
//@date: 2020-11-05-13:42
package main

import "fmt"

func join(ele []string, sep ...string) string {
	result := ele[0]
	for i, v := range ele[1:]{
		if i < len(sep) {
			result = result + sep[i] + v
		} else {
			result = result + sep[len(sep)-1]+v  //如果sep不够，则开始固定使用最后一个sep
		}
	}
	return result
}

func main() {
	ele := []string{"go","programming","language","is","effective"}
	fmt.Println(join(ele," ","  ","\n"))
}