//@author: hdsfade
//@date: 2020-11-10-17:24
package bank

import "sync"

var (
	mu sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Banlance() int{
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
