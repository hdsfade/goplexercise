//@author: hdsfade
//@date: 2020-11-04-08:43
package main

import (
	"fmt"
	"net/http"
	"os"
	"golang.org/x/net/html"
)

func main() {
	for _,url := range os.Args[1:] {
		links,err := findlinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v",url,err)
			continue
		}
		for _,link := range links{
			fmt.Println(link)
		}
	}
}

func findlinks(url string) ([]string,error) {
	resp, err := http.Get(url)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can't get %s: %s",url,resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil{
		return nil, err
	}
	return visit(nil,doc),nil
}

func visit(links []string, n *html.Node) [] string{
	if n.Type == html.ElementNode && n.Data == "a"{
		for _,a := range n.Attr{
			if a.Key == "href"{
				links = append(links,a.Val)
			}
		}
	}
	for c:=n.FirstChild;c!=nil;c = c.NextSibling{
		links = visit(links,c)
	}
	return links
}