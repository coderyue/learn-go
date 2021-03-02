package main

import "fmt"

func main() {

	/**
	指针：pointe_and_later
		存储了另一个变量的内存地址的变量
	*/

	a := 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Printf("a address: %p\n", &a)

	var p1 *int
	fmt.Println(p1) // 空指针
	p1 = &a
	fmt.Println("p1的数值", p1)
	fmt.Printf("p1 address: %p\n", &p1)
	fmt.Println("p1中存储的地址中存储的数据", *p1)

	// 操作变量, 更改数值， 不会更改变量的地址
	a = 100
	fmt.Println(a)
	fmt.Printf("%p\n", &a)

	// 通过指针更改变量的值
	*p1 = 200
	fmt.Println(a)
	fmt.Printf("%p\n", &a)
	fmt.Println(p1)

	// 指针的指针
	var p2 **int
	fmt.Println(p2)
	p2 = &p1
	fmt.Println("p2的数值", p2)
	fmt.Printf("%p\n", &p2)
	fmt.Println("p2中存储的地址对应的数据，就是p1的地址，对应的数据:", *p2)
	// 那么可以多个*获取a中存储的数据
	fmt.Println("使用p2获取a中存储的数据：", **p2)

}
