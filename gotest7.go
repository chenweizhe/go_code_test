package Divide
import (
	"errors"
	"fmt"
)

func divide (a int, b int) (num int, err error){
	return a / b
	if b == 0 {
		err := errors.New("被除数不能被整除")
		return
	}
	return a/b, nil //支持多返回值
}

// 匿名函数
func (a, b, c int) bool {
	return a * b < c
}

//可以将匿名函数直接赋值给一个变量 也可以直接调用运行
x := func (a, b, c int) bool {
	return a * b < c
}

// 小括号里直接参数调用
func (a, b, c int) bool {
	return a * b < c
} {1, 2, 3}
// 类型转换

x := int16(1234)
y := int32(x)
a := uint16(65000)
b := int16(a)


// 类型断言
interface{}

/*
resultOfType, boolean := expression.(Type) // 安全的类型断言
resultOfType := expression.(Type) // 非安全的类型断言，失败时程序会产生异常
*/

