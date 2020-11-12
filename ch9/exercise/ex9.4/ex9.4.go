//@author: hdsfade
//@date: 2020-11-11-15:52
package ex9_4

func pipeline(stages int) (in, out chan int) {
	out = make(chan int)
	first := out //first为第一个通道

	//上一个goroutine的out为下一个的in
	for i := 0; i < stages; i++ {
		in := out
		out := make(chan int)
		go func(in, out chan int) {
			for v := range in {
				out <- v
			}
			close(out)
		}(in, out)
	}

	//返回第一个和最后一个通道
	return first, out
}
