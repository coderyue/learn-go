package main

import (
    "fmt"
    "runtime"
    "time"
)

func init() {
    // 获取逻辑cpu的数量
    fmt.Println("逻辑cpu的数量-->", runtime.NumCPU())

    // 设置go程序执行的最大的cpu数量：[1,256]  //1.8版本之后， 默认让程序运行再多核上
    n := runtime.GOMAXPROCS(4)
    fmt.Println(n)
}

func main() {
    /**
    runtime包,  go代码运行在runtime中，runtime相当于jvm
    管理代码执行分配内存等
     */

    // 获取goroot目录
    fmt.Println("GOROOT-->", runtime.GOROOT())
    // 获取操作系统
    fmt.Println("so/platform-->", runtime.GOOS)

    // gosched  //让出，使不阻塞
    //go func() {
    //    for i := 0; i < 5; i++ {
    //        fmt.Println("goroutine...")
    //    }
    //}()
    //
    //for i := 0; i < 4; i++ {
    //    // 让出时间片， 让其他的goroutine执行
    //    runtime.Gosched()
    //    fmt.Println("main.....")
    //}

    go func() {
        fmt.Println("goroutine开始...")
        fun()
        fmt.Println("goroutine结束....")
    }()

    time.Sleep(1 * time.Second)
}

func fun() {
    defer fmt.Println("defer...")
    // return // 终止函数
    runtime.Goexit()  // 结束当前goroutine
    fmt.Println("fun()....")
}