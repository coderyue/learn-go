package main

import "fmt"

func main() {

	/**
	切片Slice
		1.每一个切片引用了一个底层数组
		2.切片本身不存储任何数据， 都是底层数组存储，所以修改切片也就是这个数组中的数据
		3.向切片中添加数据时，没有超过容量，直接添加，否则自动扩容(成倍增长)
	*/

	s := []int{1, 2, 3}
	fmt.Println(s)
	fmt.Printf("len:%d, cap:%d\n", len(s), cap(s))
	fmt.Printf("address: %p\n", s)

	s = append(s, 4, 5)
	fmt.Println(s)
	fmt.Printf("len:%d, cap:%d\n", len(s), cap(s))
	fmt.Printf("address: %p\n", s)

	s = append(s, 6, 7, 8)
	fmt.Println(s)
	fmt.Printf("len:%d, cap:%d\n", len(s), cap(s))
	fmt.Printf("address: %p\n", s)

	s = append(s, 9, 10)
	fmt.Println(s)
	fmt.Printf("len:%d, cap:%d\n", len(s), cap(s))
	fmt.Printf("address: %p\n", s)

	s = append(s, 11, 12, 13)
	fmt.Println(s)
	fmt.Printf("len:%d, cap:%d\n", len(s), cap(s))
	fmt.Printf("address: %p\n", s)

}
