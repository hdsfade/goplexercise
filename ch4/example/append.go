//@author: hdsfade
//@date: 2020-11-01-17:25
package appendInt

func appendInt(x []int, y int) []int{
	var z []int
	zlen := len(x) + 1
	if zlen(x) <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[zlen] = y
	return z
}