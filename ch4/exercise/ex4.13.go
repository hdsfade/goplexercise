//@author: hdsfade
//@date: 2020-11-02-22:23
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	url2 "net/url"
	"os"
)

//需要注意的是此API需要APIKEY了
const APIURL = "https://omdbapi.com/?"
const usage = "please input movie name"

type Movie struct{
	Title string
	Year string
	Poster string    //Poster以url形式给出
}

func getMovie(title string) (*Movie, error) {
	var movie Movie
	url := fmt.Sprintf("%st=%s", APIURL,url2.QueryEscape(title))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can't get %s: %s\n", url, resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&movie);err != nil {
		return nil,err
	}
	return nil,nil
}

//将海报写到磁盘上
func (m Movie) writePoster() error{
	url := m.Poster
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("can't get %s: %s", url, resp.Status)
		os.Exit(1)
	}

	file, err := os.Create(m.Title)
	if err != nil{
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_,err = writer.ReadFrom(resp.Body)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil{
		return err
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Errorf(usage)
		os.Exit(1)
	}
	title := os.Args[1]
	movie, err := getMovie(title)
	if err != nil {
		log.Fatal(err)
	}
	err = movie.writePoster()
	if err != nil{
		log.Fatal(err)
	}
}