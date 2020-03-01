/*
Go 函数中添加多个defer语句，
当函数执行到最后时，这些 defer 语句会按照逆序执行（
即最后一个defer语句将最先执行），最后该函数返回。
特别是当你在进行一些打开资源的操作时，
遇到错误需要提前返回，在返回前你需要关闭相应的资源，
不然很容易造成资源泄露等问题
*/

package main

import "fmt"

func test() (result int) {
    defer func() {
        result = 12
    }()
    return 10
}

func main() {
    fmt.Println(test())     // 12
}