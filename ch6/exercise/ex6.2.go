//@author: hdsfade
//@date: 2020-11-05-21:08
package intset

type Intset struct{
	words []int64
}

func (s *Intset) Add(x int) {
	word, bit := x/64,uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words,0)
	}
	s.words[word] |= 1 << bit
}

func (s *Intset) AddAll(eles ...int) {
	for _, v := range eles {
		s.Add(v)
	}
}
