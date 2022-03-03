package Day5

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (sc *SafeCounter) Inc(key string) {
	sc.mux.Lock()
	sc.v[key]++
	sc.mux.Unlock()
}

func (sc *SafeCounter) Values(key string) int {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	return sc.v[key]
}

// 我们已经看到信道非常适合在各个 Go 程间进行通信。
// 但是如果我们并不需要通信呢？比如说，若我们只是想保证每次只有一个 Go 程能够访问一个共享的变量，从而避免冲突？
// 这里涉及的概念叫做 *互斥（mutual*exclusion）* ，我们通常使用 *互斥锁（Mutex）* 这一数据结构来提供这种机制。
// Go 标准库中提供了 sync.Mutex 互斥锁类型及其两个方法：
// Lock
// Unlock
// 我们可以通过在代码前调用 Lock 方法，在代码后调用 Unlock 方法来保证一段代码的互斥执行。参见 Inc 方法。
// 我们也可以用 defer 语句来保证互斥锁一定会被解锁。参见 Value 方法。
func McStart() {
	var s = SafeCounter{v: make(map[string]int)}
	for i := 1; i < 10; i++ {
		s.Inc("zsp")
	}

	time.Sleep(time.Second)
	fmt.Println(s.Values("zsp"))
}
