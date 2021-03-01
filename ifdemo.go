package main

import "fmt"

func main() {

	/**
	// 方式1
	if bool {
		dosth
	}
	// 方式2
	if bool {

	} else {

	}

	// 方式3
	if bool1 {

	} else if bool2 {

	}
	 */

	num := 6
	if num > 10 {
		fmt.Println("大于10")
	}

	if num > 10 {
		fmt.Println("大于10")
	} else if num < 10 {
		fmt.Println("小于10")
	}

	score := 0
	fmt.Println("请输入一个分数")
	//fmt.Scanln(&score)

	if score > 60 {
		fmt.Println("dayu 60")
	} else if score < 60 {
		fmt.Println("xiaoyu 60")
	}

	s1 := "男"
	s2 := "男"

	fmt.Println(".......main over......")
	fmt.Printf("%p, %p\n", &s1, &s2)
	fmt.Println(s1 == s2) // true

}
