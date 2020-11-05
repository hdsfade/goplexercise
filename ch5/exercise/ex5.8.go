//@author: hdsfade
//@date: 2020-11-04-14:20
package elebyid

import (
	"golang.org/x/net/html"
)

type traverse func(*html.Node) bool

func ElementByID(doc *html.Node, id string) *html.Node {
	var result *html.Node
	//利用闭包特性
	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					result = n
					return false
				}
			}
		}
		return true
	}
	forEachNode(doc, pre, nil)
	return result
}

func forEachNode(n *html.Node, pre, post traverse){
	if pre != nil {
		if flag := pre(n); !flag {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		if flag := post(n); !flag {
			return
		}
	}
}
