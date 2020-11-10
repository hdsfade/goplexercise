//@author: hdsfade
//@date: 2020-11-08-16:34
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	tcpaddr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c *net.TCPConn, shout string, delay time.Duration, waitgroup sync.WaitGroup) {
	defer waitgroup.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintf(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintf(c, "\t", strings.ToLower(shout))
}

func handleConn(c *net.TCPConn) {
	waitgroup := sync.WaitGroup{}
	input := bufio.NewScanner(c)
	for input.Scan() {
		waitgroup.Add(1)
		go echo(c, input.Text(), 1*time.Second, waitgroup)
	}
	waitgroup.Wait()
	c.CloseWrite()
}
