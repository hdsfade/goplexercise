//@author: hdsfade
//@date: 2020-11-06-15:46
package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
)

var html = template.Must(template.New("music").Parse(`
<html>
<body>

<table>
	<tr style="text-align:left">
		<th><a href="?sort=title">title</a></th>
		<th><a href="?sort=artist">artist</a></th>
		<th><a href="?sort=album">album</a></th>
		<th><a href="?sort=year">year</a></th>
		<th><a href="?sort=length">length</a></th>
	</tr>
{{range .}}
	<tr>
		<td>{{.Title}}</td>
		<td>{{.Artist}}</td>
		<td>{{.Album}}</td>
		<td>{{.Year}}</td>
		<td>{{.Length}}</td>
	</tr>
{{end}}
</body>
</html>
`))


func main() {
	c := customSort{
		T:          tracks,
		less:       nil,
		maxColumns: 3,
	}
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		switch r:=r.FormValue("sort"); r{
		case "title":
			c.Addsort(byTitle)
		case "artist":
			c.Addsort(byArtist)
		case "length":
			c.Addsort(byLength)
		case "album":
			c.Addsort(byAlbum)
		case "year":
			c.Addsort(byYear)
		}

		sort.Sort(c)
		err := html.Execute(w, c.T)
		if err != nil{
			log.Printf("template err: %v",err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080",nil))
}