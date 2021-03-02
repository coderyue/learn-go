package main

import "fmt"

func main() {

	/*
		自定义错误
			demo2
	*/

	//length := -12.0
	//width := -2.0
	length, width := 12.0, -2.0

	area, err := rectArea(length, width)
	if err != nil {
		fmt.Println(err)
		//if err, ok := err.(*areaError); ok {
		//	fmt.Printf("错误信息是： %s, length: %.2f, width: %.2f\n", err.msg, err.length, err.width)
		//}
		if err, ok := err.(*areaError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: 长度， %.2f, 小于零\n", err.length)
			}
			if err.widthNegative() {
				fmt.Printf("error: 宽度， %.2f, 小于零\n", err.width)
			}
		}
		return
	}
	fmt.Println("面积是：", area)

}

type areaError struct {
	msg    string
	length float64
	width  float64
}

func (e *areaError) Error() string {
	return e.msg
}

func (e areaError) lengthNegative() bool {
	return e.length < 0
}

func (e areaError) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	msg := ""
	if length < 0 {
		msg = "长度小于零"
	}
	if width < 0 {
		if msg != "" {
			msg += ", 宽度也小于零"
		} else {
			msg = "宽度小于零"
		}
	}

	if msg != "" {
		return 0, &areaError{msg, length, width}
	}
	return (length + width) * 2, nil
}
