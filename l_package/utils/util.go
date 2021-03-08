package utils

import "fmt"

func Count() {
	fmt.Println("utils包下的count函数。。。。")
}

//init函数和main定义时不能有任何参数和返回值， 并且由go程序自动调用， 不可以被引用
func init() {
	fmt.Println("utils下的初始化函数....")
}

func init() {
	fmt.Println("utils包下的另一个初始化函数。。。。。")
}
