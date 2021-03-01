package main

import "fmt"

func main() {
	/*
		switch语句
		1.标准写法
			switch 变量名 {
				case 数值1: 分支1
				case 数值2: 分支2
				case 数值3: 分支3
				default:
					最后一个分支
			}
		其他写法
		其他写法1， 省略switch后的变量，相当于直接作用在true上
		switc { // true
		}

		其他写法2. case后可以同时跟随多个数值
		switch 变量 {
		case 数值1, 数值2, 数值3:
		case 数值4, 数值5:
		}

		其他写法3， switch后可以多一条初始化语句
		switch 初始化语句;变量 {
		}

	*/

	// =====================标准写法============================
	num := 5
	switch num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
		//default:
		//	fmt.Println("数据有误")
	}

	fmt.Println("=======================")
	// =====================其他写法1============================
	switch {
	case true:
		fmt.Println("ture...")
	case false:
		fmt.Println("false")
	}

	// 代码执行从上往下， 注意代码判断顺序， 否则结果错误
	score := 80
	switch {
	case score >= 60:
		fmt.Println("D")
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 0 && score < 60:
		fmt.Println("E")
	default:
		fmt.Println("成绩有误")
	}
	fmt.Println("=======================")
	// =====================其他写法1============================

	letter := "f"
	switch letter {
	case "u", "f", "o":
		fmt.Println("ufo...")
	case "l", "i", "n":
		fmt.Println("lin...")
	}
	fmt.Println("=======================")
	// =====================其他写法1============================

	// language 作用域只在switch中
	switch language := "go"; language {
	case "go":
		fmt.Println("go language")
	case "java":
		fmt.Println("java language")
	case "dart":
		fmt.Println("dart language")
	case "python":
		fmt.Println("python language")
	}
	//fmt.Println(language) //undefined: language
}
