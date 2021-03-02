package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	files, err := filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err)
		return
	}
	fmt.Println("files: ", files)

	// 忽略了错误， 这个不好， 实际是匹配模式不正常
	//files, _ := filepath.Glob("[")
	//if err != nil && err == filepath.ErrBadPattern {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("files: ", files)
}
