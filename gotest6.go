package main

// 函数调用非常简单，先将被调用函数所在的包导入，就可以直接使用该函数了
import (
	"./Add.go"
	"fmt"
)

func main()  {
	c = Add.add(1, 2)
	fmt.Println(c)
}