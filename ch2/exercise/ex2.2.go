//@author: hdsfade
//@date: 2020-10-30-17:22
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Fahrenheit float64
type Celsius float64
type inch float64
type meter float64
type pound float64
type kilogram float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (i inch) String() string    { return fmt.Sprintf("%gin", i) }
func (m meter) String() string { return fmt.Sprintf("%gm", m) }
func (p pound) String() string    { return fmt.Sprintf("%glb", p) }
func (k kilogram) String() string { return fmt.Sprintf("%gkg", k) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f + 32) * 5 / 9) }
func IToM(i inch) meter { return meter(i / 39.37)}
func MToI(m meter) inch { return inch(m * 39.37) }
func PToK(p pound) kilogram { return kilogram(p*0.4536) }
func KToP(k kilogram) pound { return pound(k / 0.4536)}

func printresults(num float64) {
	c, f := Celsius(num), Fahrenheit(num)
	i, m := inch(num), meter(num)
	p, k := pound(num), kilogram(num)
	fmt.Printf("%s = %s, %s = %s\n", c, CToF(c), f, FToC(f))
	fmt.Printf("%s = %s, %s = %s\n", i, IToM(i), m, MToI(m))
	fmt.Printf("%s = %s, %s = %s\n", p, PToK(p), k, KToP(k))
}

func main() {
	if len(os.Args[1:]) < 1{
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		num, _ := strconv.ParseFloat(input.Text(),64)
		printresults( num)
	} else {
		for _, arg := range os.Args[1:] {
			num, _ := strconv.ParseFloat(arg,64)
			printresults(num)
		}
	}
}