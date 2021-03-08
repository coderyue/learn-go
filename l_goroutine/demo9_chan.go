package main

import "fmt"

func main() {

    /**
    channel
        1. 用于channel，传递消息的
        2. 通道，每个都有相关联的数据类型
            nil chan 不能使用
        3. 使用通道传递数据 <-
        4. 阻塞
            发送数据：chan <- data,阻塞的， 直到另一条goroutine，读取数据来解除阻塞
            读取数据：data <- chan，也是阻塞的。直到另一条goroutine，写出数据解除阻塞
        5. 本身channel就是同步的，意味着同一时间，只能有一条goroutine来操作
     */

    var ch chan bool
    ch = make(chan bool)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println("字goroutine中 i:", i)
        }
        // 循环结束， 向通道中写数据， 表示要结束了
        ch <- true
        fmt.Println("结束...")
    }()

    data := <- ch // 阻塞
    fmt.Println("main---data-->", data)
    fmt.Println("maven over。。。")

}
