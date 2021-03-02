package main

import "fmt"

func main() {
	/**
	接口: interface
		在Go中接口是一组方法签名

		当某个类型为这个借口中的所有方法提供了方法的实现，它被称为实现接口

		Go语言中， 接口和类型的实现关系，是非侵入式

		// 其他语言中， 要显示的定义
		class Mouse implements USB{}

	1.当需要接口类型的对象时，可以使用任意实现类对象代替
	2.接口对象不能访问实现类中的属性

	多态：一个一个事物的多种形态
		go语言通过接口模拟多态

		就一个接口的实现

	接口的用法：
		1.一个函数如果接受接口类型作为参数，那么实际上可以传入该接口的任意实现类型对象作为参数
		2.定义一个类型为接口类型，实际上可以赋值为任意实现类的对象

	鸭子类型： 只要特性想鸭子就当成鸭子， 弱类型语言

	*/
	m := Mouse{"罗技小红"}
	fmt.Println(m.name)
	testInterface(m)

	f := FastDisk{"闪迪64G"}
	fmt.Println(f.name)
	testInterface(f)

	var usb USB
	usb = m
	usb.start()
	usb.end()

	var usb2 USB
	usb2 = f
	usb2.start()
	usb2.end()

}

//定义接口
type USB interface {
	start() //USB设备开始工作
	end()   // USB设备结束工作
}

// 实现类
type Mouse struct {
	name string
}

type FastDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "鼠标准备就绪，可以开始工作了， 点点点。。")
}

func (m Mouse) end() {
	fmt.Println(m.name, "结束工作， 可以安全退出")
}

func (f FastDisk) start() {
	fmt.Println(f.name, "准备开始工作， 可以存储数据")
}

func (f FastDisk) end() {
	fmt.Println(f.name, "安全弹出...")
}

func testInterface(usb USB) {
	usb.start()
	usb.end()
}
