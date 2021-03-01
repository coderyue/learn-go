package main

import (
	"fmt"
	"sort"
)

func main() {
	/**
	map 的遍历
		使用： for range

			数组，切片: index, value
			map: key, value

		注意：Java中map中数据不发生改变，每次遍历顺序一样，但是这里结果会发生改变，多执行几次就可以看到

	 */
	map1 := make(map[int]string)
	map1[1] = "红孩儿"
	map1[2] = "金角大王"
	map1[3] = "小钻风"
	map1[4] = "白骨精"

	// 1. 遍历map
	for k, v := range map1 {
		fmt.Println(k, "-->", v)
	}

	// 按顺序遍历， 获取所有的key，再遍历
	keys := make([]int, 0, len(map1))
	fmt.Println(keys)

	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	sort.Ints(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, map1[key])
	}

	s := []string{"apple", "android", "window", "王二狗"}
	sort.Strings(s)
	fmt.Println(s)


}
