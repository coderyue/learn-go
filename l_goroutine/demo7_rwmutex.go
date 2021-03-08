package main

import (
    "fmt"
    "sync"
    "time"
)

var rwMuetex *sync.RWMutex
var wg *sync.WaitGroup //同步等待组对象
func main() {
    /*
    读写锁
        1. 可以同时读，多个goroutine同时读
        2. 写的时候，啥也不能干。不能读也不能写

     */
    rwMuetex = new(sync.RWMutex)
    wg = new(sync.WaitGroup)

    // 读操作， 加不加锁，都可以读
    //wg.Add(2)
    //
    //go ReadData(1)
    //go ReadData(2)
    //wg.Wait()

    wg.Add(3)
    go WriteData(1)
    go ReadData(2)
    go WriteData(3)
    wg.Wait()

    fmt.Println("main over ......")

}

func WriteData(i int) {
    defer wg.Done()
    fmt.Println(i, "开始写。。。。")
    rwMuetex.Lock() // 写操作上锁
    fmt.Println(i, "正在写。。。。")
    time.Sleep(3 * time.Second)
    rwMuetex.Unlock()
    fmt.Println(i, "写结束.....")
}

func ReadData(i int) {
    defer wg.Done()
    fmt.Println("开始读。。。。。", i)
    rwMuetex.RLock() // 读操作上锁
    fmt.Println(i, "正在读取数据.....")
    time.Sleep(3 * time.Second)
    rwMuetex.RUnlock() // 读操作解锁
    fmt.Println("读结束.......", i)
}
