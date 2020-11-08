//@author: hdsfade
//@date: 2020-11-07-09:37
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollar float32
type database map[string]dollar

func (d dollar) String() string { return fmt.Sprintf("$%.2f", d) }

var listTmpl = template.Must(template.New("list").Parse(`
<html>
<body>
<table>
<tr style="text-align:left">
	<th>item</th>
	<th>price</th>
</tr>
{{range $k,$v := .}}
<tr>
	<td>{{$k}}</td>
	<td>{{$v | printf "%s"}}</td>
</tr>
{{end}}
</table>
</body>
</html>`))

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func (db database) list(w http.ResponseWriter, r *http.Request) {
	listTmpl.Execute(w,db)
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %s\n", item)
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, r *http.Request) {
	var mutex sync.Mutex
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s is already exists.", item)
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %s\n", price)
		return
	}
	mutex.Lock()
	db[item] = dollar(p)
	mutex.Unlock()
	fmt.Fprintf(w, "create %s: %s\n", item, dollar(p))
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	var mutex sync.Mutex
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")
	if _, ok := db[item]; !ok {
		fmt.Fprintf(w, "%s isn't exists.\n", item)
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %s\n", price)
		return
	}
	mutex.Lock()
	db[item] = dollar(p)
	mutex.Unlock()
	fmt.Fprintf(w, "update %s: %s\n", item, dollar(p))
}


func (db database) delete(w http.ResponseWriter, r *http.Request) {
	var mutex sync.Mutex
	item := r.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		mutex.Lock()
		delete(db, item)
		mutex.Unlock()
		fmt.Fprintf(w, "delete %s", item)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "cann't find %s", item)
	}
}

