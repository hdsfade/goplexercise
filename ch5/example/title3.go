//@author: hdsfade
//@date: 2020-11-05-17:38
package main

import (
	"fmt"
	"go/build"
)

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil{
			if title != ""{
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	},nil)
	if title == ""{
		return "",fmt.Errorf("no title element")
	}
	return title, nil
}


