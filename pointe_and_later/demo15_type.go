package main

import (
	"fmt"
	"strconv"
)

func main() {
	/**
	type: 用于类型定义和类型别名

		1.类型定义: type 类型名 Type
		2.类型别名: type 类型名 = Type

	*/

	var i1 myint
	i1 = 100
	var i2 = 200
	fmt.Printf("%T, %T\n", i1, i2) //main.myint, int
	//i1 = i2  //cannot use i2 (type int) as type myint in assignment

	var s1 mystr
	var s2 = "李小花"
	s1 = "王二狗"
	fmt.Println(s1, s2)
	fmt.Printf("%T, %T, %T, %T\n", i1, i2, s1, s2)
	//s1 = s2 // cannot use s2 (type string) as type mystr in assignment

	fmt.Println("----------------------")

	res1 := fun1()
	fmt.Println(res1(10, 20))

	fmt.Println("--------------------")
	var i3 myint2
	i3 = 1000
	fmt.Println(i3)
	i3 = i2
	fmt.Println(i3)
	fmt.Printf("%T, %T, %T\n", i1, i2, i3) //main.myint, int, int

}

// 定义一个新的类型
type myint int
type mystr string

// 定义函数类型
type myfun func(int, int) string

func fun1() myfun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}

// 类型别名
type myint2 = int // 不是重新定义新的类型， 只是给int起别名， 和int可以通用
