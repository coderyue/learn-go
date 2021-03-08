package main

import (
    "fmt"
    "reflect"
)

func main() {
    /*
    设置值
     */

    var num float64 = 1.23
    fmt.Println("num的数值是:", num)

    // 需要操作指针
    // 通过reflect.ValueOf() 获取num的Value对象
    pointer := reflect.ValueOf(&num) // 注意参数必须指针才能修改其中的值
    // 获取原始值对应的反射对象，只有原始对象才能修改， 当前反射对象不能修改
    newValue := pointer.Elem() // 获取指针类型的value对象那个
    fmt.Println(pointer.CanSet())

    fmt.Println("类型：", newValue.Type()) //float64
    fmt.Println("是否可以修改数据：", newValue.CanSet()) //true

    // 重新赋值
    newValue.SetFloat(3.14)
    fmt.Println(num)
}
