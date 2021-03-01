package main

import "fmt"

func main() {
	/**
	slice := arr[start:end]
		切片中的数据: [start,end)
		arr[:end], 从头到end
		arr[start:] 从start到末尾

	从已有数组上， 直接创建切片, 该切片的底层数组就是当前的数组， (存储的地址就是当前数组的地址)
		长度是从start到end切割的数据量
		但是容量是从start到数组的末尾
	切片从数组下标0取值创建时， 切片中的地址和原数组一致， 否者，不一致
	*/

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("--------------1.已有数组直接创建切片----------------")
	s1 := a[:5]  //[1 2 3 4 5]
	s2 := a[3:8] //[4 5 6 7 8]
	s3 := a[5:]  //[6 7 8 9 10]
	s4 := a[:]   //[1 2 3 4 5 6 7 8 9 10]
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	// 由此可知， 切片从数组下标0取值创建时， 切片中的地址和原数组一致， 否者，不一致
	fmt.Printf("a  address: %p\n", &a) //0xc000018140
	fmt.Printf("s1 address: %p\n", s1) //0xc000018140
	fmt.Printf("s2 address: %p\n", s2) //0xc000018158
	fmt.Printf("s3 address: %p\n", s3) //0xc000018168
	fmt.Printf("s4 address: %p\n", s4) //0xc000018140

	fmt.Println("--------------2.长度和容量----------------")
	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Printf("s2 len: %d, cap: %d\n", len(s2), cap(s2))
	fmt.Printf("s3 len: %d, cap: %d\n", len(s3), cap(s3))
	fmt.Printf("s4 len: %d, cap: %d\n", len(s4), cap(s4))

	fmt.Println("--------------3.更改数组的内容----------------")
	a[4] = 100
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
	s2[2] = 200
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	fmt.Println("--------------4.往切片中添加内容----------------")
	s1 = append(s1, 1, 1, 1, 1)
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
	fmt.Printf("a  address: %p\n", &a)
	fmt.Printf("s1 address: %p\n", s1)
	fmt.Printf("address: %p\n", s2)

	// 当往切片中添加元素超出原来数组最后一个地址长度时， 重新创建一个数组， 切片指向新创建的数组
	s1 = append(s1, 2, 2, 2, 2, 2)
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
	fmt.Printf("a  address: %p\n", &a)
	fmt.Printf("s1 address: %p\n", s1)
	s2 = append(s2, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1)
	fmt.Printf("address: %p\n", s2)

}
