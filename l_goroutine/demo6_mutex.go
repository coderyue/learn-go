package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// 全局变量，表示票
var ticket = 10
var mutex sync.Mutex // 创建锁头
var wg sync.WaitGroup // 同步等待组对象
func main() {
    /**
      火车站卖票
          不要以共享内存的方式去通信， 而要以通信的方式去共享内存  （不要用锁， 用channel）

    在使用互斥锁的时候，对资源操作完，一定要解锁。 否则会出现程序异常，死锁等问题
    defer 语句
    */
    wg.Add(4)

    go selTicket("售票口1")
    go selTicket("售票口2")
    go selTicket("售票口3")
    go selTicket("售票口4")

    wg.Wait()
    //time.Sleep(5 * time.Second)
    fmt.Println("程序结束了。。。。")

}

func selTicket(name string) {
    defer wg.Done()
    rand.Seed(time.Now().UnixNano())
    for {
        // 上锁
        mutex.Lock()
        if ticket > 0 {
            time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
            fmt.Println(name, "售出：", ticket)
            ticket--
        } else {
            // 条件不满足也要解锁
            mutex.Unlock()
            fmt.Println("售罄, 没有票了")
            break
        }
        mutex.Unlock()
    }
}
