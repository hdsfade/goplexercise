//@author: hdsfade
//@date: 2020-11-04-10:40
package ex5_7

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

var depth int

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func start(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		startElement(n)
	case html.TextNode:
		startText(n)
	case html.CommentNode:
		startComment(n)
	}
}

func end(n *html.Node) {
	if n.Type == html.ElementNode {
		endElement(n)
	}
}

func startElement(n *html.Node) {
	end := ">"
	if n.FirstChild == nil {
		end = "/>"
	}
	attrs := make([]string, len(n.Attr))
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf("%s=%s", a.Key, a.Val))
	}
	attrStr := ""
	if len(n.Attr) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}
	name := n.Data

	fmt.Printf("%*s<%s%s%s\n", depth*2, "", name, attrStr, end)
	depth++
}

func endElement(n *html.Node) {
	depth--
	if n.FirstChild == nil {
		return
	}
	fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
}

func startText(n *html.Node) {
	text := strings.TrimSpace(n.Data)
	if len(text) == 0 {
		return
	}
	fmt.Printf("%*s%s\n", depth*2, "", n.Data)
}
func startComment(n *html.Node) {
	fmt.Printf("<!--%s-->\n", n.Data)
}
