//@author: hdsfade
//@date: 2020-11-10-09:46
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var tokens = make(chan struct{}, 20)

func crawl2(url string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	fmt.Println(url)
	tokens <- struct{}{}
	list, err := extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	for _, link := range list {
		wg.Add(1)
		go crawl2(link, wg)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	seen := make(map[string]bool)
	for _, link := range os.Args[1:] {
		if !seen[link] {
			seen[link] = true
			wg.Add(1)
			go crawl2(link, wg)
		}
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt) //如果系统发送Interrupt，则发送信息给chan
	select {
	case <-done:
		return
	case <-interrupt:
		close(cancel)
	}
}

var cancel = make(chan struct{})

func extract(url string) ([]string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can't get %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	worklist := make([]string, 0)
	visitnode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := resp.Request.URL.Parse(a.Val)
					if err != nil {
						continue
					}
					worklist = append(worklist, link.String())
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
