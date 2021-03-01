package main

import "fmt"

func main() {
	/**
	深拷贝， 浅拷贝

	func copy(dst, src []Type) int
	可以实现切片的拷贝
	*/

	s1 := [4]int{1, 2, 3, 4}
	s2 := make([]int, 0) // len:0 cap:0
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1)
	fmt.Println(s2)

	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)

	// copy()
	s3 := []int{7, 8, 9}
	//copy(s2, s3) // 将s3中的元素， 拷贝到s2中
	//copy(s3, s2) // 将s3中的元素， 拷贝到s2中
	copy(s2[1:3], s3[1:])
	fmt.Println(s2)
	fmt.Println(s3)

}
