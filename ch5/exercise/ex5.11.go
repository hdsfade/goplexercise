//@author: hdsfade
//@date: 2020-11-04-19:25
package main

import "fmt"

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":{"linear algebra"},
	"linear algebra":{"calculus"},

	"compilers":{
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math":{"intro to programming"},
	"formal languages":{"discrete math"},
	"networds": {"operating systems"},
	"operating systems":{"data structures", "computer organization"},
	"programming languages":{"data structures","computer organization"},
}

func toposort(m map[string][]string) ([]string,bool){
	var order []string
	flag := true
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item]=true
				visitAll(m[item])
				order = append(order,item)
			} else {
				flag = false
				break
			}
		}
	}
	for k,_ := range m {
		visitAll([]string{k})
	}
	return order,flag
}

func main() {
	toposequence, flag := toposort(prereqs)
	fmt.Println(flag)
	if flag {
		for i, course := range toposequence {
			fmt.Printf("%d:\t%s\n", i+1, course)
		}
	}else {
		fmt.Printf("The graph has loop.\n")
	}

}