package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	/**
	输入和输出
		fmt包： 输入， 输出

		输出：
			Print() //打印
			Printf() //格式化打印
			Println() //打印之后换行
		格式化打印占位符
	 		%v 原样输出
			%T 打印类型
			%t bool类型
			%s string类型
			%f 浮点
			%d 10进制整数
			%b 2进制整数
			%o 8进制整数
			%x, %X 16进制
				%x 0-9, a-f
				%X 0-9, A-F
			%c 打印字符
			%p 打印地址
			...
		输入：
			func Scanln(a ...interface{}) (n int, err error)
				Scanln is similar to Scan, but stops scanning at a newline and after the final item there must be a newline or EOF.
			Scanf()
		bufio包

	 */

	/*
	var x int
	var y float64
	fmt.Println("请输入一个整数， 一个浮点数")
	fmt.Scanln(&x, &y) // 读取键盘的输入，通过操作地址，赋值给x和y  阻塞式
	fmt.Printf("x: %d, y: %f\n", x, y)

	// 按格式度读取，这里两个值之间用逗号隔开， 和%d,%f保持一致
	fmt.Scanf("%d,%f", &x, &y)
	fmt.Printf("x: %d, y: %f\n", x, y)
	*/
	fmt.Println("请输入一个字符串")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println(s1)


}