//@author: hdsfade
//@date: 2020-11-09-16:16
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const delay = 1 * time.Second //计算显示间隔时间

func main() {
	roots := os.Args
	if len(roots) == 0 {
		roots = []string{"."}
	}

	tick := time.Tick(1 * delay)
	stop := make(chan int)

	go func() {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
		}
		stop <- 1
	}()

loop:
	for {
		select {
		case <-tick:
			nfiles, nbytes := dirInfo(roots)
			printDiskUsage(nfiles, nbytes)
		case <-stop:
			break loop
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1fGB\n", nfiles, float64(nbytes)/1e9)
}

func dirInfo(roots []string) (nfiles, nbytes int64) {
	wg := &sync.WaitGroup{}
	fileSizes := make(chan int64)
	go func() {
		wg.Wait()
		close(fileSizes)
	}()
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, fileSizes, wg)
	}
	for filesize := range fileSizes {
		nfiles++
		nbytes += filesize
	}
	return
}

func walkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			walkDir(subdir, fileSizes, wg)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex8.9: %s\n", err)
		return nil
	}
	return entries
}
