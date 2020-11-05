//@author: hdsfade
//@date: 2020-11-04-14:04
//unfinished, it will be finished after reading chapter 11
package outline

import (
	"strings"
	"testing"
	"golang.org/x/net/html"
)

func TestforEachNode(t *testing.T) {
	input := `
<html>
!---test
<body>
	<h1>Hello!</h1><br/>
	<a href="https:www.baidu.com" color="red"><b>www.baidu.com</b></a>
</body>
<html>
`
	doc,err := html.Parse(strings.NewReader(input))
	if err != nil{
		t.Errorf(err)
	}
	forEachNode(doc,pre,post)
}
