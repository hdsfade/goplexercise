//@author: hdsfade
//@date: 2020-11-09-09:03
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var token = make(chan struct{}, 20) //令牌，限制并发数为20
var maxDepth int                    //最大跳转
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}
var base *url.URL

//深度历遍
func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()

	tokens <- struct{}{}
	urls, err := visit(url)
	<-tokens
	if err != nil {
		log.Printf("visit %s: %v", url, err)
	}

	if depth >= maxDepth {
		return
	}
	for _, link := range urls {
		seenLock.Lock()
		if seen[link] {
			seenLock.Unlock()
			continue
		}
		seen[link] = true
		seenLock.Unlock()
		wg.Add(1)
		go crawl(link, depth, wg)
	}
}

//结点历遍
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

//返回根结点下所有的链接元素结点
func linkNodes(n *html.Node) []*html.Node {
	var links []*html.Node
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			links = append(links, n)
		}

	}
	forEachNode(n, visitNode, nil)
	return links
}

//处理链接元素结点，返回基于base_url解析的链接
func linkURLs(linkNodes []*html.Node, base *url.URL) []string {
	var urls []string
	for _, n := range linkNodes {
		for _, a := range n.Attr {
			if a.Key != "href" {
				continue
			}
			link, err := base.Parse(a.Val)
			if err != nil {
				log.Printf("parsing failed %q: %s", a.Val, err)
				continue
			}
			if link.Host != base.Host {
				log.Printf("skipping %q: non-local host", a.Val)
				continue
			}
			urls = append(urls, link.String())
		}
	}
	return urls
}

//重写链接为基于base_url的相对链接
func rewriteLocalLinks(linkNodes []*html.Node, base *url.URL) {
	for _, n := range linkNodes {
		for i, a := range n.Attr {
			if a.Key != "href" {
				continue
			}
			link, err := base.Parse(a.Val)
			if err != nil || link.Host != base.Host {
				continue
			}

			link.Scheme = ""
			link.Host = ""
			link.User = nil
			a.Val = link.String()
			n.Attr[i] = a
		}
	}
}

//处理链接
func visit(rawurl string) (urls []string, err error) {
	fmt.Println(rawurl)
	resp, err := http.Get(rawurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("GET %s: %s", rawurl, resp.Status)
	}

	u, err := base.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	if base.Host != u.Host {
		log.Printf("not saving %s: non-local", rawurl)
		return nil, nil
	}

	var body io.Reader
	contentType := resp.Header["Content-Type"]
	//html才需要解析链接
	if strings.Contains(strings.Join(contentType, ","), "text/html") {
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("parsing %s as HTML: %v", u, err)
		}
		nodes := linkNodes(doc)
		urls := linkURLs(nodes, u)
		rewriteLocalLinks(nodes, u)
		b := &bytes.Buffer{}
		err = html.Render(b, doc) //将以doc为根结点的html内容写入b
		if err != nil {
			log.Printf("render %s: %s", u, err)
		}
		body = b
	}
	err = save(resp, body)
	return urls, err
}

//保存网页内容
func save(resp *http.Response, body io.Reader) error {
	u := resp.Request.URL
	filename := filepath.Join(u.Host, u.Path)
	if filepath.Ext(u.Path) == "" {
		filename = filepath.Join(u.Host, u.Path, "index.html")
	}
	err := os.MkdirAll(filepath.Dir(filename), 0777)
	if err != nil {
		return err
	}
	fmt.Println("filename:", filename)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	if body != nil {
		_, err = io.Copy(file, body)
	} else {
		_, err = io.Copy(file, resp.Body)
	}
	if err != nil {
		log.Print("save: ", err)
	}
	err = file.Close()
	if err != nil {
		log.Print("save: ", err)
	}
	return nil
}

func main() {
	flag.IntVar(&maxDepth, "d", 3, "max crawl depth")
	flag.Parse()
	wg := &sync.WaitGroup{}
	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "usage: mirror URL ...")
		os.Exit(1)
	}
	u, err := url.Parse(os.Args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid url: %s\n", err)
	}
	base = u
	for _, link := range flag.Args() {
		wg.Add(1)
		go crawl(link, 1, wg)
	}
	wg.Wait()
}
