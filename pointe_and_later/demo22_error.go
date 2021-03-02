package main

import (
	"fmt"
	"math"
)

func main() {
	/**
	自定义error
		demo1
	*/

	radius := -3.0
	area, err := circlrArea(radius)
	if err != nil {
		fmt.Println(err)
		if err, ok := err.(*areaError); ok {
			fmt.Printf("半径是：%.2f\n", err.radius)
		}
		return
	}
	fmt.Println("圆形的面积是：", area)

}

// 1.定义一个结构题， 表示错误的类型
type areaError struct {
	msg    string
	radius float64
}

// 2.实现error接口，就是实现Error()方法
func (e *areaError) Error() string {
	return fmt.Sprintf("error: 半径：%.2f, %s", e.radius, e.msg)
}

func circlrArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"半径是非法的", radius}
	}
	return math.Pow(radius, 2) * math.Pi, nil
}
