/*
select用于处理异步 IO 问题，
它的语法与switch非常类似。由select开始一个新的选择块，
每个选择条件由case语句来描述，
并且每个case语句里必须是一个channel操作。
它既可以用于channel的数据接收，也可以用于channel的数据发送。
如果select的多个分支都满足条件，
则会随机的选取其中一个满足条件的分支。
*/

package main
import "time"
import "fmt"
func main() {
    c1 := make(chan string)
    c2 := make(chan string)
    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "two"
    }()
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}

// 超时机制
// t := make(chan bool)
// go func {
//     time.Sleep(1e9) //等待1秒
//     t <- true
// }

// select {
//     case <-ch:  //从ch中读取数据

//     case <-t:  //如果1秒后没有从ch中读取到数据，那么从t中读取，并进行下一步操作
// }

// channel的关闭
// ch := make(chan int)
// close(ch)

// 关闭了channel后如何查看channel是否关闭成功了
// x, ok := <-ch