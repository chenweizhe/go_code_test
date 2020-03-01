/*
Go 语言中的面向对象和 C++，Java 中的面向对象不同，
因为 Go 语言不支持继承，Go 语言只支持组合
*/
package main

import "fmt"

// 自定义结构类型
// type Man struct{
// 	name string
// 	age int
// }

// // 对象的创建和初始化
// man := new(Man)
// man := &Man{}
// man := &Man{"Tom", 18}
// man := &Man{name: "Tom", age: 18}

// 为了更加方便的创建对象，
// 我们一般会使用一个全局函数来完成对象的创建，这和传统的“构造函数”类似。

// func NewMan(name string, age int) *Man {
// 	return &Man{name, age}
// }

/*
方法是作用在自定义类型上的一类特殊函数，
通常自定义类型的值会被传递给该函数，
该值可能是以指针或者复制值的形式传递。定义方法和定义函数几乎相同，
只是需要在func关键字和方法名之间必须写上接接受者。
*/

type Count int

func (count *Count) Increment() { *count++ }
func (count *Count) decrement() { *count-- }
func (count Count) IsZero() bool {return count == 0}

type Part struct {
	stat string
	Count // 匿名字段
}

func (part Part) IsZero() bool { // 覆盖了匿名字段Count的IsZero()方法
    return part.Count.IsZero() && part.stat == "" // 调用了匿名字段的方法
}

func (part Part) String() string { // 定义String()方法，自定义了格式化指令%v的输出
    return fmt.Sprintf("<<%s, %d>>", part.stat, part.Count)
}

func main() {
    var i Count = -1
    fmt.Printf("Start \"Count\" test:\nOrigin value of count: %d\n", i)
    i.Increment()
    fmt.Printf("Value of count after increment: %d\n", i)
    fmt.Printf("Count is zero t/f? : %t\n\n", i.IsZero())
    fmt.Println("Start: \"Part\" test:")
    part := Part{"232", 0}
    fmt.Printf("Part: %v\n", part)
    fmt.Printf("Part is zero t/f? : %t\n", part.IsZero())
    fmt.Printf("Count in Part is zero t/f?: %t\n", part.Count.IsZero()) // 尽管覆盖了匿名字段的方法，单还是可以访问

}


