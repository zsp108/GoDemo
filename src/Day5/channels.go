package Day5

import "fmt"

func sum(s []int, c chan int) {
	var sum = 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// go 通道，其实可以理解为一个队列模式的通道，当发送方或者接收方在准备阶段，通道处于锁状态，一般如下不带缓存参数的通道需要即发即接收，
// 通道在使用钱需要先用make创建，make函数可以创建map，chan，slice，创建使用chan关键字语法：c:=make(chan int)
// 创建后向通道投递发送数据，c <-v,输出或者给变量参数赋值可以使用：v := <-c
// 以下例子采用go 开启两个sum的线程计算list的，同时在主线程内通过赋值给xy打印，

func ChannelMain() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

}
