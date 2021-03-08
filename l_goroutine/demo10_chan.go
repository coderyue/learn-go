package main

import (
    "fmt"
)

func main() {

    /**
    从通道中读写数据都是阻塞的
     */

    ch := make(chan int)

    go func() {
        fmt.Println("子goroutine开始执行....")
        //time.Sleep(1 * time.Second)
        data := <- ch
        fmt.Println("data: ", data)
    }()

    // 在main的goroutine中添加睡眠， 子goroutine大概率没有执行完， 都不加或者在子goroutine加睡眠，子goroutine大都能执行完
    //time.Sleep(1 * time.Second)
    ch <- 10
    fmt.Println("main over...")
    ch2 := make(chan int)

    ch2 <- 100 // 阻塞， 没有goroutine来读数据解除阻塞
}
