package main

import "fmt"

func main() {
	/**
	数组排序


	*/

	// 冒泡排序
	arr1 := [5]int{2, 54, 1, 45, 76}
	// 第一轮排序
	//for i := 0; i < len(arr1) - 1; i++ {
	//	if arr1[i] > arr1[i + 1] {
	//		arr1[i], arr1[i + 1] = arr1[i + 1], arr1[i]
	//	}
	//}
	//fmt.Println(arr1)
	//
	//// 第二轮排序
	//for i := 0; i < 3; i++ {
	//	if arr1[i] > arr1[i + 1] {
	//		arr1[i], arr1[i + 1] = arr1[i + 1], arr1[i]
	//	}
	//}
	//fmt.Println(arr1)

	// 最后总结, 排序最后用的是下一个循环，
	//for i := 0; i < len(arr1) - 1; i++ {
	//	for j := 0; j < i; j++ {
	//		if arr1[j] > arr1[j+1] {
	//			arr1[j], arr1[j+1] = arr1[j+1], arr1[j]
	//		}
	//	}
	//}
	for i := 1; i < len(arr1); i++ {
		for j := 0; j < len(arr1) - i; j++ {
			if arr1[j] > arr1[j+1] {
				arr1[j], arr1[j+1] = arr1[j+1], arr1[j]
			}
		}
		fmt.Println(arr1)
	}
	//fmt.Println(arr1)

}
