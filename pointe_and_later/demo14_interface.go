package main

func main() {
	/**
	接口断言
		方式一：
			1. instance := 接口对象.(实际类型)  // 不安全， 会panic()
			2. instance, ok := 接口对象.(实际类型) // 安全

		方式二： switch
			switch instance := 接口对象.(type) {
			case 实际类型1:
				.....
			case 实际类型2:
				....
			....
			}

	*/

}

// 定义一个接口
type Shape interface {
	peri() float64 // 形状的周长
	area() float64 // 形状的面积
}
