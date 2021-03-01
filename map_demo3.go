package main

import "fmt"

func main() {

	/**
	1.和slice结合使用

	2.map是引用类型

		make() 创建都是引用类型

	 */

	map1 := make(map[string]string)
	map1["name"] = "熊大"
	map1["age"] = "12"
	map1["sex"] = "男"
	fmt.Println(map1)

	map2 := make(map[string]string)
	map2["name"] = "小翠"
	map2["age"] = "23"
	map2["sex"] = "女"
	fmt.Println(map2)

	s := make([]map[string]string, 0, 2)
	s = append(s, map1)
	s = append(s, map2)
	fmt.Println(s)

	// 遍历切片
	for i, val := range s {
		fmt.Printf("第%d个人的信息是：\n", i + 1)
		fmt.Println("\tname:", val["name"])
		fmt.Println("\tage:", val["age"])
		fmt.Println("\tsex", val["sex"])
	}

	// map是引用类型
	fmt.Println("--------------")
	map3 := make(map[string]map[string]string)
	fmt.Printf("%T\n", map1)
	fmt.Printf("%T\n", map3)
	map3["熊大"] = map1
	map3["熊二"] = map2
	fmt.Println(map3)

}
