package Day5

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func Racmain() {
	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c1)
	for v := range c1 {
		fmt.Println(v)
	}
}
