package main

import "fmt"

func main() {
	/**
	匿名结构体和匿名字段

	匿名结构体: 没有名字的结构体
		在创建匿名结构体时，同事创建对象
		变量名 ：= struct {
			定义字段field
		}{
			字段进行赋值
		}

	匿名字段： 一个结构体的字段没有字段名

	匿名函数: 没有名字的函数


	结构体嵌套：一个结构体中的字段，是另一个结构体类型
		has a
		嵌套的时候 使用指针， 更改才能修改另一个值， 否则是另外开辟一块空间

	*/

	s1 := Student{name: "张三", age: 18}
	fmt.Println(s1.name, s1.age)

	func() {
		fmt.Println("匿名函数。。。。")
	}()

	// 匿名结构体
	s2 := struct {
		name string
		age  int
	}{
		name: "李四",
		age:  19,
	}
	fmt.Println(s2)

	//w1 := Worker{name: "王二狗", age: 18}
	//fmt.Println(w1)

	w2 := Worker{"李小花", 19}
	fmt.Println(w2)
	fmt.Println(w2.string)
	fmt.Println(w2.int)

}

type Worker struct {
	//name string
	//age int
	string // 匿名字段
	int    // 匿名字段 默认使用数据类型作为字段，可以通过类型进行访问，，-> 匿名字段类型不能重复
}

type Student struct {
	name string
	age  int
}
