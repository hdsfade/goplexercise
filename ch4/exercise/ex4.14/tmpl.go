//@author: hdsfade
//@date: 2020-11-03-15:27
package main

import (
	"html/template"
)

var issueTmpl = `
<h1>{{.Title}}</h1>

<div>
<table>
<tr style='text-align: left'>
<th>id</th>
<th>description</th>
<th>url</th>
{{range .Bugs}}
<tr>
	<td> {{.Id}}</td>
	<td>{{.Description}}</td>
	<td>{{.HTMLURL}}</td>
</tr>
{{end}}
</table>
</div>

<div>
<table>
<tr style='text-align: left'>
<th>title</th>
<th>description</th>
<th>url</th>
</tr>
<tr>
<td>{{.Milestone.Title}}</td>
<td>{{.Milestone.Description}}</td>
<td>{{.Milestone.HTMLURL}}</td>
</tr>
</table>
</div>

<div>
<table>
<tr style='text-align: left'>
<th>login</th>
<th>url</th>
</tr>
<tr>
<td>{{.User.Login}}</td>
<td>{{.User.HTMLURL}}</td>
</tr>
</table>
</div>
`

var issueTemplate = template.Must(template.New("issue").Parse(issueTmpl))

type Issue struct {
	Number int
	Title string
	User *User
	Bugs []*Bug
	Milestone *Milestone
}

type Bug struct {
	Id int `json:"id"`
	Description string `json:"description"`
	HTMLURL string `json:"html_url"`
}

type Milestone struct {
	Title string `json:"title"`
	Description string `json:"description"`
	HTMLURL string `json:"html_url"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type IssueCache struct {
	Issues map[int]Issue
}