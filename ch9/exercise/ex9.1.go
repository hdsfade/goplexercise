//@author: hdsfade
//@date: 2020-11-10-17:13
package bank

var (
	deposits  = make(chan int)
	balances  = make(chan int)
	withdraws = make(chan int)
	flag      = make(chan bool)
)

func withdraw(amount int) bool {
	withdraws <- amount
	return <-flag
}
func deposit(amount int) { deposits <- amount }
func banlance() int      { return <-balances }

func teller() {
	var balance int
	select {
	case amount := <-deposits:
		balance += amount
	case balances <- balance:
	case amount := <-withdraws:
		if amount > balance {
			flag <- false
		} else {
			balance -= amount
			flag <- true
		}
	}
}

func init() {
	go teller()
}
