package main

import (
    "fmt"
    "reflect"
)

func main() {
    // 反射操作： 通过反射， 可以获取一个接口类型变量的 类型和数值

    var x float64 = 3.14
    fmt.Println("type:", reflect.TypeOf(x))
    fmt.Println("value:", reflect.ValueOf(x))

    fmt.Println("--------------------")
    // 根据反射的值， 来获取对应的类型和数值
    v := reflect.ValueOf(x)
    fmt.Println("kind float64:", v.Kind() == reflect.Float64)
    fmt.Println("type:", v.Kind())
    fmt.Println("value:", v.Float())

}
