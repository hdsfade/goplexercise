//@author: hdsfade
//@date: 2020-11-12-08:57
//题目翻译的不好
//此程序可以得出指定包的依赖
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var usage = "usage:ex10.4.exe [pkg]"

//保存依赖包名
type Dependence struct {
	Deps []string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	args := []string{"list", "-json", os.Args[1]}
	cmd := exec.Command("go", args...)
	var deps Dependence

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(strings.NewReader(string(output))).Decode(&deps)
	fmt.Println(deps)
}
