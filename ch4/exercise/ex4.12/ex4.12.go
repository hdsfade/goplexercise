//@author: hdsfade
//@date: 2020-11-03-08:38
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

const (
	APIURL   = "https://xkcd.com"
	suffix   = "info.0.json"
	filename = "comics.json"
	usage    = ` id num
words query`
)

type Comic struct {
	Num              int
	Year, Month, Day string
	Title            string
	Transcript       string
	Alt              string
	Img              string //url
}

type WordIndex map[string]map[int]bool
type NumIndex map[int]Comic

//使用提示信息
func tips() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

//判断文件是否存在
func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func init() {
	//如果索引文件不存在则初始化
	if !isExist(filename) {
		err := index(filename)
		if err != nil {
			log.Fatal("failed to initial indexes\n", err)
		}
	}
}

//通过id获取comic
func get(num int) (Comic, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Comic{}, err
	}
	defer file.Close()

	var wordIndex WordIndex
	var numIndex NumIndex
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&wordIndex); err != nil {
		return Comic{}, err
	}

	if err := decoder.Decode(&numIndex); err != nil {
		return Comic{}, err
	}
	return numIndex[num], nil
}

func main() {
	if len(os.Args) < 2 {
		tips()
	}
	cmd := os.Args[1]
	switch cmd {
	case "id":
		if len(os.Args) != 3 {
			tips()
		}
		num, err := strconv.Atoi(os.Args[2])
		if err != nil {
			tips()
		}
		comic, err := get(num)
		if err != nil {
			log.Fatalf("can't get %d: %v", num, err)
		}
		fmt.Println(comic)
	case "words":
		if len(os.Args) != 3 {
			tips()
		}
		query := os.Args[2]
		err := search(query, filename)
		if err != nil {
			log.Fatal("Error searching index", err)
		}
	default:
		tips()
	}
}

func getComicCount() (int, error) {
	url := strings.Join([]string{APIURL, suffix}, "/")
	//返回的是最后一个comic的信息
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("can't get %s: %s", url, resp.Status)
	}

	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return 0, err
	}
	return comic.Num, err
}

func getComic(num int) (Comic, error) {
	url := fmt.Sprintf(APIURL+"/%d/"+suffix, num)
	resp, err := http.Get(url)
	if err != nil {
		return Comic{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Comic{}, fmt.Errorf("can't get %s: %s", url, resp.Status)
	}

	var comic Comic
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return Comic{}, err
	}
	return comic, nil
}

func getComics() ([]Comic, error) {
	comicNum, err := getComicCount()
	if err != nil {
		return nil, err
	}

	waitgroup := sync.WaitGroup{}
	waitgroup.Add(comicNum)
	comics := make([]Comic, 0)

	//将漫画划分成5份，并发处理
	partsNum := 5
	part := comicNum / partsNum
	for i := 0; i < partsNum; i++ {
		var st, ed int
		if i != partsNum-1 {
			st, ed = i*part, (i+1)*part
		} else {
			st, ed = i*part, comicNum
		}
		go func(st, ed int) {
			for j := st + 1; j <= ed; j++ {
				mutex := sync.Mutex{}
				comic, err := getComic(j)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					waitgroup.Done()
					continue
				}

				mutex.Lock()
				comics = append(comics, comic)
				mutex.Unlock()
				waitgroup.Done()
			}
		}(st, ed)
	}
	waitgroup.Wait()
	return comics, nil
}

//scanner划分函数
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	i, st, ed := 0, 0, 0

	//找到起始unicode字符位置
	for i < len(data) {
		r, nbytes := utf8.DecodeRune(data[i:])
		i += nbytes
		if unicode.IsLetter(r) {
			st = i - nbytes
			break
		}
	}

	//找到单词unicode字符结束位置
	for i < len(data) {
		r, nbytes := utf8.DecodeRune(data[i:])
		i += nbytes
		if !unicode.IsLetter(r) {
			ed = i - nbytes
			break
		}
	}
	if st < ed {
		token = data[st:ed]
	}
	return i, token, nil
}

func indexComics(comics []Comic) (WordIndex, NumIndex) {
	wordIndex := make(WordIndex)
	numIndex := make(NumIndex)

	for _, comic := range comics {
		numIndex[comic.Num] = comic
		scanner := bufio.NewScanner(strings.NewReader(comic.Transcript)) //将字符串转为io.Reader
		scanner.Split(ScanWords)
		for scanner.Scan() {
			token := strings.ToLower(scanner.Text())
			if _, ok := wordIndex[token]; !ok {
				wordIndex[token] = make(map[int]bool, 1)
			}
			wordIndex[token][comic.Num] = true
		}
	}
	return wordIndex, numIndex
}

//查找完全匹配关键词的comic
func comicsContainingWords(words []string, wordIndex WordIndex, numIndex NumIndex) []Comic {
	found := make(map[int]int) //comic num -> count words found
	comics := make([]Comic, 0)
	//统计匹配的关键词数目
	for _, word := range words {
		for num := range wordIndex[word] {
			found[num]++
		}
	}
	for num, nfound := range found {
		if nfound == len(words) {
			comics = append(comics, numIndex[num])
		}
	}
	return comics
}

func search(query string, filename string) error {
	wordIndex, numIndex, err := readIndex(filename)
	if err != nil {
		return err
	}
	comics := comicsContainingWords(strings.Fields(query), wordIndex, numIndex) //strings.Fields()完成对查询条件的分词
	for _, comic := range comics {
		fmt.Printf("%#v\n\n", comic)
	}
	return nil
}

//将索引以文件形式保存
func index(filename string) error {
	comics, err := getComics()
	if err != nil {
		return err
	}
	wordIndex, numIndex := indexComics(comics)
	file, err := os.Create(filename)
	if err != nil {
		return nil
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(wordIndex)
	if err != nil {
		return nil
	}
	err = encoder.Encode(numIndex)
	if err != nil {
		return nil
	}
	return nil
}

//从文件中读取索引
func readIndex(filename string) (WordIndex, NumIndex, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	decoder := json.NewDecoder(file)
	var wordIndex WordIndex
	var numIndex NumIndex
	err = decoder.Decode(&wordIndex)
	if err != nil {
		return nil, nil, err
	}
	err = decoder.Decode(&numIndex)
	if err != nil {
		return nil, nil, err
	}
	return wordIndex, numIndex, err
}
