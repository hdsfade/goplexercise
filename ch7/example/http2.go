//@author: hdsfade
//@date: 2020-11-07-08:13
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe(":8000", db))
}

type dollar float32

func (d dollar) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollar

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range db{
			fmt.Fprintf(w,"%s: %s\n",item,price)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok{
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n",item)
			return
		}
		fmt.Fprintf(w,"%s\n",price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"no such page: %s\n",r.URL)
	}
}
