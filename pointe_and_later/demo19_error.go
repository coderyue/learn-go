package main

import (
	"errors"
	"fmt"
)

func main() {
	/**
	error: 内置的数据类型， 内置的接口
		定义方法：Error() string

	使用go语言提供好的包
		errors包下的函数： New(),创建一个error对象
		fmt包下的Errorf()函数
			func Errorf(format string, a ...interface{}) error
	*/

	// 创建一个error数据
	err1 := errors.New("自己创建玩的..")
	fmt.Println(err1)
	fmt.Printf("%T\n", err1)

	// 另一个创建error的方法
	err2 := fmt.Errorf("错误信息码： %d\n", 100)
	fmt.Println(err2)
	fmt.Printf("%T\n", err2)

	fmt.Println("-------------------")
	err3 := checkAge(-30)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println("程序。。go。。 on。。。")

}

func checkAge(age int) error {
	if age < 0 {
		// 返回error对象
		//return errors.New("年龄不合法")
		return fmt.Errorf("您给定的年龄是：%d, 不合法", age)
	}
	fmt.Println("年龄是：", age)
	return nil
}
