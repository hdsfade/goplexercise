//@author: hdsfade
//@date: 2020-11-04-15:05
package expand

import "strings"

func expand(s string, f func(string) string) string {
	replace := f("foo")
	return strings.Replace(s, "$foo", replace, -1)
}
