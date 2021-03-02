package main

func main() {
	//结构体的匿名字段

	/**
	  面向对象: OOP

	  go语言的结构体嵌套;
		1.模拟继承性： is -> a
			type A struct {
				field
			}
			type B struct {
				A // 匿名字段
				field
			}
			此时B可以直接访问A中字段
		2.模拟聚合关系： has -> a
			type C {
				field
			}
			type D {
				c C // 聚合关系
				field
			}
			这种情况不能直接A中的字段

	*/

}
