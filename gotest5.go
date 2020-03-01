package main
/*
Go 语言提供的流程控制语句包括if、switch、for、goto、select,
 其中select用于监听 channel（通道）。
*/
import (
	"fmt"
)

//Go 语言中使用goto关键字实现跳转。
// goto语句的语义非常简单，就是跳转到本函数内的某个标签

func myfunc()  {
	i := 0
	THIS: //定义一个THIS标签
	fmt.Println(i)
	i++
	if i < 1 {
		goto THIS //跳转到THIS标签
	}
}


func classchecker(items ...interface{}) {
	for i,x := range items {
		switch x := x.(type) { //创建影子变量
		case bool:
			fmt.Printf("param #%d is a bool, value: %t\n", i, x)
		case float64:
            fmt.Printf("param #%d is a float64, value: %f\n", i, x)
        case int, int8, int16, int32, int64:
            fmt.Printf("param #%d is a int, value: %d\n", i, x)
        case uint, uint8, uint16, uint32, uint64:
            fmt.Printf("param #%d is a uint, value: %d\n", i, x)
        case nil:
            fmt.Printf("param #%d is a nil\n", i)
        case string:
            fmt.Printf("param #%d is a string, value: %s\n", i, x)
        default:
            fmt.Printf("param #%d's type is unknow\n", i)
		}
	}
}

func main() {
	classchecker(-5, -17.98, "AIDEN", nil, true, complex(1,1))
}