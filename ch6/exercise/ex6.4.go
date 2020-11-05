//@author: hdsfade
//@date: 2020-11-05-21:20
package intset

type Intset struct {
	words []int
}

func (s *Intset) Elems() []int {
	ret := make([]int, 0)
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if (word << j) != 0 {
				ret = append(ret, i*64+j)
			}
		}
	}
	return ret
}
