package main

import (
    "fmt"
    "time"
)

// 全局变量，表示票
var ticket = 200
func main() {
    /**
    火车站卖票
        不要以共享内存的方式去通信， 而要以通信的方式去共享内存  （不要用锁， 用channel）
     */

    go selTicket("售票口1")
    go selTicket("售票口2")
    go selTicket("售票口3")

    time.Sleep(1 * time.Second)

}

func selTicket(name string) {
    for {
        if ticket > 0 {
            fmt.Println(name, "售出：", ticket)
            ticket--
        } else {
            fmt.Println("售罄, 没有票了")
            break
        }
    }
}
