//@author: hdsfade
//@date: 2020-11-08-15:22
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

type clock struct {
	name, host string
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprint(os.Stderr, "usage: <name=host>")
		os.Exit(1)
	}
	clocks := make([]*clock, 0)
	for _, arg := range os.Args[1:] {
		fields := strings.Split(arg, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad arg: %s\n", arg)
			continue
		}
		clocks = append(clocks, &clock{fields[0], fields[1]})
	}
	for _, c := range clocks {
		conn, err := net.Dial("tcp", c.host)
		if err != nil {
			log.Print(err)
			continue
		}
		defer conn.Close()
		go c.watch(os.Stderr, conn)
	}
}

func (c clock) watch(w io.Writer, r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Fprintf(w, "%s: %s\n", c.name, s.Text())
	}
	fmt.Fprintf(w, "%s is done\n", c.name)
	if s.Err() != nil {
		log.Printf("can't read from %s", c.name)
	}

}
