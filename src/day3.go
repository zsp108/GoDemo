package src

import "fmt"

// 指针是保存了值在内存的地址，类型*T 是指向T类型值的指针，其零值为nil，指针使用&指向其他操作数，
// *操作符表示指针指向的底层值
func mypointers() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

//结构体，其实就是一组字段，类似java的实体类，结构体字段使用点号来访问，
type Vertex struct {
	X, Y int
}

var (
	//结构体文法
	v1 = Vertex{1, 2}  // 利用var 来赋值结构体
	v2 = Vertex{X: 1}  //没有赋值的时候Y默认为：0
	v3 = Vertex{}      // 没有赋值都是默认为0
	p  = &Vertex{1, 2} // 创建结构体指针
)

func mystruct() {
	fmt.Println("结构体学习")
	// 正常结构体实例化和调用，结构体字段使用点号来访问
	v := Vertex{1, 2}
	v.Y = 4
	fmt.Println(v.Y)
	a := new(Vertex)
	a.X = 7
	a.Y = 8
	fmt.Printf("X:%d,Y:%d \n", a.X, a.Y)

	// 结构体指针，和指针相同，通过&指向实例化出来的结构体，可以修改内存中已存在的值
	fmt.Println("==========结构体指针=============")
	b := Vertex{1, 2}
	p := &b
	p.X = 1e9
	fmt.Println(b)

	fmt.Println("===========结构体文法=============")
	fmt.Println(v1, v2, v3, p)

}

// 数组
// 和其他语言相似，数组是存储多个值的一个组合
// [n]T 表示拥有n个T类型值的数组
// var a [10]int 声明a为拥有10个整数的数组
func myarray() {
	var a [2]string                                         // 无值初始化数组
	var b = [10]float64{1000.0, 2.0, 3.4, 7.0, 50.0}        //带值初始化。初始化值{}内的长度不能大于数组[]限制长度，如果是[]放空，则是自由设置，和下面的...一样
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0} // 无固定长度初始化
	balance1 := [5]float32{1: 2.0, 3: 7.0}                  //  将索引为 1 和 3 的元素初始化

	fmt.Println(a, b, balance, balance1)
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

// 切片
// 每个数组的大小都是固定的，二切片则为数组提供动态大小的，灵活的视角，
// 类型 []T 表示一个元素类型为 T 的切片。
// 切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：a[low : high]它会选择一个半开区间，包括第一个元素，但排除最后一个元素。
// 以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素a[1:4]
func myslices() {

}

func Day3Start() {
	fmt.Println("第三天学习指针，struct，slice和映射")
	mypointers() // 指针学习调用
	mystruct()   // 结构体学习调用
	myarray()    // 数组学习调用
	myslices()   // 切片学习调用
}
