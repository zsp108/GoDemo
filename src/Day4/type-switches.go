package Day4

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%d %T\n", v, v)
	case string:
		fmt.Printf("%v %T\n", v, v)
	default:
		fmt.Printf("未知类型")
	}
}
func Tsmain() {
	fmt.Println("类型选择")
	do("Hello")
	do(11)
	do(true)
}
