package main

import "fmt"

func main() {

	/**
	for循环: 某些代码会被多次执行

	1. 标准写法
		for init; condition; post { }
	2.同时省略表达式1和表达式3,   分号和两个表达式都省略，相当于while(条件)
		for condition {  }
	3.同时省略三个表达式
		for ;; { }   分号省略掉   for { }
		相当于while(true)
	4.其他写法 ：for循环中同时省略几个表达式都可以
		省略表达式二: 死循环
	*/
	// 标准写法
	fmt.Println("标准写法============")
	for i := 0; i < 3; i++ {
		fmt.Printf("i: %d\n", i)
	}

	// 同时省略表达式1和表达式3
	fmt.Println("同时省略表达式1和表达式3============")
	i := 1
	// go里面支持吧分号也省略
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println("->", i)
	for i < 7 {
		fmt.Println(i)
		i++
	}

	//同时省略三个表达式
	fmt.Println("同时省略三个表达式============")
	i2 := 2
	// 这种写法不如直接把条件放到for里, 除非写死循环， 不跳出循环
	for {
		if i2 < 4 {
			fmt.Println("--->", i2)
		} else {
			break
		}
		i2++
	}

}
