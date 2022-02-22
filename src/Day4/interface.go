package Day4

import (
	"fmt"
	"math"
)

// GO的接口和java等语言的接口意思相同，接口也是一种数据类型，它把所有具有共性的方法定义在一起，任何其他类型的只要实现了这些方法就是实现了这个接口，
//接口内写上需要实现的方法名和返回值，如下
type Abser interface {
	Abs() float64
}

type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64((-f))
	}
	return float64(f)
}

type Verte struct {
	X, Y float64
}

func (v *Verte) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 接口隐式实现示例
type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func Interfacemain() {
	var a Abser
	mf := Myfloat(-math.Sqrt2)
	v := Verte{3, 4}
	a = mf
	fmt.Println(a.Abs())
	a = &v
	//a = v 在方法中，接受者是不区分指针或者值，但是接口中，Verte的Abs方法实现的是指针类型的Abs()，所以调用的时候也是需要使用指针类型的Abs()
	fmt.Println(a.Abs())

	// 接口隐式实现不用提前定义接口，直接使用以下语法可以定义实现
	fmt.Println("接口隐式转换")
	var i I = T{"Hello"}
	i.M()
	// 接口值，接口也是值，它们可以像值一样进行传递，
	// 当接口内的具体值即结构体底层值为nil时，方法仍然可以被nil接收者调用，并不会触发空指针异常
	//	但是当接口值为nil时，nil 接口值既不保存值，也不保存具体类型
	fmt.Printf("(%v ,%T)\n", i, i)
}
