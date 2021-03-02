package main

import "fmt"

/**
结构体在嵌套的时候
*/

type Person struct {
	name string
}

func (p Person) show() {
	fmt.Println("Person -->", p.name)
}

// 类型别名
type People = Person

func (p People) show2() { //Person.show redeclared in this block
	fmt.Println("PeoPle-->", p.name)
}

type Student struct {
	Person
	People
}

func main() {
	var s Student
	//s.name = "王二狗" // ambiguous selector s.name
	s.People.name = "李小花"
	s.Person.name = "王二狗"
	//s.show() //ambiguous selector s.show
	s.Person.show()
	fmt.Printf("%T, %T\n", s.Person, s.People)

	s.People.show()
	s.People.show2()

}
