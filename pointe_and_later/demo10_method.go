package main

import "fmt"

func main() {
	/**
	OOP中的继承性：
		如果两个类(class)存在继承关系， 其中一个是子类，另一个作为父类，那么：

		1.子类可以直接访问父类的属性和方法
		2.子类可以新增自己的属性和方法
		3.子类可以重写父类的方法(override，就是将父类已有的方法，重新实现)

	go语言的结构体嵌套
		1.模拟继承性： is -> a
			type A struct {
				field
			}
			type B struct {
				A //匿名字段
			}

		2.模拟聚合关系： has -> a
			type C struct {
				field
			}
			type D struct {
				c C // 聚合关系
			}

	*/

	// 1.创建Person类型
	p1 := Person{name: "王二狗", age: 30}
	fmt.Println(p1.name, p1.age)
	p1.eat() // 父类对象，访问父类方法

	// 2. 创建Student类型
	s1 := Student{Person: Person{name: "ruby", age: 18}, school: "家里蹲大学"}
	fmt.Println(s1.name, s1.age, s1.school) // 子类对象可以直接访问父类的字段, (其实就是提升字段)， s1.Person.name

	s1.eat() // 子类对象访问父类方法，
	// 当子类重新写接受者是子类的同样方法时， 相当于重写了父类的方法，此时调用就是调用的子类的方法

	s1.Study() // 子类的对象， 访问自己新增的方法

}

// 1. 定义一个"父类"
type Person struct {
	name string
	age  int
}

// 2. 定义一个"子类"
type Student struct {
	Person // 结构体嵌套， 模拟继承性
	school string
}

func (p Person) eat() {
	fmt.Println(p.name, "父类在吃窝窝头")
}

func (s Student) eat() {
	fmt.Println(s.name, "子类在吃炸鸡，喝啤酒")
}

func (s Student) Study() {
	fmt.Println("子类新增的方法,学生学习啦")
}
