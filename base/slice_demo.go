package main

import "fmt"

func main() {
	/**
	数组array:
		存储一组相同数据类型的数据结构
			特点：定长

	切片 slice:
		同数组类似， 也叫做变长数组或者动态数组
			特点：变长

		是一个引用类型的容器，指向了一个底层的数组

	make()
	func make(t Type, size ...IntegerType) Type

	append()  // 专门用于向切片的尾部追加元素
		slice = append(slice, elem1, elem2)
		slice = append(slice, anotherSlice...)

	第一个参数： 类型
		slice， map，channel
	第二个参数： 长度len
		实际存储元素的数量
	第三个参数：容量cap
		最多能够存储的元素的数量
	*/

	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)

	var s1 []int // 不给出长度就是切片
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4}
	fmt.Println(s2)

	fmt.Printf("%T, %T\n", arr, s1) //[4]int, []int

	// ---------------使用make创建-------------------

	s3 := make([]int, 3, 10) // 默认长度是0
	fmt.Println(s3)
	fmt.Printf("length: %d, cap: %d\n", len(s3), cap(s3))
	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	fmt.Println(s3)
	//fmt.Println(s3[3]) // panic: runtime error: index out of range [3] with length 3

	// append()
	//slice = append(slice, elem1, elem2)
	//slice = append(slice, anotherSlice...) 三个点表示把后面切片中的元素拿出来放到第一个切片中
	fmt.Println("--------------")
	s4 := make([]int, 3, 5)
	fmt.Println(s4)
	s4 = append(s4, 4, 5) // 扩容后改变了地址，所以这里要重新赋值一下
	fmt.Println(s4)
	s4 = append(s4, s3...)
	fmt.Println(s4)

	fmt.Println("-----------------")
	// 数组的一些操作对切片也是适用的
	// 遍历切片
	for i := 0; i < len(s4); i++ {
		fmt.Print(s4[i], "\t")
	}
	fmt.Println()
	for i, val := range s4 {
		fmt.Printf("%d-->%d\n", i, val)
	}

}
