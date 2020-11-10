//@author: hdsfade
//@date: 2020-11-09-09:26
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintf(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintf(c, "\t", strings.ToLower(shout))
	wg.Done()
}

func handleConn(c net.Conn) {
	wg := &sync.WaitGroup{}
	defer func() {
		wg.Wait()
		c.Close()
	}()
	lines := make(chan string)
	timeout := time.Second
	timer := time.NewTimer(2 * time.Second)
	for {
		select {
		case line := <-lines:
			timer.Reset(timeout)
			wg.Add(1)
			go echo(c, line, 1*time.Second, wg)
		case <-timer.C:
			return
		}
	}
}

func scan(r io.Reader, lines chan<- string) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		lines <- s.Text()
	}
	if s.Err() != nil {
		log.Print("san: ", s.Err())
	}
}
