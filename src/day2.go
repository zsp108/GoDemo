package src

import (
	"fmt"
	"math"
	"runtime"
	"time"
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

//if 简短写法和else
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// 由于v是if内进行赋值的，所以v 的作用域仅在if 和else内，
		fmt.Printf("%g>=%g \n", v, lim)
	}
	return lim
}
func myif() {
	fmt.Println("======== if =========")
	fmt.Printf(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
}

// Switch 就类似多个if else嵌套起来执行，当匹配到合适的值会停止运行，或者没有合适值的时候会调用default结果，如果Switch没有条件的话，会一直为真
func myswitch() {
	fmt.Println("=========switch 循环流程==========")
	fmt.Println("=========查询本服务运行环境==========")
	switch osname := runtime.GOOS; osname {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s \n", osname)
	}
	fmt.Println("=========Switch 求值顺序==========")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("timorrow")
	case today + 2:
		fmt.Println("in tow days")
	default:
		fmt.Println("too far away")
	}
	fmt.Println("=========Switch 无值例子==========")
	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("早上好")
	case t.Hour() < 17:
		fmt.Println("下午好")
	default:
		fmt.Println("晚上好")

	}

}

// defer 语句会将函数推迟到外层函数执行后再执行，推迟的函数会放进栈，然后逐个回去输出,
//可以理解为defer语句存在栈最低端，等函数调用完后，再从栈顶逐个输出
func mydefer() {
	fmt.Println("==========defer 实例==============")
	defer fmt.Println("wold")

	fmt.Println("hello")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

//第二天学习启动入口，今天学习流程控制
func Day2Start() {
	myfor()
	myif()
	myswitch()
	mydefer()
}
