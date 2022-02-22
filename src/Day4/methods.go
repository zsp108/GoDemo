package Day4

import (
	"fmt"
	"math"
)

//方法，go语言中没有类，不过可以通过给结构体定义方法的方式，实现类的效果
//语法如下,和函数不同的是，在func 和方法名之间带上接收者，其实方法就是带了特殊接收者的函数，结构体内设置初始值，调用时对结构体进行初始化
//方法只能对同一包内定义的类型接收者声明，而不能对其他包定义的类型，包括int之类的内建类型接收者声明方法，也就是说，除了下面示例的结构体声明，还可以支持type Myfloat float64
//然后结构体.带有接收者的方法进行调用,
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//同样，方法的接收者也可以使用指针语法，和方法接收者规则一样，只能接受同一个包内的结构体或者自建类型，不支持类似*int等内建类型，
// 使用指针接收者和值接收者的区别在于，指针接收者可以修改接收者指向的值，如果方法需要经常修改他的接收者，指针接收者比值接收者要好用很多
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func Mymethod() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

// 指针与函数区别示例
func Fabs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Fscale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func mymethod_pointers_explained() {
	v := Vertex{3, 4}
	//fmt.Println(Fabs(v))
	Fscale(&v, 10)
	//v.Scale(10)
	//fmt.Println(Fabs(v))
	// 指针重定向，函数指针参数必须引入指针才可以执行，而方法指针接收者即可接受指针接收者，也可以接受值接收者，
	// 相同道理，如果函数内使用指针参数，那也只能接受指针传入的值，方法不受限制，方法即可以接受指针接收者，也可以使用值接收者
	// 选择指针或值为接接收者：一、方法可以修改器接收者指向的值，二、避免每次调用的时候复制该值，如果值是大型结构体的话，这样做更高效，
	p := &v
	p.Scale(10)
	fmt.Println(p.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())

}
