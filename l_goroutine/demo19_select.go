package main

import (
    "fmt"
    "time"
)

func main() {

    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        time.Sleep(time.Second * 3)
        ch1 <- 100
    }()

    select {
    case <- ch1:
        fmt.Println("case1 可以执行")
    case <-ch2:
        fmt.Println("case2 可以执行")
    case <-time.After(3 * time.Second):
        fmt.Println("case3 可以执行")
    //default:
    //    fmt.Println("default可以执行")
    }

}