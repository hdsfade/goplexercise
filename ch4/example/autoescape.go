//@author: hdsfade
//@date: 2020-11-03-15:13
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const tmpl = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(tmpl))
	var data struct {
		A string
		B template.HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil{
		log.Fatal(err)
	}
}
