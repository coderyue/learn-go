package main

import (
    "fmt"
    "time"
)

func main() {
    /**
    通过range访问通道
     */
    ch := make(chan int)
    go SendData(ch)

    for v := range ch { // v <- ch
        fmt.Println("读取到的数据:", v)
    }

    fmt.Println("main over...")
}

func SendData(ch chan int)  {
    for i := 0; i < 10; i++ {
        time.Sleep(1 * time.Second)
        ch <- i
    }
    close(ch) // 通知对方， 通道关闭， 否则会死锁
}