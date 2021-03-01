package main

import "fmt"

var a = 1000 //全局变量

func main() {
	var num int
	num = 100
	fmt.Printf("num是：%d, address: %p\n", num, &num)

	num = 200
	fmt.Printf("num是：%d, address: %p\n", num, &num)

	var name string
	name = "lin"
	fmt.Printf("name is :%s\n", name)

	// 简短定义, 左边至少要有一个是新定义的变量， 非新定义的变量会进行赋值
	num, name, sex := 300, "yue", "femail"
	fmt.Printf("test: %d, %s, %s\n", num, name, sex)

	fmt.Println(a)

	fmt.Println("==================")

	var m int
	fmt.Println(m)
	var n float64
	fmt.Println(n)
	var s string	 // ""
	fmt.Println(s)
	var s2 [] int // 切片
	fmt.Println(s2)
	fmt.Println(s2 == nil)


}
