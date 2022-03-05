package Day4

import "fmt"

type peoson struct {
	name string
	age  int
}

func (p peoson) String() string {
	return fmt.Sprintf("name:%v,age:%v\n", p.name, p.age)
}

// Stringer是一个普遍的接口，个人理解是这个接口可以实现格式化输出，
func SMain() {
	a := peoson{"a", 11}
	b := peoson{"b", 24}
	fmt.Println(a, b)
}
