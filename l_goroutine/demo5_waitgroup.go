package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup
func main() {
    /**
    WaitGroup: 同步等待组
        Add(), 设置等待组中要执行的子  goroutine的数量
        Wait(), 让主goroutine处于等待
        Done(), 让等待组中的counter计数器的值减1，同Add(-1)
     */

    wg.Add(2)

    go fun1()
    go fun2()

    fmt.Println("main 进入阻塞状态， 等待wg中的子goroutine结束。。")
    wg.Wait() // 表示main goroutine进入等待， 意味着阻塞
    fmt.Println("main 解除阻塞")

}

func fun1() {
    for i := 0; i < 10; i++ {
        fmt.Println("....A", i)
    }
    wg.Done()
}

func fun2() {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        fmt.Println(".....B", i)
    }
}