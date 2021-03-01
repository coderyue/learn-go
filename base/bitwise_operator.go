package main

import "fmt"

func main() {

	/**
	位运算符
		将数值， 转为二进制后， 按位操作
	按位&：
	按位|：
	异或^：
		二元：a^b
			不同为1， 相同为0
		一元: ^a		// 相当于其他语言的取反 ~
			按位取反
	位清空 &^：
		对于 a &^ b
			对于b上的每个数值
			如果为0， 则取a对应位上的数值
			如果为1， 则结果位就取0
	移位运算符
	<<: 按位左移
		a << b
	>>: 按位右移
		a >> b
	 */

	a := 60
	b := 13
	/*
	a 60 0011 1100
	b 13 0000 1101
	&    0000 1100
	|    0011 1101
	^    0011 0001
	&^   0011 0000
	 */
	fmt.Printf("a:%d, %b\n", a, a)
	fmt.Printf("b:%d, %b\n", b, b)

	ret1 := a & b
	fmt.Println(ret1)

	ret2 := a | b
	fmt.Println(ret2)

	ret3 := a ^ b
	fmt.Println(ret3)

	ret4 := a &^ b
	fmt.Println(ret4)

	ret5 := ^a
	fmt.Println(ret5)
	fmt.Printf("ret5: %b\n", ret5)

	c := 8
	fmt.Println(c << 2)
	fmt.Println(c >> 2)



}
