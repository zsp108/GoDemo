package src

import (
	"fmt"
	"math"
)

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
	//简单切片
	primes := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	j := primes[0:4]
	fmt.Println(j)
	//切片不存储数据，它只是描述底层数组中的一段，更改切片元素会修改底层数组
	// 切片默认下界为0，上界为数组长度
	j[0] = 12
	fmt.Println(j)
	fmt.Println(primes)
	//数组和切片的长度可以用len来求出，容量可以用cap函数来求出
	fmt.Println(len(j))
	fmt.Println(cap(j))
	//切片的零值是nil
	//nil的切片长度和容量都是0,且是没有底数的数组
	var s []int
	fmt.Printf("零值切片长度:%d,容量:%d\n", len(s), cap(s))

	// 使用make 创建切片
	// make([]type,len,cap)// len(b)=len, cap(b)=cap
	var a = make([]int, 5, 10) // len(b)=5, cap(b)=10
	fmt.Printf("零值切片长度:%d,容量:%d ,数组：%v ;\n", len(a), cap(a), a)

	//使用append函数向切片追加新元素，使用上面的s []int做示例，s切片是一个nil切片,
	// 上面说nil 切片的长度和容量都是0，如果追加新元素的容量大于底层数组容量，它将分配一个更大的数组，返回的切片也会指向新分配的数组
	fmt.Printf("空切片s：%v\n", s)
	s = append(s, 1, 2, 3)
	fmt.Printf("零值切片长度:%d,容量:%d ,数组：%v ;\n", len(s), cap(s), s)

}

// for 循环的range 用于遍历切片或者映射，range 返回两个值，一个当前下标，一个当前下标值的副本
// range 如果只需要下标，不需要值，可以直接忽略第二个参数，如果只需要值，可以使用_下划线忽略下标
func myrange() {
	var pow = []int{1, 2, 3, 4, 5, 6, 7, 8}

	for k, v := range pow {
		fmt.Printf("当前下标：%d,值：%d\n", k, v)
	}

	for _, v := range pow {
		fmt.Printf("当前下标：无需,值：%d\n", v)
	}

}

// 映射，也叫map集合，是无序键值对集合，通过key可以快速搜索到value，由于map是无序的，使用hash表实现的，所以我们无法保证返回顺序
// 映射语法： var m map[key_type]value_type
type Val struct {
	Lat, Long float64
}

func mymaps() {
	// 如果不初始化map，那map默认是一个nil map，nilmap不能用来存放键值对
	var m map[string]Val
	m = make(map[string]Val)
	m["zz"] = Val{31, 52}
	fmt.Println(m["zz"])
	//和结构体一样，我们可以使用make初始化，也可以直接使用变量声明方式初始化
	m1 := map[string]Val{"z1": Val{58, 64}, "z2": Val{1, 2}}
	fmt.Println(m1["z1"])
	// 插入新键值对，可以通过语法 m["key"] = elem 方式新增，
	// 获取键值对可以通过elem := m[key],双复制的话可以判断key是否存在elem， ok := m[key]
	// 删除键值对可以通过delete函数进行删除 delete(m,key)
	fmt.Println(m1)
	delete(m1, "z2")
	fmt.Println(m1)
	elem, ok := m1["z2"]
	fmt.Println(elem, ok)
}

//函数值，函数也是一个值，他们可以像值一样调用，函数值可以作为函数的参数或者值
//            |--可以相当于参数名--------| 函数类型   返回类型
func computer(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func myfunc_value() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(computer(hypot))
	fmt.Println(computer(math.Pow))
}

//函数值闭包，Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。
// 其实也可以理解为闭包是吧函数值绑定在函数体外的变量，用以下例子来说，就是adder函数的函数值func绑定在引用其的pos变量上，初学理解，有误请指教
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func myfunclosures() {
	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}
}

func Day3Start() {
	fmt.Println("第三天学习更多类型：指针，struct，slice和映射")
	mypointers()    // 指针学习调用
	mystruct()      // 结构体学习调用
	myarray()       // 数组学习调用
	myslices()      // 切片学习调用
	myrange()       // range学习调用
	mymaps()        // map 映射学习调用
	myfunc_value()  // 函数值学习调用
	myfunclosures() //函数值闭包学习调用
}
