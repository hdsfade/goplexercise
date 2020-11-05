//@author: hdsfade
//@date: 2020-11-03-20:47
package elecount

import "golang.org/x/net/html"

func elecount(n *html.Node, count *int) {
	if n.Type == html.ElementNode {
		*count++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elecount(c, count)
	}
}
