//@author: hdsfade
//@date: 2020-11-04-17:30
package links

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
)

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("getting %s: %s",url,resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil{
		return nil,err
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _,a := n.Attr{
				if a.Key == "href"{
					link, err := resp.Request.URL.Parse(a.Val) //解析成基于resp.Request.URL的绝对路径
					if err != nil{
						continue
					}
					links = append(links,link.String())
				}
			}
		}
	}
	forEachNode(doc, visitNode,nil)
	return links, nil
}

func forEachNode(n *html.Node,pre,post func(*html.Node)) {
	if pre != nil{
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = n.NextSibling{
		forEachNode(c,pre,post)
	}
	if post != nil{
		post(n)
	}
}
