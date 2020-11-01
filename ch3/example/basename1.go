//@author: hdsfade
//@date: 2020-11-01-08:44
package basename

//basename移除路径部分和.后缀
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i-1]
			break
		}
	}
	return s
}
