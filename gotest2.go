package main
import (
	"fmt"
	"unsafe"
)

/***常量****/
// 常量关键字 const
const limit = 512
const top uint16 = 1421
const Pi float64 = 3.1415926
const x,y int = 1,3 //多重赋值

// 在定义多常量时 不需要重复使用const
const (
	Cyan = 0
	Black = 1
	White = 2
)

// go的预定义常量 true false iota, iota是一个可以被编译器修改的常量
// iota在const的修饰下重置为0
// 在下一个const出现之前 每出现一次iota 所代表的数字加1
const (
	a = iota //0
	b = iota //1
	c = iota //2
)

const d = iota //0

/********数据类型********/
// go提供十一种整形

/*
byte	等同于 uint8
int	依赖于不同平台下的实现，可以是 int32 或者 int64
int8	[-128, 127]
int16	[-32768, 32767]
int32	[-2147483648, 2147483647]
int64	[-9223372036854775808, 9223372036854775807]
rune	等同于 int32
uint	依赖于不同平台下的实现，可以是 uint32 或者 uint64
uint8	[0, 255]
uint16	[0, 65535]
uint32	[0, 4294967295]
uint64	[0, 18446744073709551615]
uintptr	一个可以恰好容纳指针值的无符号整型（对 32 位平台是 uint32, 对 64 位平台是 uint64）
*/

// unsafe.Sizeof查看字节长度 
func main()  {
	fmt.Println("hello world!")
	a := 12
	fmt.Println("length of a:", unsafe.Sizeof(a))
	var b int = 12
	fmt.Println("length of b:", unsafe.Sizeof(b))
	var c int8 = 12
	fmt.Println("length of c:", unsafe.Sizeof(c))
	var d int16 = 12
	fmt.Println("length of d:", unsafe.Sizeof(d))
	var e int32 = 12
	fmt.Println("length of e:", unsafe.Sizeof(e))
	var f int64 = 12
	fmt.Println("length of f:", unsafe.Sizeof(f))

}
