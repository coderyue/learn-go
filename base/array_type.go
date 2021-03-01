package main

import "fmt"

func main() {

	/**
	数组的数据类型
		[size] type

	值类型: 理解为存储的数据本身 (地址处存储的就是值)
		将数据传递给其他的变量, 传递的是数据的备份(备份)
			int float string bool array
	引用类型: 存储的是地址
		传递的也是地址
			slice map...

	 */
	num := 10
	fmt.Printf("%T\n", num)

	//arr1 := [3] int {}
	var arr1 [3] int
	fmt.Printf("%T\n", arr1) // [3]int

	// 赋值
	num2 := num		// 值传递
	fmt.Println(num, num2)
	num2 = 20
	fmt.Println(num, num2)

	arr2 := arr1	// 数组也是值传递
	fmt.Println(arr1, arr2)
	arr2[1] = 100
	fmt.Println(arr1, arr2)



}
