//@author: hdsfade
//@date: 2020-11-10-18:13
package memo

import (
	"sync"
)

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]result
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo{
	return &Memo{f:f,cache:make(map[string]result)}
}

func(memo *Memo) Get(key string) (value interface{},err error) {
	memo.mu.Lock()
	res,ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok{
		res.value, res.err = memo.f(key)
		memo.cache[key]=res
	}
	return res.value,res.err
}
