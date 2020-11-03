//@author: hdsfade
//@date: 2020-11-01-20:50
package rmadjsame

func rmadjsame(strings []string) []string {
	i := 1
	for _, s := range strings[1:] {
		if s != strings[i-1] {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
