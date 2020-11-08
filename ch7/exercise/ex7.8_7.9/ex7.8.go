//@author: hdsfade
//@date: 2020-11-06-16:08
package main

import (
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

//定义三种比较结构，只有相等的情况才需要根据上一次列进行比较
type comparison int

const (
	less  comparison = -1
	equal comparison = iota
	greater
)

var tracks = []*Track{{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Read 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}


type compare func(x,y *Track) comparison
type customSort struct {
	T          []*Track
	less       []compare
	maxColumns int
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil{
		panic(s)
	}
	return d
}

//注意此处接受者应为指针，否则无法保存函数
//如果x.less == nil，则加入函数相当于修改了地址
func (x *customSort) Addsort(cmp compare) {
	x.less = append([]compare{cmp}, x.less...)
	if len(x.less) > x.maxColumns{
		x.less = x.less[:x.maxColumns]
	}
}

func (x customSort) Len() int { return len(x.T) }
func (x customSort) Less(i, j int) bool {
	for _, f := range x.less {
		cmp := f(x.T[i], x.T[j])
		switch cmp {
		case less:
			return true
		case equal:
			continue
		case greater:
			return false
		}
	}
	return false
}
func (x customSort) Swap(i, j int) { x.T[i], x.T[j] = x.T[j], x.T[i] }

func byArtist(x, y *Track) comparison {
	switch {
	case x.Artist < y.Artist:
		return less
	case x.Artist == y.Artist:
		return equal
	default:
		return greater
	}
}
func byYear(x, y *Track) comparison {
	switch {
	case x.Year < y.Year:
		return less
	case x.Year == y.Year:
		return equal
	default:
		return greater
	}
}
func byTitle(x, y *Track) comparison {
	switch {
	case x.Title < y.Title:
		return less
	case x.Title == y.Title:
		return equal
	default:
		return greater
	}
}
func byAlbum(x, y *Track) comparison {
	switch {
	case x.Album < y.Album:
		return less
	case x.Album == y.Album:
		return equal
	default:
		return greater
	}
}
func byLength(x, y *Track) comparison {
	switch {
	case x.Length < y.Length:
		return less
	case x.Length == y.Length:
		return equal
	default:
		return greater
	}
}
