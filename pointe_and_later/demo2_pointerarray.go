package main

import "fmt"

func main() {
	/**
	数组指针: 首先是一个指针，一个数组的地址
		*[4]Type
	指针数组: 首先是一个数组， 存储的数据类型是指针
		[4]*Type
	*/

	// 创建一个数组
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Printf("%p\n", &arr)

	// 创建一个指针， 指向该数组的地址
	var p1 *[4]int
	p1 = &arr
	fmt.Println(p1)         // &[1 2 3 4]  前面的&表示这是一个指针，是数组的指针
	fmt.Printf("%p\n", p1)  // 0xc0000121a0
	fmt.Printf("%p\n", &p1) // 0xc000006030

	// 根据数组指针， 操作数组
	(*p1)[0] = 100
	fmt.Println(arr)

	// 简化写法, 在c中数组变量就是该数组的地址， 指针指向这个数组变量， 那该指针就相当于该数组，简化写法和c一致
	p1[0] = 200
	fmt.Println(arr)

	// 4.指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr2)
	fmt.Println(arr3)
	arr2[0] = 100
	fmt.Println(arr2)
	fmt.Println(a)

	*arr3[0] = 200
	fmt.Println(arr3)
	fmt.Println(a)

	b = 1000
	fmt.Println(arr2)
	fmt.Println(arr3)

	for _, v := range arr3 {
		fmt.Printf("%d\t", *v)
	}
	fmt.Println()
	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i])
	}

}
