//@author: hdsfade
//@date: 2020-11-07-08:34
package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollar float32

func (d dollar) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollar

func main() {
	db := database{"shoes":50,"socks":5}
	mux := http.NewServeMux()
	mux.Handle("/list",http.HandlerFunc(db.list))
	mux.Handle("/price",http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe(":8000",mux))
}

func (db database) list(w http.ResponseWriter,r *http.Request) {
	for item, price := range db{
		fmt.Fprint(w,"%s: %s\n",item,price)
	}
}

func (db database) price(w http.ResponseWriter,r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no sucn item: %s\n",item)
		return
	}
	fmt.Fprintf(w,"%s\n",price)
}