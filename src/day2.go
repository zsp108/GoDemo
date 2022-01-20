package src

import (
	"fmt"
	"math"
)

// for 流程控制
// for 循环初始化语句和后置语句可选，但是一定要有，没有的话可能会产生死循环,其实初始化语句和后置语句去掉之后，再去掉;号就是while语句
func myfor() {
	fmt.Println("========for循环=========")
	sum := 0
	for i := 0; i < 10; i++ {
		sum = i + sum
	}
	fmt.Println(sum)

	j := 1
	for j < 100 {
		j += j
	}
	fmt.Println(j)
}

//if
// 正常写法
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//简短写法
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func myif() {
	fmt.Println("======== if =========")
	fmt.Printf(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10))
}

//第二天学习启动入口，今天学习流程控制
func Day2Start() {
	myfor()
	myif()
}
