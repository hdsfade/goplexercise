//@author: hdsfade
//@date: 2020-10-30-14:02
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mutex sync.Mutex
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	count++
	mutex.Unlock()
	fmt.Fprintf(w, "URL.PATH: %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	fmt.Fprintf(w, "count: %d\n", count)
	mutex.Unlock()
}
