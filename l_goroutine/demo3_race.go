package main

import (
    "fmt"
    "runtime"
)

func main() {
    /*
    临界资源. 多个goroutine共享的资源

     */

    a := 1
    go func() {
        a = 2
        fmt.Println("goroutine...", a)
    }()

    a = 3
    runtime.Gosched() // 让出时间片
    fmt.Println("main...", a)
}
