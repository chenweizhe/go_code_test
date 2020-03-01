/*
	go支持两种字符类型 一种是byte（uint8的别名）
	另一个是rune 代表单个Unicode字符
*/
/*
数组 是一个定长序列 其中元素类型相同
省略符go会自动计算数组长度
Go 语言的切片比数组更加灵活，强大而且方便。数组是按值传递的（即是传递的副本），而切片是引用类型，传递切片的成本非常小，
而且是不定长的。而且数组是定长的，而切片可以调整长度。
另外可以通过内置的函数append()来增加切片的容量
*/
package main

import (
	"fmt"
)

func main()  {
	
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
    fmt.Printf("len and cap of array %v is: %d and %d\n", a, len(a), cap(a))
    fmt.Printf("item in array: %v is:", a)
    for _, value := range a {
        fmt.Printf("% d", value)
    }

    fmt.Println()

    s1 := a[3:6]
    fmt.Printf("len and cap of slice: %v is: %d and %d\n", s1, len(s1), cap(s1))
    fmt.Printf("item in slice: %v is:", s1)
    for _, value := range s1 {
        fmt.Printf("% d", value)
    }

    fmt.Println()

    s1[0] = 456
    fmt.Printf("item in array changed after changing slice: %v is:", s1)
    for _, value := range a {
        fmt.Printf("% d", value)
    }

    fmt.Println()

    s2 := make([]int, 10, 20)
    s2[4] = 5
    fmt.Printf("len and cap of slice: %v is: %d and %d\n", s2, len(s2), cap(s2))
    fmt.Printf("item in slice %v is:", s2)
    for _, value := range s2 {
        fmt.Printf("% d", value)
    }

    fmt.Println()

}
