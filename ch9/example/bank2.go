//@author: hdsfade
//@date: 2020-11-10-17:22
package bank

var (
	sema = make(chan struct{})
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

func Balance() int{
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
