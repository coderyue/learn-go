package main

import "fmt"

func main() {

	/**
	指针作为参数：

	参数的传递：值传递， 引用传递

	*/

	a := 10
	fmt.Println("函数调用前：a", a)
	func1(a)
	fmt.Println("函数调用后：a", a)

	func2(&a)
	fmt.Println("函数修改后，a：", a)

	arr := [4]int{1, 2, 3, 4}
	func3(arr)
	fmt.Println("函数调用后arr:", arr)

	func4(&arr)
	fmt.Println("函数调用后arr：", arr)

	s1 := []int{1, 2, 3, 4} // 切片本身就是引用类型的数据，传递的就是地址
	fmt.Println(s1)

}

func func4(p1 *[4]int) {
	fmt.Println("func4函数中的arr:", *p1)
	p1[0] = 100
	fmt.Println("func4修改后 arr：", *p1)
}

func func3(arr [4]int) {
	fmt.Println("func3函数中数组arr：", arr)
	arr[0] = 100
	fmt.Println("func3函数中修改后arr：", arr)
}

func func2(p1 *int) { // 引用传递
	fmt.Println("funf2函数中：num的值：", *p1)
	*p1 = 200
	fmt.Println("func2中修改num：", *p1)
}

func func1(num int) { // 值传递
	fmt.Println("func1函数中num：", num)
	num = 100
	fmt.Println("func1中num修改后：", num)
}
