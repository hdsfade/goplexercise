//@author: hdsfade
//@date: 2020-11-06-11:31
package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type celsiusFlag struct {
	Celsius
}

func (c Celsius) String() string {
	return fmt.Sprintf("%f°C", c)
}

func FToC(f Fahrenheit) Celsius { return Celsius((f + 32) * 5 / 9) }

func (c *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		c.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		c.Celsius = Celsius(FToC(Fahrenheit(value)))
		return nil
	}
	return fmt.Errorf("invaild temperature %q", s)
}

func CelsiusFalg(name string, value Celsius, usage string) *Celsius {
	 f := celsiusFlag{value}
	 flag.CommandLine.Var(&f,name,usage)
	 return &f.Celsius
}

var temp = CelsiusFalg("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}