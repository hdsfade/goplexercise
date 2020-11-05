//@author: hdsfade
//@date: 2020-11-05-21:29
package intset

import (
	"bytes"
	"fmt"
)

//system bit
const d = 32 << (^uint(0) >> 63)

type Intset struct {
	words []uint
}

func (s *Intset) Has(x int) bool {
	word, bit := x/d, uint(x%d)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *Intset) Add(x int) {
	word, bit := x/d, uint(x%d)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *Intset) UnionWith(t *Intset) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *Intset) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < d; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", d*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
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
	word, bit := x/d, uint(x%d)
	if word < len(s.words) {
		s.words[word] = s.words[word] &^ (1 << bit)
	}
}

func (s *Intset) Clear() {
	s.words = make([]uint, 0)
}

func (s *Intset) Copy() *Intset {
	ret := Intset{}
	copy(ret.words, s.words)
	return &ret
}

func (s *Intset) AddAll(eles ...int) {
	for _, v := range eles {
		s.Add(v)
	}
}

func (s *Intset) Intersaction(t *Intset) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}
func (s *Intset) Symsub(t *Intset) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *Intset) Elems() []int {
	ret := make([]int, 0)
	for i, word := range s.words {
		for j := 0; j < d; j++ {
			if (word << j) != 0 {
				ret = append(ret, i*d+j)
			}
		}
	}
	return ret
}
