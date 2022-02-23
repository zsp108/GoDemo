package Day4

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (me MyError) Error() string {
	return fmt.Sprintf("at %v, %s", me.When, me.What)
}
func run() error {
	return MyError{time.Now(), "zzzzz"}
}

type DiviedError struct {
	diviee int
	divier int
}

func (d *DiviedError) Error() string {
	strFormat := `
	除数不能为零，diviee：%d，divier:0
`
	return fmt.Sprintf(strFormat, d.diviee)
}

func Divide(vardivdee int, vardivder int) (int, error) {
	if vardivder == 0 {
		var i int
		vdata := &DiviedError{vardivdee, vardivder}
		return i, vdata
	} else {
		return vardivdee / vardivder, nil
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))

}

func MSqrt(x float64) (float64, error) {
	if x < 0 {
		se := ErrNegativeSqrt(x)
		return 0, se
	} else {
		return math.Sqrt(x), nil
	}
}

// error,go使用error来表示错误，error和Stringer类似，也是一个内建类型，可以通过重写Error方法，返回特定的错误信息，然后通过判断返回的错误信息，是否为nil，是则失败，否则成功
func Emain() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// 除数不能为0
	if res, errmsg := Divide(10, 0); errmsg != nil {
		fmt.Printf("错误信息：%v", errmsg)
	} else {
		fmt.Printf("结果集：%d,错误信息：%v", res, errmsg)
	}
	fmt.Println(MSqrt(9))
	fmt.Println(MSqrt(-2))
}
