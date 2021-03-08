package main

import "fmt"

func main() {
    /*
    双向：
        chan T
            chan <- data , 发送数据，写出
            data <- chan， 获取数据，读取

    单向：定向
        chan <- T, 只支持写
        <- chan T, 只读

     */

    ch1 := make(chan string)
    done := make(chan bool)
    go SendData(ch1, done)

    data := <- ch1 // 读取
    fmt.Println("子goroutine传来：", data)
    ch1 <- "main" // 发送

    // 这个布尔的通道就是当子的goroutine结束后写个数据通知主goroutine结束了， 在主goroutine没有收到这个数据时先不要结束
    // 这里没有接收数据的变量， 就是做了个阻塞, 因为接收了也不使用
    <- done
    fmt.Println("main over....")

}

func SendData(ch chan string, done chan bool) {
    ch <- "lin"
    data := <- ch // 读取
    fmt.Println("main goroutine传来：", data)

    done <- true // true或false都可以， 就是通知主goroutine子goroutine结束了
}
