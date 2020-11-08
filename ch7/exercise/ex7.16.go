//@author: hdsfade
//@date: 2020-11-07-18:47
package main

import (
	"fmt"
	"gopl.io/ch7/eval"
	"log"
	"net/http"
)

func computer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	s := r.Form.Get("expr")
	if s == "" {
		http.Error(w, "empty expression", http.StatusBadRequest)
		return
	}
	expr, err := eval.Parse(s)
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, expr.Eval(eval.Env{}))

}

func main() {
	http.HandleFunc("/", computer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
