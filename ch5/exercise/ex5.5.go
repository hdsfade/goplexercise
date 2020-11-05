//@author: hdsfade
//@date: 2020-11-04-09:01
package countwi

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"unicode"
	"unicode/utf8"
	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil{
		return
	}
	doc,err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s",err)
		return
	}
	words,images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, image int) {
	for c := n.FirstChild; c!=nil;c = c.NextSibling{
		w, i := countWordsAndImages(c)
		words += w
		image += i
	}
	if  n.Type == html.ElementNode && n.Data == "img"{
		image++
	}  else if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(ScanWords)
		for scanner.Scan(){
			words++
		}
	}
	return
}

func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	for width := 0; start < len(data); start += width{
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if unicode.IsLetter(r){
			break
		}
	}
	for width, i := 0, start; i < len(data) ; i +=width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if !unicode.IsLetter(r) {
			return i + width, data[start:i],nil
		}
	}
	if atEOF && len(data) > start {
		return len(data), data[start:],nil
	}
	return start,nil,nil
}