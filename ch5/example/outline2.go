//@author: hdsfade
//@date: 2020-11-04-10:10
package outline

import (
	"fmt"
	"golang.org/x/net/html"
)

var depth int

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil{
		pre(n)
	}
	for c := n.FirstChild;c!=nil;c=c.NextSibling{
		forEachNode(c,pre,post)
	}
	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n",2*depth,"",n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n",2*depth,"",n.Data)
	}
}