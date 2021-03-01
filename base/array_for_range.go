package main

import "fmt"

func main() {
	/*
	数组的遍历


	 */

	arr1 := [5] int {1,2,3,4,5}
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2 + 1
		fmt.Println(arr1[i])
	}
	fmt.Println(arr1)

	fmt.Println("---------------")
	for index, value := range arr1 {
		fmt.Println("index:", index, "value: ", value)
	}

	sum := 0
	for _, value := range arr1 {
		sum += value
	}
	fmt.Println(sum)

}
