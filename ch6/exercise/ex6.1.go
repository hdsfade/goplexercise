//@author: hdsfade
//@date: 2020-11-05-20:50
package intset

type Intset struct{
	words []int64
}

func (s *Intset) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *Intset) Len() int {
	len := 0
	for _, word := range s.words {
		for word != 0 {
			len++
			word &= word - 1
		}
	}
	return len
}

func (s *Intset) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] = s.words[word] &^ (1 << bit)
	}
}

func (s *Intset) Clear() {
	s.words = make([]int64, 0)
}

func (s *Intset) Copy() *Intset {
	ret := Intset{}
	copy(ret.words, s.words)
	return &ret
}