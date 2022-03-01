package Day5

import (
	"fmt"
	"time"
)

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
	//c := make(chan int)
	//quit := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	//
	//myfibonacci(c, quit)

	intChan := make(chan int, 10)
	func() {
		for i := 0; i < 10; i++ {
			intChan <- i
		}
	}()

	strChan := make(chan string, 5)
	func() {
		for j := 0; j < 5; j++ {
			strChan <- "hello " + fmt.Sprintf("%d", j)
		}
	}()

	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据-%d\n", v)
			time.Sleep(time.Second)
		case v := <-strChan:
			fmt.Printf("从stringChan读取的数据%v\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Println("没有数据可取")
			time.Sleep(time.Second)
			return

		}
	}

}
