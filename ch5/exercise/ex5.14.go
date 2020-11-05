//@author: hdsfade
//@date: 2020-11-05-08:53
package main

import "fmt"

var subway = map[string][]string{
	"宁波火车站": {"柳西新村", "兴宁桥西", "城隍庙"},
	"柳西新村":  {"宁波火车站", "大御桥"},
	"大御桥":   {"柳西新村", "西门口"},
	"西门口":   {"鼓楼", "大御桥"},
	"鼓楼":    {"西门口", "城隍庙", "天一广场"},
	"天一广场":  {"鼓楼", "江厦桥东"},
	"江厦桥东":  {"天一广场", "舟孟北路"},
	"舟孟北路":  {"江厦桥东", "樱花公园"},
	"樱花公园":  {"舟孟北路", "儿童乐园"},
	"儿童乐园":  {"白鹤公园", "樱花公园"},
	"白鹤公园":  {"儿童乐园", "兴宁桥东"},
	"兴宁桥东":  {"白鹤公园", "兴宁桥西"},
	"兴宁桥西": {"兴宁桥东","宁波火车站"},
}

func visitnode(n string) []string {
	return subway[n]
}

func breadthFirst(worklist []string, f func(n string) []string)[]string {
	var order []string
	seen := map[string]bool{}
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				order = append(order, item)
				worklist = append(worklist, f(item)...)
			}
		}
	}
	return order
}

func main() {
	//从某一结点出发历遍图
	for key := range subway {
		fmt.Println(breadthFirst([]string{key},visitnode))
		break
	}

}
