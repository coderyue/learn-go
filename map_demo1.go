package main

import "fmt"

func main() {
	/**
	map: 映射，存储键值对，属于引用类型

	存储特点
		A:无序键值对
		B:键不能重复，并且和value一一对应
			若key重复， 则会新值覆盖旧值

	语法结构：
		1.创建map
			var map1 map[key类型]value类型
				nil map, 无法直接使用

			var map2 = make(map[key类型]value类型)

			var map3 = map[key类型]value类型{key:value, key:value, key:value...}

	数据类型的初始值
		int:0
		array: array类型的空值

		slice: nil  但是指向的是数组， 可以直接使用
		map: nil  不能直接使用
	 */

	// 1.创建map
	var map1 map[int]string
	var map2 = make(map[int]string)
	map3 := map[string]int{"go":12, "python":23, "c":25}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

	// 2. nil map  当map是nil时， 不能往里面存储数据
	if map1 == nil {
		map1 = make(map[int]string)
		fmt.Println(map1 == nil)
	}

	// 3.存储键值对到map中
	map1[1] = "hello"
	map1[2] = "world"
	map1[3] = "memeda"
	fmt.Println(map1)

	// 4.获取数据， 根据key获取对应的value值
	// key 不存在， 则返回value对应类型的零值
	fmt.Println(map1[3])
	fmt.Println(map1[30]) // ""

	value, ok := map1[30]
	if ok {
		fmt.Println("get the value: ", value)
	} else {
		fmt.Println("can not get the value: ", value)
	}

	// 5.修改数据
	fmt.Println(map1)
	map1[3] = "如花"
	fmt.Println(map1)

	// 6.删除数据
	delete(map1, 3)
	fmt.Println(map1)
	delete(map1, 30)
	fmt.Println(map1)

	// 7.长度
	fmt.Println(len(map1))

}
