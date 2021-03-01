package main

import "fmt"

func main() {
	/**
	go中的字符串是一个字节的切片
		可以通过将其内容封装在“”中来创建字符串。go中的字符串是Unicode兼容的，并且是utf-8编码的

	字符串是一些字节的集合

	语法: "", ``
		""
		"a", "中"
		"hello"

	字符：--> 对应编码表中的编码值
		A -->65
		B -->66
		a -->97

	字节： byte-->uint8

	字符串遍历的时候， 如果有中文，使用len遍历的时候中文部分会出现乱码， 但是使用for range不会

	*/

	// 定义字符串
	s1 := "hello中国"
	s2 := `hello world`
	fmt.Println(s1)
	fmt.Println(s2)

	// 2. 字符串的长度： 返回的是字节的个数
	fmt.Println(len(s1))
	fmt.Println(len(s2))

	// 3. 获取某个字节
	fmt.Println(s2[0]) // 获取字符串中的第一个字节
	a := 'h'
	b := 104
	fmt.Printf("%c, %c, %c\n", s2[0], a, b)

	// 字符串的遍历
	for i := 0; i < len(s2); i++ {
		//fmt.Printf("%c\t", s2[i])
		fmt.Printf("%c\t", s1[i])
	}
	fmt.Println()

	// for range
	for _, v := range s1 {
		fmt.Printf("%c", v)
	}
	fmt.Println()
	for _, v := range s2 {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	// 5. 字符串是字节的集合
	slice1 := []byte{65, 66, 67, 68}
	s3 := string(slice1) // 根据一个字节切片， 构建字符串
	fmt.Println(s3)

	s4 := "abcdef"
	slice2 := []byte(s4)
	fmt.Println(slice2)

	// 6. 字符串o不能修改
	//s4[2] = 'B' //cannot assign to s4[2] (strings are immutable)

}
