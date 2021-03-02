package main

import "fmt"

func main() {
	/**
	go语言数据类型
		1.基本数据类型
			布尔： bool
				取值： true， false
			数值类型：
				整数， int
					有符号： 最高位表示符号位， 其余位表示数值
						int8  	// 后面数字表示占用多少个bit位
						int16
						int32
						int64
					无符号: 所有的位都表示数值
						uint8
						uint16
						uint32
						uint64
					int uint 根据平台表示32位或者64位整数   int 和int64 不等价

					byte: uint8
					rune: int32
				浮点数，
					float32
					float64
				复数
			字符串
		2， 复合数据类型
			eg： array, slice, map, function, pointe_and_later,	struct, interface, channel...

	*/
	// 整数
	var b1 bool
	b1 = true
	fmt.Printf("%T, %t\n", b1, b1)

	var b2 = false
	fmt.Printf("%T, %t\n", b2, b2)
	fmt.Println(b1, b2)

	var i1 int
	i1 = 1000
	fmt.Printf("%d\n", i1)
	//var i2 int64
	//i2 = i1 //cannot use i1 (type int) as type int64 in assignment

	var i3 uint8
	i3 = 100
	var i4 byte
	i4 = i3
	fmt.Printf("%d, %d\n", i3, i4)

	// 浮点
	var f1 float32
	f1 = 3.14
	fmt.Printf("%T, %.2f\n", f1, f1)

	var f2 float64
	f2 = 3.14
	fmt.Printf("%T, %f\n", f2, f2)
	fmt.Println(f1, f2)

	var f3 = 2.55
	fmt.Printf("%T", f3)
}
