package main

import (
	"fmt"
	"os"
)

func main() {

	/**
	FileInfo: 文件信息
		interface
			Name(), 文件名
			Size(), 文件大小，字节为单位
	 */

	//fileInfo, err := os.Stat("F:\\goworkspace\\l_file\\test.txt")
	fileInfo, err := os.Stat("F:/goworkspace/l_file/test.txt")
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Printf("%T\n", fileInfo)
	// 文件名
	fmt.Println(fileInfo.Name())
	// 文件大小
	fmt.Println(fileInfo.Size())
	// 是否是目录
	fmt.Println(fileInfo.IsDir())
	// 修改时间
	fmt.Println(fileInfo.ModTime())
	// 权限
	fmt.Println(fileInfo.Mode()) // -rw-rw-rw-
	/**
	Linux下有两种文件权限表示方式，即“符号表示”和“八进制表示”
	1. 符号表示
	-      ---    ---    ---
	type  owner  group  others
	文件权限分配， 读写可执行 分别对应 r w x， 如果没有哪一个权限，用 - 代替
	(-文件 d目录 |连接符号)

	2.八进制表示
	r -> 004
	w -> 002
	x -> 001
	- -> 000

	eg: 0755, 0777,

	 */


}
