package main

import "fmt"

func main() {
	/**
	函数指针：一个指针，指向了一个函数的指针
		因为go语言中， function，默认看做一个指针， 没有*

	指针函数：一个函数， 该函数的返回值是一个指针
	*/

	var a func()
	a = func1
	a()
	fmt.Println(a)

	arr1 := func2()
	fmt.Printf("arr1的类型：%T, 地址： %p, 数值： %v\n", arr1, &arr1, arr1)

	arr2 := func3()
	fmt.Printf("arr2的类型：%T, 地址: %p, 数值: %v\n", arr2, &arr2, arr2)
	fmt.Printf("arr2中存储的是数组的地址：%p\n", arr2)

}

func func3() *[4]int {
	arr := [4]int{1, 2, 3, 4}
	fmt.Printf("arr的地址：%p\n", &arr)
	return &arr
}

func func2() [4]int { //普通函数
	arr := [4]int{1, 2, 3, 4}
	return arr
}

func func1() {
	fmt.Println("func1()...")
}
