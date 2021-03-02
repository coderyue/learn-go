package main

import "fmt"

func main() {
	/**
	结构体：是由一系列具有相同类型或不同类型的数据构成的数据集合
		结构体成员是由一系列的成员变量构成，这些成员变量被称为“字段”

	*/
	// 创建一个结构体对象
	// 方法1
	var p1 Person
	fmt.Println(p1)
	p1.name = "王二狗"
	p1.age = 12
	p1.sex = "男"
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n", p1.name, p1.age, p1.sex)

	// 方法2
	p2 := Person{}
	fmt.Println(p2)
	p2.name = "ruby"
	p2.age = 23
	p2.sex = "女"
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n", p1.name, p1.age, p1.sex)

	// 方法3
	p3 := Person{name: "李小花", age: 25, sex: "女"}
	fmt.Println(p3)
	p4 := Person{
		name: "ruby",
		age:  14,
		sex:  "男",
	}
	fmt.Println(p4)

	// 方法4
	p5 := Person{"如花", 25, "女"}
	fmt.Println(p5)
}

type Person struct {
	name string
	age  int
	sex  string
}
