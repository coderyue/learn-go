package main

import "fmt"

func main() {

	/**
	结构体指针

	数据类型
		值类型  	int, float, bool, string, array, struct
		引用类型	slice, map, function, pointer

	通过指针：
		new(), 不是nil， 空指针
			指向了新分配的类型和内存空间， 里面存储的零值
	*/

	// 结构体是值类型
	p1 := Person{"abc", 12, "女"}
	fmt.Println(p1)
	fmt.Printf("%T,%p\n", p1, &p1)
	p2 := p1
	p2.name = "test"
	fmt.Println(p1)
	fmt.Println(p2)

	// 定义指针结构体
	var pp1 *Person
	pp1 = &p1
	//(*pp1).age = 20
	pp1.age = 14
	fmt.Println(p1)

	// 使用内置函数new， go语言中专门用于创建某种类型的指针的函数
	pp2 := new(Person)
	fmt.Printf("%p, %T\n", pp2, pp2)
	fmt.Println(*pp2)
	pp2.name = "李小花"
	pp2.age = 16
	pp2.sex = "女"
	fmt.Println(*pp2)
	fmt.Println(pp2)

	pp3 := new(int)
	fmt.Println(pp3)
	*pp3 = 12
	fmt.Println(*pp3)

}

type Person struct {
	name string
	age  int
	sex  string
}
