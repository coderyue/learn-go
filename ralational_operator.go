package main

import "fmt"

func main() {

	/**
	关系运算符 > < >= <= == !=
		结果都是bool类型  true， false
		== 表示比较两个数值是相等的
		!= 表示比较两个数值是不相等的
	 */

	a := 3
	b := 5
	c := 3

	ret1 := a > b
	ret2 := b > c
	fmt.Printf("%T, %t\n", ret1, ret1)
	fmt.Printf("%T, %t\n", ret2, ret2)

}
