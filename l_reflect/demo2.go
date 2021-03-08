package main

import (
    "fmt"
    "reflect"
)

func main() {
    /**
    //获取变量里存的数值
    convertValue := value.Interface().(float64)
     */

    var num float64 = 1.23
    // “接口类型变量” -> "反射类型对象"
    value := reflect.ValueOf(num)

    //"反射类型对象" -> "接口类型变量"
    convertValue := value.Interface().(float64)
    fmt.Println(convertValue)

    /**
    反射类型对象 -> 接口类型变量，理解为“强制转换”
    golang对类型要求非常严格， 类型一定要完全符合
    一个是float64一个是*float64， 不能混淆， 否则panic
     */
    pointer := reflect.ValueOf(&num)
    convertPointer := pointer.Interface().(*float64) //interface conversion: interface {} is *float64, not float64
    fmt.Println(convertPointer)
}
