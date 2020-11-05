//@author: hdsfade
//@date: 2020-11-04-17:53
package main

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/html"
	"os"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string)[]string{
	fmt.Println(url)
	links, err := Extract(url)
	if err != nil{
		log.Print(err)
	}
	return links
}

func Extract(url string) ([]string,error) {
	resp, err := http.Get(url)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil{
		return nil, err
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _,a := range n.Attr{
				if a.Key == "href"{
					link, err := resp.Request.URL.Parse(a.Val)
					if err != nil{
						continue
					}
					links = append(links, link.String())
				}
			}
		}
	}
	forEachNode(doc,visitNode,nil)
	return links,nil
}

func forEachNode(n *html.Node, pre,post func(n *html.Node)) {
	if pre != nil{
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling{
		forEachNode(c,pre,post)
	}
	if post != nil{
		post(n)
	}
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}