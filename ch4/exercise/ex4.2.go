//@author: hdsfade
//@date: 2020-11-01-16:04
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var s = flag.String("s", "SHA256", "encoding by SHA256, SHA512 or SHA384")

func main() {
	flag.Parse() //先解析命令行参数
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex4.2: while reading %v\n", err)
		os.Exit(1)
	}
	switch *s {
	case "SHA512":
		fmt.Printf("%x\n", sha512.Sum512(input))
	case "SHA384":
		fmt.Printf("%x\n", sha512.Sum384(input))
	default:
		fmt.Printf("%x\n", sha256.Sum256(input))
	}
}
