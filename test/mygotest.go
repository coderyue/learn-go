package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func swap(x, y string) (string, string) {
	return y, x
}
///*添加方法*/
func add(x, y int) int {
	return x + y
}

// 明明返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// var 语句用于声明一个变量列表, 可以出现在包或函数级别
var c, python, java = true, false, "no"

// 变量初始化， 变量声明可以包含初始值， 见上面
//   := 只能在函数内部使用，函数外的每个语句都必须以关键字开始  eg：var func等
// 常量声明和变量类似， 关键字是const   常量不能用:=语法声明

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Sqrt(x float64) float64 {
	var z = x / 2
	for math.Abs(z * z - x) > 0.000000000000001 {
		z -= (z * z - x) / (2 * z)
		fmt.Println(z, "==========")
	}
	return z
}

func main() {
	fmt.Println("hello world:", rand.Intn(10), math.Sqrt(7), "", math.Pi)
	fmt.Println()
	fmt.Println(add(2, 79))
	a, b := swap("x", "y")
	fmt.Println(a, b)
	fmt.Println(split(3))
	var str string
	fmt.Println(str)

	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(u)

	v := 0.867 + 0.5i // 修改这里！
	fmt.Printf("v is of type %T\n", v)

	sum := 1
	for ; sum < 1000;	 {
		sum += sum
	}
	fmt.Println(sum)

	while := 1
	for while < 100 {
		while += while
	}
	fmt.Println(while)
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(Sqrt(5))
	fmt.Println(sqrt(5))

	testSwitch()
	testSwitchWeekday()
	fmt.Println(int(time.Sunday + 1))
	noConditionSwitch()

	i = 10
	i ++
	defer fmt.Println(i)
	i ++
	fmt.Println(i)

	fmt.Println("counting")

	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}

	fmt.Println("done")
	fmt.Println(testtext())


	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j

	*p = *p / 37
	fmt.Println(j)

	fmt.Println(MyStruct{1, 2})

	t := MyStruct{1, 2}
	tt := MyStruct2{1, 2}
	t.X = 4
	fmt.Println(t.X)

	pp := &t
	pp.X = 1e9

	pp = (*MyStruct)(&tt)
	fmt.Println(pp)

	fmt.Println(v1, v2, v3, v4)

	// 数组  [n]T 表示拥有n个T类型的值的数组	
	var aa [2]int
	aa[0] = 1
	aa[1] = 2
	fmt.Println(aa[0], aa[1])
	fmt.Println(a)

	var abc [2]string
	abc[0] = "dsf"
	abc[1] = "wey"
	fmt.Println(abc[0], abc[1])
	fmt.Println(abc)

	// 切片
	params := [6]int{2, 3, 4, 5, 6, 7}
	var s []int = params[1:4]
	var testsplice []int = params[0:]
	fmt.Println(s)
	fmt.Println(testsplice)



}


var (
	v1 = MyStruct{1, 2}
	v2 = MyStruct{X: 2}
	v3 = MyStruct{}
	v4 = &MyStruct{1, 2}
)

type MyStruct struct {
	X int
	Y int
}

type MyStruct2 struct {
	X int
	Y int
}



func testtext() string {
	return "this is a test"
}

func noConditionSwitch() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour() < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}

}

func testSwitchWeekday() {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("in two days")
	default:
		fmt.Println("too far away")
	}

}


func testSwitch() {
	num := 10
	switch num {
	case 5:
		fmt.Println("os 5")
	case 10:
		fmt.Println("linux")
	default:
		fmt.Println("%s. \n", num)
	}

	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s. \n", os)
	}

}

