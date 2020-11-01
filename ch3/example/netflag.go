//@author: hdsfade
//@date: 2020-11-01-10:03
package main

import "fmt"

type Flags uint

const (
	FlagUP Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool {return v&FlagUP == FlagUP}
func TurnDown(v *Flags) {*v &^=FlagUP}  //&^位清除
func SetBroadcast(v *Flags) {*v |= FlagBroadcast}
func IsCast(v Flags) bool {return v&(FlagBroadcast|FlagMulticast) != 1}

func main() {
	var v Flags = FlagMulticast|FlagUP
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
}

