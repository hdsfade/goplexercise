//@author: hdsfade
//@date: 2020-11-10-18:13
package memo

type entry struct {
	res   result
	ready chan struct{}
}
type request struct {
	key      string
	response chan<- result
	done     chan struct{}
}

type Memo struct {
	requests chan request
	cancel   chan request
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

//若done不为空，则取消操作，参考联系8.10
type Func func(key string, done chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func (memo *Memo) Get(key string, done chan struct{}) (interface{}, error) {
	response := make(chan result)
	req := request{key, response, done}
	memo.requests <- req
	res := <-response
	select {
	case <-done: //如果请求被取消，则done通道是关闭的
		memo.cancel <- req
	default:
	}
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for {
		select {
		case cancel := <-memo.cancel:
			delete(cache, cancel.key)
		case req := <-memo.requests:
			e := cache[req.key]
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, req.done)
			}
			go e.deliver(req.response)
		}
	}
}

func (e *entry) call(f Func, key string, done chan struct{}) {
	e.res.value, e.res.err = f(key, done)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
