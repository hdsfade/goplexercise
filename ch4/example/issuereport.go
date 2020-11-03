//@author: hdsfade
//@date: 2020-11-03-14:42
package main

import (
	"gopl.io/ch4/github"
	"html/template"
	"log"
	"os"
	"time"
)

const tmpl = `{{.TotalCount}} issues:
{{range .Items}}----------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%6.4"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours()/24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo":daysAgo}).
	Parse(tmpl))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil{
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout,result); err != nil {
		log.Fatal(err)
	}
}