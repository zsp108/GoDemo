package Day5

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//goroutine 是由go运行时管理的轻量级线程，调用模式为：go funcname(x,y,x),启动的时候会启动一个新线程并执行
//而我们普通调用func(x,y,z) 则是发生在当前go线程内的，

func Startgoroutine() {
	go say("Hello")
	say("world")
}
