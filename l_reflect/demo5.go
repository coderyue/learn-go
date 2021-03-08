package main

import (
    "fmt"
    "reflect"
)

type Student struct {
    Name string
    Age int
    School string
}
func main() {
    /**
    修改结构体中的值
     */
    s1 := Student{"王二狗", 19 , "家里蹲"}
    // 通过反射， 更改对象中的值， 前提是数据可以被更改
    fmt.Println(s1)

    // 改变数值
    value := reflect.ValueOf(&s1)
    if value.Kind() == reflect.Ptr {
        newValue := value.Elem()
        fmt.Println(newValue.CanSet())

        f1 := newValue.FieldByName("Name")
        f1.SetString("lin")
        f2 := newValue.FieldByName("School")
        f2.SetString("母鸡")
        fmt.Println(s1)

    }

}
