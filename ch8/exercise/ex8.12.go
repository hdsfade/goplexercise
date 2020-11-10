//@author: hdsfade
//@date: 2020-11-10-15:49
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
	users    = make(map[string]bool)
	usersM   = sync.Mutex{}
)

func broadcast() {
	clients := make(map[client]bool)
	for {
		select {
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		}
	}
}

func handleConn(conn net.Conn) {
	fmt.Fprintln(conn, "The chat room has:")
	for user := range users {
		fmt.Fprintln(conn, user)
	}

	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()

	usersM.Lock()
	users[who] = true
	usersM.Unlock()

	ch <- "You are" + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	delete(users, who)
	leaving <- ch
	messages <- who + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
