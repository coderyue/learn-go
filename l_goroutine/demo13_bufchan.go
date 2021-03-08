package main

import (
    "fmt"
    "strconv"
)

func main() {
    /**
    非缓冲通道： make(chan T)
        一次发送上一次接收， 都是阻塞的
    缓冲通道： make(chan T, capacity)
        发送： 缓冲区满了， 才会阻塞
        接受： 缓冲区空了， 才会阻塞
     */

    ch1 := make(chan int) // 非缓冲通道
    fmt.Println(len(ch1), cap(ch1)) // 0 0
    //ch1 <- 100 // 阻塞式的， 需要有其他的goroutine解除阻塞， 否则deadlock

    ch2 := make(chan int, 5)  // 缓冲通道，缓冲区大小是5
    fmt.Println(len(ch2), cap(ch2)) // 0 5

    ch2 <- 100
    ch2 <- 200
    ch2 <- 300
    ch2 <- 400
    ch2 <- 500
    fmt.Println(len(ch2), cap(ch2))
    //ch2 <- 600

    fmt.Println("---------------------")
    ch3 := make(chan string, 4)
    go SendData(ch3)

    //for {
    //    v, ok := <- ch3
    //    if !ok {
    //        fmt.Println("读完了...", v)
    //        break
    //    }
    //    fmt.Println("\t读取的数据:", v)
    //}
    for v := range ch3 { //range遍历好像是等通道满了再一下子都读出来
        fmt.Println("\t读取的数据:", v)
    }
    fmt.Println("main over...")

}

func SendData(ch chan string) {
    for i := 0; i < 10; i++ {
        ch <- "数据" + strconv.Itoa(i)
        fmt.Println("字goroutine写入数据:", i)
    }
    close(ch)
}