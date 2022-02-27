package Day5

import "fmt"

func Day5Start() {
	fmt.Println("第五天学习并发")
	//Startgoroutine()
	fmt.Println("Go channel通道学习")
	ChannelMain()
	fmt.Println("带缓冲的channel通道")
	Bcmain()
	fmt.Println("通道的range 和close")
	Racmain()
	//RacStart()
}
