package main

import "fmt"

func main() {

	/*
	变量转换
		数字在需要时自动转换
		变量需要手动转换, 注意， bool不能转换为int等
	 */
	var a int8
	a = 10

	var b int16
	b = int16(a)
	fmt.Println(a, b)

	f1 := 3.14
	var c int
	c = int(f1)
	fmt.Println(f1, c)

	f1 = float64(a)
	fmt.Println(f1, a)

	//b1 := true
	//a = int8(b1) //cannot convert b1 (type bool) to type int8

	sum := f1 + 100
	fmt.Printf("%T, %f\n", sum, sum)
}
