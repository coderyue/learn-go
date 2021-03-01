package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/**
	生成随机数
	 	伪随机数， 根据一定的算法公式算出来的
	math/rand
	*/

	num1 := rand.Int()
	fmt.Println(num1) // 由于种子相同，所以多次运行结果一样

	//num1 = rand.Intn(10)
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}

	rand.Seed(3)
	fmt.Println("-->", rand.Intn(10))

	t1 := time.Now()
	fmt.Println(t1)

	// 时间戳：指定时间，距离1970年1月1日0点0分0秒， 之间的时间差值： 秒 纳秒
	timestamp1 := t1.Unix() // 秒
	fmt.Println(timestamp1)

	timestamp2 := t1.UnixNano()
	fmt.Println(timestamp2)

	// 所以生成随机数， 可以设置种子为时间戳
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		// 调用生成随机数的函数
		fmt.Println("-->", rand.Intn(10))
	}

	// 获取 1-100的随机数
	fmt.Println(rand.Intn(100) + 1)

}
