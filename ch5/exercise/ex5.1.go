//@author: hdsfade
//@date: 2020-11-03-20:37
package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil{
		fmt.Fprintf(os.Stderr,"ex5.1: %v\n",err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) []string{
	var links []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _,a := range n.Attr{
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visit(n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(n.NextSibling)
	}
	return links
}
