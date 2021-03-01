package main

import "fmt"

func main() {

	/**
	算术运算符 : + - * / % ++ --

	++: 给自己加1
		i++
	--: 给自己减1
		i--

	c中不支持  --i等复杂写法, 只有简单的i++ i--
	 */

	a := 10
	b := 3
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	div := a / b
	mod := a % b
	fmt.Printf("%d / %d = %d\n", a, b, div)
	// 这里两个%表示输出%
	fmt.Printf("%d %% %d = %d\n", a, b, mod)

	c := 3
	c++
	fmt.Println(c)

}
