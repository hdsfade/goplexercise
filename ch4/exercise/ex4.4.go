//@author: hdsfade
//@date: 2020-11-01-20:25
package rotate

//n表示左移的位数
func rotate(s []int, n int) {
	tmp := make([]int, n)
	copy(tmp, s[:n])
	copy(s[:len(s)-n], s[n:])
	copy(s[len(s)-n:], tmp)
}
