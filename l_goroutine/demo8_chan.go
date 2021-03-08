package main

import "fmt"

func main() {
    /**
    channel, 通道

     */

    // 表示a这个通道能传递int
    var a chan int
    fmt.Printf("%T, %v\n", a, a) //chan int, <nil>   仅仅声明， 没有创建

    if a == nil {
        fmt.Println("channel是nil的，不能使用， 需要先创建通道...")
        a = make(chan int)
        fmt.Println(a)
    }
    test1(a)
}

// 通道传递的是地址
func test1(ch chan int)  {
    fmt.Printf("%p, %v\n", ch, ch)
}
