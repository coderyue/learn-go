package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/**
	读取数据：
		Reader接口：
			Read(p []byte)(n int, error)

	*/

	// 读取本地text.txt文件中的数据
	fileName := "F:/goworkspace/l_file/test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	// 关闭文件
	defer file.Close()
	// 读取数据
	bs := make([]byte, 4, 4)
	/*
	// 第一次读取数据
	n, err := file.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(bs) //[97 98 99 100]
	fmt.Println(string(bs))

	// 第二次读取数据
	n, err = file.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(bs) // [101 102 103 104]
	fmt.Println(string(bs))

	// 第三次读取数据
	n, err = file.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(bs) // [105 106 103 104]
	fmt.Println(string(bs))

	// 第四次读取数据
	n, err = file.Read(bs)
	fmt.Println(err) // 到结尾了是 EOF
	fmt.Println(n)
	fmt.Println(bs)
	*/

	for {
		n, err := file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("读取到了文件的末尾，结束读取操作..")
			break
		}
		fmt.Println(string(bs[:n]))
	}


}
