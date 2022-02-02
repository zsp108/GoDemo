package src

import (
	"fmt"
	"math/rand"
)

//学习go基础语法和包函数等

//导入包学习
func pakg() {
	fmt.Println("=======初始导入包======")
	fmt.Println("My favorite number is", rand.Int())
}

//函数可以接受0个或者多个参数，多个函数入参类型相同可以简写为一个类型
func myadd(x, y int) int {
	fmt.Println("=======函数入参======")
	return x + y
}

//多值返回,单值返回无需括号，多值返回需要括号加上
func myswap(x, y string) (string, string) {
	fmt.Println("=======多值返回======")
	return x, y
}

//命名参数返回,返回值和返回命名参数不同时会报错，当函数内存在和返回命名参数相同名称参数时直接使用return返回即可
func mysplit(sum int) (x, y int) {
	fmt.Println("=======命名参数返回======")
	x = sum * 4 / 9
	y = sum - x
	return
}

//变量声明可以使用var 关键字（var 可以不带初始值,可以在函数体内或者函数体外进行声明），也可以使用 := 短变量声明（短变量声明仅使用于函数体内，并且需要赋值）,
//每个变量对应一个初始值，如果初始值已存在，可以省略声明变量类型
//如果没有明确赋值的话，数值类型为 0，布尔类型为 false，字符串为 ""（空字符串）
func myvar() {
	fmt.Println("=======变量声明======")
	//var 关键字声明变量，带有初始值
	var i, j int = 1, 2
	// var 不带变量
	//var c, python, java = true, false, "no!"
	//fmt.Println(i, j, c, python, java)
	fmt.Println(i, j)

	//短变量声明
	x, y := 3, 4
	fmt.Println(x, y)

	var a int
	var b bool
	var s string
	var f float64
	fmt.Println(a, b, s, f)
}

// 类型转换
func mytype_conversions() {
	fmt.Println("=======类型转换======")
	//var i int = 1
	//var f float64 = float64(i)
	//var z uint = uint(f)
	//fmt.Println(z)

	// 也可以使用短变量声明的转换写法
	i := 1
	f := float64(i)
	z := uint(f)
	fmt.Println(z)
}

//类型推导
// 在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出。
// 右边包含未指明类型的数值常量时，新变量的类型就可能是 int, float64 或 complex128 了，这取决于常量的精度
func mytype_inferenc() {
	fmt.Println("=======类型推导======")
	i := 42 // 修改这里！
	fmt.Printf("i is of type %T\n", i)
	f := 42.0
	fmt.Printf("f is of type %T\n", f)
	g := 0.867 + 0.5i
	fmt.Printf("g is of type %T\n", g)

}

//常量的声明与变量类似，只不过是使用 const 关键字。
//常量可以是字符、字符串、布尔值或数值。
//常量不能用 := 语法声明
const Pi = 3.14

func myconstants() {
	fmt.Println("=======常量======")
	const World = "world"
	fmt.Printf("helo %s\n", World)
	fmt.Printf("Happy %s Day\n", Pi)
	const Truth = true
	fmt.Printf("Go rules? %v \n", Truth)
}

//数值常量是高精度的 值。
//一个未指定类型的常量由上下文来决定其类型。
//再尝试一下输出 needInt(Big) 吧。
//（int 类型最大可以存储一个 64 位的整数，有时会更小。）
//（int 可以存放最大64位的整数，根据平台不同有时会更少。）
const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 - 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
func mynumeric() {
	fmt.Println("=======数值常量======")
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

//当前学习启动入口
func Day1Start() {
	// 第一天学习的基础调用
	fmt.Println("第一天学习包、变量和函数等基础调用")
	fmt.Println("hello world")
	pakg()
	fmt.Println(myadd(1, 2))
	fmt.Println(myswap("Hello", "world!"))
	m, n := mysplit(17)
	fmt.Println(m, n) // 7 10
	myvar()
	mytype_conversions()
	mytype_inferenc()
	myconstants()
	mynumeric()

}
