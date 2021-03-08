package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/**
	写出数据

	*/

	fileName := "F:/goworkspace/l_file/a.txt"
	// 打开文件
	// 写出数据
	// 关闭文件
	//file, err := os.Open(fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	defer file.Close()

	// 写出数据
	//bs := []byte{65, 66, 67, 68, 69, 70}
	bs := []byte{97, 98, 99, 100}
	//n, err := file.Write(bs)
	n, err := file.Write(bs[:2])
	fmt.Println(n)
	HandleErr(err)
	file.WriteString("\n")

	// 直接写出字符串
	n, err = file.WriteString("HelloWorld")
	fmt.Println(n)
	HandleErr(err)
	file.WriteString("\n")

	n ,err = file.Write([]byte("today"))
	fmt.Println(n)
	HandleErr(err)


}

func HandleErr(err error) {
	if err != nil {
		log.Fatalln(err)
		return
	}
}
