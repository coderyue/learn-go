package main

import (
	"fmt"
	"math"
)

func main() {
	/**
	for 的练习

	for中循环控制语句
		break  		强制结束循环(停止循环)
		continue	强制结束本次循环， 执行下一次循环
		注意， 多层循环嵌套时， 对离自己最近的生效
	 */

	// 打印1-100内，能被3整除但是不能被5整除的所有整数，输出共个数，每行打印5个
	count := 0
	for i := 0; i < 100; i++ {
		if i % 3 ==0 && i % 5 != 0 {
			fmt.Print(i, "\t")
			count++
			if count % 5 == 0 {
				fmt.Println()
			}
		}

	}
	fmt.Println()
	fmt.Println("->", count)

	// 打印99乘法表
	out:
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			//fmt.Print(i, "*", j, "=", i * j, "\t")
			fmt.Print(j, "*", i, "=", i * j, "\t")
			if j == 3 {
				//break out
				fmt.Println()
				continue out   // 结束循环到标签处， 但是外部循环依然存在
			}
		}
		fmt.Println()
	}

	// 求水仙花数
	for i := 1; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				num := i * 100 + j * 10 + k
				//if num == i * i * i + j * j * j + k * k * k {
				if float64(num) == math.Pow(float64(i), 3) + math.Pow(float64(j), 3) + math.Pow(float64(k), 3) {
					fmt.Println(num)
				}
			}
		}
	}

	// 求 素数
	for i := 2; i <= 100; i++ {
		flag := true
		//for j := 2; j < i; j++ { // 判断到根号i就可以
		for j := 2; j < int(math.Sqrt(float64(i))); j++ {
			if i % j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}


}
