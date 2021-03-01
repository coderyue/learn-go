package main

import "fmt"

func main() {
	/*
		if语句的其他写法
			if 初始化语句; 条件 {

			}
	*/

	// 初始化语句定义的num只能在该if中使用, 在下面的else if中也可以使用这个num
	if num := 3; num > 0 {
		fmt.Println("num 是正数")
	} else if num < 0 {
		fmt.Println("num 是负数")
	}

}
