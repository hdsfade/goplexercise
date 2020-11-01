//@author: hdsfade
//@date: 2020-11-01-09:49
package sameletter

func sameletter(s1 string, s2 string) bool {
	Lfreq1 := make(map[rune]int) //s1字符频次
	Lfreq2 := make(map[rune]int) //s2字符频次

	for _, v := range s1 {
		Lfreq1[v]++
	}
	for _, v := range s2 {
		Lfreq2[v]++
	}

	for k, v := range Lfreq1 {
		if Lfreq2[k] != v {
			return false
		}
	}
	for k, v := range Lfreq2 {
		if Lfreq1[k] != v {
			return false
		}
	}
	return true
}
