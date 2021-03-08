package main

import (
    "fmt"
    "time"
)

func main() {
    /**
    1. func NewTimer(d Duration) *Timer
        创建一个计时器， d时间以后出发
     */

    //timer := time.NewTimer(3 * time.Second)
    //fmt.Printf("%T\n", timer) //*time.Timer
    //fmt.Println(time.Now()) //2021-03-06 22:43:19.7104245 +0800 CST m=+0.003122801
    //
    //// 此处等待channel中数值， 会阻塞三秒
    //ch := timer.C
    //fmt.Println(<- ch) //2021-03-06 22:43:22.7166687 +0800 CST m=+3.009367001

    // 这个计时器可以停止
    timer := time.NewTimer(5 * time.Second)
    go func() {
        <- timer.C
        fmt.Println("Timer 结束了。。。。")
    }()

    time.Sleep(3 * time.Second)
    flag := timer.Stop()
    if flag {
        fmt.Println("Timer 停止了")
    }


}
