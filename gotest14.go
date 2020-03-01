/*
channel是goroutine之间互相通讯的东西。
类似我们 Unix 上的管道（可以在进程间传递消息），
用来goroutine之间发消息和接收消息。
其实，就是在做goroutine之间的内存共享。channel是类型相关的，
也就是说一个channel只能传递一种类型的值，
这个类型需要在channel声明时指定。
*/

var a chan int //声明一个传递元素类型为int的channel
var b chan float64
var c chan string

// 初始化一个channel也非常简单
a := make(chan int) //初始化一个int型的名为a的channel
b := make(chan float64)
c := make(chan string)

// channel最频繁的操作就是写入和读取，这两个操作也非常简单
