package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	/**

	Seek(offset int64, whence int) (int64, error) // 设置指针光标的位置
		第一个参数：偏移量
		第二个参数：如何设置

	*/
	fileName := "test.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// 读写
	bs := []byte{0}
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(2, io.SeekStart)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(4, 0) // SeekStart
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(3, io.SeekCurrent)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(0, io.SeekEnd)
	file.WriteString("ABC")


}
