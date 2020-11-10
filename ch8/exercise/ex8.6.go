//@author: hdsfade
//@date: 2020-11-09-08:16
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

//确保并发数量在二十个以内
var tokens = make(chan struct{}, 20)

func crawl2(link URL) []URL {
	fmt.Println(link.url)
	tokens <- struct{}{}
	list, err := extract(link)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []URL)
	args := make([]URL, 0)
	for _, v := range os.Args[1:] {
		args = append(args, URL{v, 0})
	}
	var n int
	n++
	go func() {
		worklist <- args
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			//链接未被访问过或深度小于3，访问
			if !seen[link.url] && link.depth <= 3 {
				seen[link.url] = true
				n++
				go func(link URL) {
					worklist <- crawl2(link)
				}(link)
			}
		}
	}
}

type URL struct {
	url   string
	depth int
}

func extract(link URL) ([]URL, error) {
	resp, err := http.Get(link.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can't get %s: %s", link.url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	worklist := make([]URL, 0)
	visitnode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					l, err := resp.Request.URL.Parse(a.Val)
					if err != nil {
						continue
					}
					worklist = append(worklist, URL{l.String(), link.depth + 1})
				}
			}
		}
	}
	forEachNode(doc, visitnode, nil)
	return worklist, nil
}

type visit func(n *html.Node)

func forEachNode(n *html.Node, pre, post visit) {
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
