package main
import (
	"fmt"
	"unsafe"
)

// go提供两种浮点型和两种复数类型
/*
float32	±3.402 823 466 385 288 598 117 041 834 845 169 254 40x1038 计算精度大概是小数点后 7 个十进制数
float64	±1.797 693 134 862 315 708 145 274 237 317 043 567 981x1038 计算精度大概是小数点后 15 个十进制数
complex32	复数，实部和虚部都是 float32
complex64	复数，实部和虚部都是 float64
*/

// 布尔类型 true flase 不支持强制转换
func main()  {
	var a bool
	a = true
	b := (2 == 3)
	fmt.Println(a)
	fmt.Println(unsafe.Sizeof(b))

	// 字符串的创建 \是转义的意思
	t1 := "\"hello\""
	t2 := `"hello"`
	t3 := "\u6B22\u8FCE"
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)

	// Go 语言中单个字符可以使用单引号( ' )来创建

}