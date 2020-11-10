//@author: hdsfade
//@date: 2020-11-09-15:44
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress message")
var sema = make(chan struct{}, 20)

func main() {
	flag.Parse()
	roots := flag.Args()
	fileSizes := make(chan int64)
	wg := &sync.WaitGroup{}
	if len(roots) == 0 {
		roots = []string{"."}
	}
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64

	go func() {
		wg.Wait()
		close(fileSizes)
	}()
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, fileSizes, wg)
	}
loop:
	for {
		select {
		case filesize, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += filesize
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			go walkDir(subdir, fileSizes, wg)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3: %s\n", err)
		return nil
	}
	return entries
}
