package main

import "fmt"

func main() {
	/**
	字符串
	1.概念：多个byte的集合， 理解为一个字符序列
	2.语法： 使用双引号
		“abc”
		也可以使用``
	 */

	// 定义字符串
	var s1 string
	s1 = "王二狗"
	fmt.Printf("%T, %s\n", s1, s1)

	var s2 = `hello world`
	fmt.Printf("%T, %s\n", s2, s2)

	// 'A' "A"
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T, %d\n", v1, v1)
	fmt.Printf("%T, %s\n", v2, v2)

	v3 := '中'
	fmt.Printf("%T, %d, %c, %q\n", v3, v3, v3, v3)

	// 转义字符    \n \t 特殊作用
	fmt.Println("\"helloword\"")

	fmt.Println("hello`wor`ld")
	fmt.Println(`hell"ow"orld`)
}
