package main

import "fmt"

func main() {
	// 常数
	fmt.Println(100)
	fmt.Println("test")

	// 定义常量
	const PATH string = "http://www.baidu.com"
	const PI = 3.14
	fmt.Println(PATH)
	//fmt.Println(PI)

	// 常量不能修改
	//PATH = "testset"

	// 定义一组常量
	const C1, C2, C3 = 100, 3.14, "haha"
	fmt.Println(C1, C2, C3)
	const (
		MAIL   = 0
		SEX    = 1
		UNKNOW = 3
	)

	// 在一组常量中， 如果某个常量没有初始值，默认和上一行一致
	const (
		a int    = 100
		b        // 这里没有初始值， 默认和上一行一致， 类型和初始值都一致
		c string = "ruby"
		d
		e
	)
	fmt.Printf("%T, %d\n", a, a)
	fmt.Printf("%T, %d\n", b, b)
	fmt.Printf("%T, %s\n", c, c)
	fmt.Printf("%T, %s\n", d, d)
	fmt.Printf("%T, %s\n", e, e)

	// 枚举类型： 使用常量组作为枚举类型， 理解成  一组相关数值的数据
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)



}
