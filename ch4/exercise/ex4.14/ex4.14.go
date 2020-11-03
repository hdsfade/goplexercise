//@author: hdsfade
//@date: 2020-11-03-17:32
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func NewIssueCache(owner, repo string) (ic IssueCache, err error) {
	ic.Issues = make(map[int]Issue)
	issues, err := GetIssues(owner, repo)
	if err != nil{
		return
	}
	for _, issue := range issues{
		ic.Issues[issue.Number] = issue
	}
	return
}

func (ic IssueCache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	num,err := strconv.Atoi(strings.Split(r.URL.Path,"/")[1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_,err := w.Write([]byte(fmt.Sprintf("please input a number")))
		if err != nil{
			log.Printf("Error writing response for %s: %s", r, err)
		}
		return
	}
	issue, ok := ic.Issues[num]
	if !ok{
		w.WriteHeader(http.StatusNotFound)
		_,err := w.Write([]byte(fmt.Sprintf("No issue %d",num)))
		if err != nil{
			log.Printf("Error writing response for %s: %s", r,err)
		}
		return
	}
	issueTemplate.Execute(w, issue)
}

func main() {
	if len(os.Args) != 3{
		fmt.Fprintln(os.Stderr,"usage:owner repo")
		os.Exit(1)
	}
	owner := os.Args[1]
	repo := os.Args[2]

	issueCache, err := NewIssueCache(owner, repo)
	if err != nil{
		log.Fatal(err)
	}

	http.Handle("/", issueCache)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
