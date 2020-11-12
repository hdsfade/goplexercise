//@author: hdsfade
//@date: 2020-11-10-16:39
package bank

var deposits = make(chan int)
var balances = make(chan int)

func deposit(amount int) { deposits <- amount }
func banlance() int      { return <-balances }

func teller() {
	var balance int
	select {
	case amount := <-deposits:
		balance += amount
	case balances <- balance:
	}
}

func init() {
	go teller()
}