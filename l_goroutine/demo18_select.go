package main

import (
    "fmt"
    "time"
)

func main() {
    /*
    分支语句： if， switch， select
    select语句类型与switch语句类似
        但是select语句会随机执行一个可运行的case
        如果没有可运行的case， 要看是否有default，如果有就执行default， 否则进入阻塞，直到有case可以运行

    select中每个case必须是一个通信

     */

    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        time.Sleep(time.Second * 3)
        ch1 <- 100
    }()

    go func() {
        time.Sleep(time.Second * 3)
        ch2 <- 200
    }()

    select {
    case num1 := <-ch1:
        fmt.Println("ch1中获取数据:", num1)
    case num2, ok := <-ch2:
        if ok {
            fmt.Println("ch2中获取数据：", num2)
        } else {
            fmt.Println("ch2通道已关闭")
        }
    default:
        fmt.Println("default语句")
    }
}