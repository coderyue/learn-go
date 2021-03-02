package main

import "fmt"

func main() {
	/**
	接口的嵌套

	*/
	cat := Cat{}
	cat.test1()
	cat.test2()
	cat.test3()

	var a A = cat
	a.test1()

	var b B = cat
	b.test2()

	var c C = cat
	c.test1()
	c.test2()
	c.test3()

	var c2 A = c
	c2.test1()

}

type A interface {
	test1()
}

type B interface {
	test2()
}

type C interface {
	A
	B
	test3()
}

type Cat struct {
}

func (c Cat) test1() {
	fmt.Println("test1()...")
}

func (c Cat) test2() {
	fmt.Println("test2()...")
}

func (c Cat) test3() {
	fmt.Println("test3()...")
}
