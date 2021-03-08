package main

import (
    "fmt"
    "time"
)

func main() {
    /*
       2. func After(d Duration) <-chan Time
           返回一个通道： chan， 存储的d时间之后的当前时间

        相当于： return NewTimer(d).C
    */

    ch := time.After(3 * time.Second)
    fmt.Printf("%T\n", ch) //<-chan time.Time

    fmt.Println(time.Now()) //2021-03-06 22:53:53.5449462 +0800 CST m=+0.003786701
    time2 := <- ch
    fmt.Println(time2) //2021-03-06 22:53:56.5539003 +0800 CST m=+3.012740801

}
