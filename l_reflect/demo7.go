package main

import (
    "fmt"
    "reflect"
    "strconv"
)

func main() {
    /**
    函数的反射
    思路： 函数也是看做接口变量类型
    1. 函数 -> 反射对象， Value
    2. kind -> func
    3. Call()
     */
    f1 := fun1
    value1 := reflect.ValueOf(f1)
    fmt.Printf("kind: %s, type: %s\n", value1.Kind(), value1.Type()) //kind: func, type: func()

    value2 := reflect.ValueOf(fun2)
    value3 := reflect.ValueOf(fun3)
    fmt.Printf("kind: %s, type: %s\n", value2.Kind(), value2.Type()) //kind: func, type: func(int, string)
    fmt.Printf("kind: %s, type: %s\n", value3.Kind(), value3.Type()) //kind: func, type: func(int, string) string

    // 通过反射调用函数
    value1.Call(nil)
    value2.Call([]reflect.Value{reflect.ValueOf(100), reflect.ValueOf("lin")})
    resValue := value3.Call([]reflect.Value{reflect.ValueOf(100), reflect.ValueOf("test")})
    fmt.Printf("%T\n", resValue) // []reflect.Value
    fmt.Println(len(resValue)) //1
    fmt.Printf("kind: %s, type: %s\n", resValue[0].Kind(), resValue[0].Type()) //kind: string, type: string

    s := resValue[0].Interface().(string)
    fmt.Println(s)
    fmt.Printf("%T\n", s)
}

func fun1() {
    fmt.Println("我是fun1， 无参的。。。")
}

func fun2(i int, s string) {
    fmt.Println("我是fun2， 有参的...", i, s)
}

func fun3(i int, s string) string {
    fmt.Println("我是fun3， 有参的，，也有返回值。。。")
    return s + strconv.Itoa(i)
}