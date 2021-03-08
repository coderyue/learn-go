package main

import (
    "fmt"
    "time"
)

func main() {
    /**
      单向：定向
          chan <- T, 只支持写
          <- chan T, 只读
     */

    ch1 := make(chan int) // 双向通道
    //ch2 := make(chan <- int) // 只能写
    //ch3 := make(<- chan int) // 只能读

    //ch1 <- 100
    //ch2 <- 100
    //data := <- ch2 //invalid operation: <-ch2 (receive from send-only type chan<- int)

    //data := <- ch3
    //ch3 <- 100 //invalid operation: ch3 <- 100 (send to receive-only type <-chan int)

    //go fun1(ch1)
    //go fun1(ch2) // 只写
    //data := <- ch1
    //fmt.Println("fun1读取到数据", data)

    go fun2(ch1)
    go fun1(ch1)
    //ch1 <- 100
    time.Sleep(1 * time.Second)
    fmt.Println("main over...")
}

// 该函数，只能操作只写通道，，，单向通道用于函数中， 限制通道在函数内只读或者只写
func fun1(ch chan <- int)  {
    // 在函数内部， 对于ch1通道，只能写数据，不能读取数据
    ch <- 100 // 写入数据
    fmt.Println("fun1函数结束..。")
}

// 该函数只能操作只读通道
func fun2(ch <- chan int)  {
    data := <- ch
    fmt.Println("fun2读取到数据：", data)
}