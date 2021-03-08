package main

import (
    "fmt"
    "time"
)

func main() {
    /**
    关闭通道：close(ch)
        子goroutine： 写出10个数据
            每写一个，阻塞一次，主goroutine读取一次，解除阻塞

        主goroutine，读取数据
            每次读取数据，阻塞一次，子goroutine，写出一个，解除阻塞
     */

    ch := make(chan int)
    go sendData(ch)

    // 读取通道数据
    for {
        time.Sleep(1 * time.Second)
        v, ok := <- ch
        if !ok {
            fmt.Println("已经读取完毕", ok, v) // 当通道关闭的时候读取到的是0
            break
        }
        fmt.Println("读取的数据: ", v, ok)
    }

    fmt.Println("main over...")

}

func sendData(ch chan int) {
    for i := 0; i < 10; i++ {
        ch <- i // 将i写入到通道中
    }
    close(ch) //  将ch通道关闭
}