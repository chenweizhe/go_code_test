package main

import (
	"fmt"
)

func main()  {
	t0 := "\u6B22\u8FCE\u6765\u5230" // t0内容：欢迎来到
    t1 := "\u5B9E\u9A8C\u697C"       // t1内容：实验楼
    t2 := t0 + t1
    for index, char := range t2 {
        fmt.Printf("%-2d    %U      '%c'    %X      %d\n",
            index, char, char, []byte(string(char)), len([]byte(string(char))))
    }
    fmt.Printf("length of t0: %d, t1: %d, t2: %d\n", len(t0), len(t1), len(t2))
    fmt.Printf("content of t2[0:2] is: %X\n", t2[0:2])
	text := "\u5B9E\u9A8C\u697C"
    fmt.Printf("bool output:\n%t\n%t\n\n", true, false)
    fmt.Println("number output, origin value: 64")
    fmt.Printf("|%b|%8b|%-8b|%08b|% 8b|\n", 64, 64, 64, 64, 64)
    fmt.Printf("|%x|%8x|%-8x|%08X|% 8X|\n\n", 64, 64, 64, 64, 64)
    fmt.Println(`text output, origin value: \u5B9E\u9A8C\u697C`)
    fmt.Printf("content: %s\n", text)
    fmt.Printf("hex value: % X\nUnicode value: ", text)
    for _, char := range text {
        fmt.Printf("%U ", char)
    }
    fmt.Println()
    bytes := []byte(text)
    fmt.Printf("value of bytes: %s\n", bytes)
    fmt.Printf("hex value of bytes: % X\n", bytes)
    fmt.Printf("origin value of bytes: %v\n", bytes)
}