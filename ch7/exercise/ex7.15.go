//@author: hdsfade
//@date: 2020-11-08-08:29
package main

import (
	"bufio"
	"fmt"
	"gopl.io/ch7/eval"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Expression: ")
	input.Scan()
	exprStr := input.Text()
	fmt.Printf("Variables (<var>=<val>), eg: x=2): ")
	input.Scan()
	envStr := input.Text()
	if input.Err() != nil {
		fmt.Fprintln(os.Stderr, input.Err())
		os.Exit(1)
	}
	env := eval.Env{}
	assignments := strings.Fields(envStr)
	for _, a := range assignments {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad assignment: %s\n", a)
		}
		ident, valStr := fields[0], fields[1]
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bad value for %s, using zero: %s\n", ident, err)
		}
		env[eval.Var(ident)] = val
	}
	expr, err := eval.Parse(exprStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bad expression: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(expr.Eval(env))
	os.ErrNotExist
}
