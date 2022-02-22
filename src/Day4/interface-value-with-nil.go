package Day4

import "fmt"

type i interface {
	M()
}

type tt struct {
	s string
}

func (t *tt) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.s)
}
func describe(i i) {
	fmt.Printf("%v %T\n", i, i)
}

func describe1(i interface{}) {
	fmt.Printf("%v %T\n", i, i)
}

// 1、当接口内的具体值即结构体底层值为nil时，方法仍然可以被nil接收者调用，并不会触发空指针异常
// 不管你的接口是否有具体值，只要有方法实现该接口，即使传的是nil底层值，程序还是可以正常运行的，因为可以识别出实现方法的值类型，所以不会有空指针异常，但是一般会做一些展示优化
// 2、如果是一个nil接口，没有方法实现它，那运行的时候会产生报错，nil值既不保存值，也不保存具体类型，输出nil接口的值和类型的时候都是“<nil> <nil>”
// 3、空接口就是指定了零个方法的接口，空接口可以接受所有值，因为每个类型都实现了零个方法，一般用来处理未知类型的值
func Ivwnstart() {
	var t1 *tt
	var i1 i = t1
	describe(i1)
	i1.M()
	i1 = &tt{"Hello"}
	describe(i1)
	i1.M()

	fmt.Println("i2 nil接口值示例")
	var i2 i
	describe(i2)
	//i2.M()

	fmt.Println("空接口 示例")
	var i3 interface{}
	describe1(i3)
	i3 = 4
	describe1(i3)
	i3 = "hello"
	describe1(i3)
}
