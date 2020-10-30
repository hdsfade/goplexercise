//@author: hdsfade
//@date: 2020-10-30-11:02
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const filepath = "data.txt"

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}

	file, err := os.OpenFile(filepath,os.O_CREATE, 0755)
	if err != nil {
		ch <- fmt.Sprintf("while opening file: %s: %v\n", url, err)
		return
	}
	file.Seek(0, 2)

	nbytes, err := io.Copy(file, resp.Body)
	file.Close()
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
