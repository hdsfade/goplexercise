//@author: hdsfade
//@date: 2020-11-08-20:15
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

func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	/*err := save(url)
	if err != nil {
		log.Printf(`can't cache "%s": %s`, url, err)
	}*/
	list, err := extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

/*var originhost = ""
func save(rawurl string) error{
	url,err := url.Parse(rawurl)
	if err != nil{
		return err
	}
	if originhost == ""{
		originhost = url.Host
	}
	if url.Host != originhost{
		return nil
	}
	var filename string
	dir := url.Host
	if filepath.Ext(url.Path)=="" {
		dir = filepath.Join(dir, url.Path)
		filename = filepath.Join(dir,"index.html")
	}else {
		filename  = filepath.Join(dir, url.Path)
		dir = filepath.Join(dir,filepath.Dir(url.Path))
	}
	err = os.MkdirAll(dir,0777)
	if err != nil{
		return err
	}

	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file,err := os.Create(filename)
	if err !=nil{
		return err
	}
	_,err = io.Copy(file,resp.Body)
	if err != nil{
		return err
	}
	err = file.Close()
	if err != nil{
		return err
	}
	return nil
}*/

func main() {
	worklist := make(chan []string)
	var n int
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl2(link)
				}(link)
			}
		}
	}
}

func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
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
