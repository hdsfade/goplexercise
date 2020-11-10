//@author: hdsfade
//@date: 2020-11-10-10:24
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var cancel = make(chan struct{})

func mirroredQuery() string {
	responses := make(chan string, 1)
	go func() { responses <- request("www.baidu.com") }()
	go func() { responses <- request("www.tencent.com") }()
	go func() { responses <- request("www.qq.com") }()
	return <-responses
}

func request(hostname string) (response string) {
	if !strings.HasPrefix(hostname, "http://") {
		hostname = "http://" + hostname
	}
	defer func() {
		close(cancel)
	}()

	req, err := http.NewRequest("GET", hostname, nil)
	req.Cancel = cancel //req.Cancel赋值一个通道，当通道关闭时，自动断开连接
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "can't get %s: %s", hostname, resp.Status)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	return string(buf)
}

func main() {
	fmt.Println(mirroredQuery())
}
