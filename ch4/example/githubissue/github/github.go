//@author: hdsfade
//@date: 2020-11-02-16:04
package github

import "time"

//github提供的api
const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"created_at"`
	Body string  //MarkDowm格式
}

type User struct{
	Login string
	HTMLURL string `json:"html_url"`
}