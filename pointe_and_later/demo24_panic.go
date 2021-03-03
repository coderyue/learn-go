package main

import "fmt"

func main() {
	/**
	panic: 词意“恐慌”
	recover：“恢复”
	go 语言中利用panic(), recover(),实现程序中的极特殊的异常处理
		panic(), 让当前的程序进入恐慌，中断程序的执行
		recover(), 让程序恢复， 必须在defer函数中执行

	从发生恐慌处， 后面的代码都不会执行，
	*/
	defer func() {
		//recover() 的返回值就恐慌中的值，不需要参
		if msg := recover(); msg != nil {
			fmt.Println(msg, "程序恢复啦。。。")
		}
	}()
	funcA()
	defer myprint("defer main...3...")
	funcB()
	defer myprint("defer main.....4...")

	fmt.Println("main over .......")
}
func myprint(s string) {
	fmt.Println(s)
}

func funcA() {
	fmt.Println("我是一个函	数funcA。。。。")
}

func funcB() { // 外围函数
	//defer func() {
	//	//recover() 的返回值就恐慌中的值，不需要参
	//	if msg := recover(); msg != nil {
	//		fmt.Println(msg, "程序恢复啦。。。")
	//	}
	//}()

	defer myprint("defer funcB。。1。。。")
	for i := 0; i < 10; i++ {
		fmt.Println("i", i)
		if i == 5 {
			panic("funB函数， 恐慌了。。。")
		}
	} //当外围函数的代码中发生了运行恐慌，只有其中所有的已经defer的函数全部都执行完毕后，该运行恐慌才会真正扩展至调用处
	defer myprint("defer funcB...2..")
}
