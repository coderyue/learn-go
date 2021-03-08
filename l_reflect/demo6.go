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

func (p Person) Test(i, j int, s string) {
    fmt.Println(i, j, s)
}
func main() {
    /**
    通过反射来进行方法的调用
    1. 接口变量 -> 对象反射对象： Value
    2. 获取对应的方法对象：MethodByName()
    3. 将方法对象进行调用： Call()
     */

    p1 := Person{"ruby", 20, "男"}
    value := reflect.ValueOf(p1)
    fmt.Printf("kind:%s, type: %s\n", value.Kind(), value.Type()) //kind:struct, type: main.Person

    methodValue1 := value.MethodByName("PrintInfo")
    fmt.Printf("kind:%s, type: %s\n", methodValue1.Kind(), methodValue1.Type()) //kind:func, type: func()

    // 没有参数， 直接调用
    methodValue1.Call(nil)

    methodValue2 := value.MethodByName("Say")
    fmt.Printf("kind:%s, type: %s\n", methodValue2.Kind(), methodValue2.Type()) //kind:func, type: func(string)
    args2 := []reflect.Value{reflect.ValueOf("反射机制")}
    methodValue2.Call(args2)

    methodValue3 := value.MethodByName("Test")
    fmt.Printf("kind: %s, type: %s\n", methodValue3.Kind(), methodValue3.Type()) //kind: func, type: func(int, int, string)
    args3 := []reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200), reflect.ValueOf("hello world")}
    methodValue3.Call(args3)
}
