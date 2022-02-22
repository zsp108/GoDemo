package Day4

import "fmt"

// 类型断言，提供了一个访问接口底层具体值的方法，
//t := i.(T)
//该语句断言接口值 i 保存的T类型值，并把符合T类型的底层值赋予t，如果没有T类型的底层值，则将nil赋予t，
// 类型断言会返回两个值，一个是T类型的底层值，一个是判断断言是否正确的布尔值，t, ok := i.(T)
// 若i保存了T类型值，则t等于i的底层值，ok等于true，若i没有T类型值，则t等于nil，ok等于false，程序不会产生恐慌，但是如果是接受值仅t没有ok，而且断言类型不符，则会报错panic恐慌
func Tastart() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)
	//f = i.(float64)
	//fmt.Println(f)
}
