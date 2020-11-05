//@author: hdsfade
//@date: 2020-11-05-21:19
package intset

type Intset struct {
	words []int
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
