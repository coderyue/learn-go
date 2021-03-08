package main

import "fmt"

func main() {

    /**
    一个goroutine打印数字， 一个goroutine打印字母，观察运行结果
     */
    //可以发现当main执行完之后，就会停止所有goroutine， 所以需要通信，channel
    go printNum()
    for i := 0; i < 100; i++ {
        fmt.Println("主goroutine中的字母: ", i,  "A")
    }

}

func printNum()  {
    for i := 0; i < 100; i++ {
        fmt.Println("子goroutine中的数字:", i)
    }
}
