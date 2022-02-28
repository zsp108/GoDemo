package Day5

import "fmt"

func myfibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// go的select其实有点像Switch，只不过select 需要配合channel使用，每个case 需要有个通信channel，select 会阻塞到某个分支可以继续执行为止，
//多个分支准备好时，会随机选择一个执行，如果有default子句，则执行该语句，如果没有则select 会阻塞到某个case可执行

func Smain() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	myfibonacci(c, quit)
}
