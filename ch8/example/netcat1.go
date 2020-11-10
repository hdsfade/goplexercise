//@author: hdsfade
//@date: 2020-11-08-15:02
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(w io.Writer, c net.Conn) {
	if _, err := io.Copy(w, c); err != nil {
		log.Fatal(err)
	}
}
