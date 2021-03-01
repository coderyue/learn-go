package main

import (
	"fmt"
	"strconv"
)

func main() {
	/**
	  strconv包：字符串和基本类型之间的转换
	      string convert

	*/

	// 1.bool类型
	s1 := "true"
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T, %t\n", b1, b1)

	ss1 := strconv.FormatBool(b1)
	fmt.Printf("%T, %s\n", ss1, ss1)

	// 整数
	s2 := "100"
	// 第二个参数那个是转换成多少进制的， 第三个参数是多少bit
	i2, err := strconv.ParseInt(s2, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T, %d\n", i2, i2)

	ss2 := strconv.FormatInt(i2, 10)
	fmt.Printf("%T, %s\n", ss2, ss2)

	// itoa(),  atoi()
	i3, err := strconv.Atoi("-42")
	fmt.Printf("%T, %d\n", i3, i3)
	ss3 := strconv.Itoa(i3)
	fmt.Printf("%T, %s\n", ss3, ss3)

}
