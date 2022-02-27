package Day5

// 带缓冲通道的channel通道可以保留一定缓冲的数据直到数据被接收方接收，如果通道缓冲满了，发送端会出现阻塞，如果通道内没有数据，则接收方会阻塞
func Bcmain() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	println(<-c)
	println(<-c)
}
