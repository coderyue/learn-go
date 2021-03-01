package main

import "fmt"

func main() {
	/**
	赋值运算符
		= += -= %= <<= >>= &=
	 */
	// a 3 0101
	a := 3
	fmt.Println(a)
	// 取反运算， 取反后的值等于取反前的值加负号再减1
	// 原理 ，  取反前和取反后的值加起来全都是1， 此时再加1 则为0， 那么一个数的取反就等于取负再减1
	fmt.Println(^a)

	a += 2
	fmt.Println(a)

	a <<= 2
	fmt.Println(a)

	a >>= 2
	fmt.Println(a)

	a &= 2
	fmt.Println(a)


}
