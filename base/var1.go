package main

import "fmt"

func main() {

	// 定义变量
	var num1 int
	num1 = 12
	fmt.Printf("%d\n", num1)

	var num2 int = 13
	fmt.Printf("%d\n", num2)

	var num3 = 15
	fmt.Printf("%d\n", num3)

	// 简短定义
	sum := 23
	fmt.Printf("%d\n", sum)

	fmt.Println("gogogogo..")

	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)

	var m, n int = 12, 13
	fmt.Println(m, n)

	var n1, n2, n3 = 100, 3.14, "Go"
	fmt.Println(n1, n2, n3)

	var (
		studentName = "name"
		age         = 18
		sex         = "女"
	)
	fmt.Printf("学生姓名: %s, 年龄: %d, 性别: %s\n", studentName, age, sex)
	fmt.Println("学生姓名:", studentName, "年龄:", age, "性别:", sex)

}
