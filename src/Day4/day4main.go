package Day4

import "fmt"

func Day4Start() {
	fmt.Println("第四天学习方法和接口")
	//总结来说，方法即是待了接收者的函数，接收者必须在同一包内，可以是结构体，或者自定义类型，类似type Myfloat float64,不支持其他包或者int之类的内建类型
	// 方法和函数都支持值和指针，不同的是，函数参数如果是值类型的，那传入参数只能是值，如果传入参数为指针对象，那只能传到包含指针参数的函数中，传入参数为值对象的只能在参数为值的函数中使用
	//指针类型的只能是指针类型，但是方法可以是值传递也可以指针传递，如果方法的接收者为指针对象的时候，Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)
	//当方法的接收者为指针对象时，修改接收者会直接修改接收者指向的结构体内的值，这样可以避免每次调用方法时复制该值，在大型结构体使用上效率更高效
	Mymethod()
	fmt.Println("=====指针与函数=====")
	mymethod_pointers_explained()
	fmt.Println("=====接口=====")
	Interfacemain()
	fmt.Println("接口nil值")
	Ivwnstart()
	fmt.Println("类型断言")
	Tastart()
	fmt.Println("类型断言选择")
	Tsmain()
	fmt.Println("Stringer 接口")
	SMain()
	fmt.Println("Error 接口")
	Emain()
	fmt.Println("Reader 接口")
	Rmain()
}
