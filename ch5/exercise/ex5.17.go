//@author: hdsfade
//@date: 2020-11-05-14:00
package main

import (
	"golang.org/x/net/html"
)

func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	var result []*html.Node
	start := func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, v := range name {
				if n.Data == v {
					result = append(result, n)
					break
				}
			}
		}
	}
	forEachNode(doc, name, start, nil)
	return result
}

func forEachNode(n *html.Node, name []string, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, name,pre, post)
	}
	if post != nil{
		post(n)
	}
}
