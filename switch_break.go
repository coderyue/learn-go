package main

import "fmt"

func main() {

	/*
	switch 中的break和fallthrough语句
	break: 可以使用在switch中， 也可以使用在for循环中
		强制结束case语句， 从而结束switch分支
	fallthrou : 用于穿透switch
		当switch中某个case匹配成功后， 执行该case语句
		如果遇到fallthrough，那么后面紧跟的case无需匹配，直接穿透执行
		fallthrough 需要位于某个case的最后一行
	 */

	n := 2
	switch n {
	case 1:
		fmt.Println("我是熊大")
		fmt.Println("我是熊大")
		fmt.Println("我是熊大")
	case 2:
		fmt.Println("我是熊二")
		fmt.Println("我是熊二")
		break	// 用于强制结束case，switch也跟着结束
		fmt.Println("我是熊二")
	}

	fmt.Println("------------------------")

	m := 2
	switch m {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
		fallthrough	// 若执行这个case，则fallthrough可以再执行下一个case
	default:
		fmt.Println("do default")
		//if true {
		//	fallthrough //fallthrough statement out of place
		//}
	case 3:
		fmt.Println(3)
	}

	fmt.Println("main...over....")


}
