package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string
    Age int
    Sex string
}

func (p Person) Say(msg string) {
    fmt.Println("say:", msg)
}

func (p Person) PrintInfo()  {
    fmt.Println(p.Name, p.Age, p.Sex)
}

func main() {
    /**
    结构体
     */
    p1 := Person{"王二狗", 18, "男"}
    GetMessage(p1)
}

// 获取空接口的信息
func GetMessage(input interface{}) {
    getType := reflect.TypeOf(input)
    fmt.Println("get type is:", getType.Name()) //get type is: Person
    fmt.Println("get kind is:", getType.Kind()) //get kind is: struct

    getValue := reflect.ValueOf(input)
    fmt.Println("get all field is:", getValue)

    // 获取字段
    /*
    1. 先获取Type对象，refect.Type
        NumField()
        Field(index)
    2. 通过Field()获取每一个Field字段
    3. Interface(), 得到对应的value
     */
    for i := 0; i < getType.NumField(); i++ {
        field := getType.Field(i)
        value := getValue.Field(i).Interface()
        fmt.Printf("字段名称：%s, 字段类型：%s, 字段的值：%v\n", field.Name, field.Type, value)
    }

    // 获取方法
    for i := 0; i < getType.NumMethod(); i++ {
        method := getType.Method(i)
        fmt.Printf("方法名称: %s, 方法类型： %v\n", method.Name, method.Type)
    }

}