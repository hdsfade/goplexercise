//@author: hdsfade
//@date: 2020-11-01-20:13
package reverse

const arraynum = 5

func reverse(s *[arraynum]int) {
	for i := 0; i <= (len(s)-1)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
