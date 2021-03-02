package main

import "fmt"

func main() {

	/**
	数据类型
	 	基本类型：整数 浮点 布尔 字符串
		复合类型：array, slice, map, struct, pointe_and_later, function, channel...
	数组：
		var 数组名 [长度] 数据类型
		var 数组名 = [长度] {元素1, 元素2...}
		数组名 := [...] 数据类型 {元素...}
	通过下标访问
		下标，索引  index
		默认从0 到长度减1
		数组名[index]
			赋值
			取值
		不能越界: [0, 长度-1]
	长度和容量: go 语言的内置函数
		len() // 长度
		cap() // 容量
	*/

	// step1: 创建数组
	var arr1 [4]int
	arr1[0] = 0
	arr1[1] = 1
	//arr1[2] = 2
	//arr1[3] = 3
	fmt.Println(arr1[2])
	//fmt.Println(arr1[4])  //invalid array index 4 (out of bounds for 4-element array)

	fmt.Println("数组的长度", len(arr1)) // 容器中实际存储的数量
	fmt.Println("数组的容量", cap(arr1)) // 容器中能够存储的最大数量
	// 因为数组定长， 所以长度和容量相同

	// 数组其他创建方式
	var a [4]int // 同 var a = [4] int
	fmt.Println(a)

	var b = [4]int{1, 3, 4, 6}
	fmt.Println(b)

	var c = [5]int{1, 4, 5}
	fmt.Println(c) // 其余补0

	var d = [5]int{1: 1, 3: 2}
	fmt.Println(d)

	var e = [5]string{"asdfa", "王二狗"}
	fmt.Println(e)

	f := [...]int{1, 2, 3, 4, 5}
	fmt.Println(f)

	g := [...]int{1: 1, 6: 3}
	fmt.Println(g)

}
