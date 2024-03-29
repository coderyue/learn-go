package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	/**
	文件操作
	1.路径：
		相对路径：relative
			ab.txt
			相对于当前工程
		绝对路径：absolute
			F:/goworkspace/l_file/test.txt

	2. 创建文件夹，如果文件夹存在，创建失败
		os.MkDir()， 创建一层
		os.MkDirAll(), 可以创建多层

	3. 创建文件， Create采用模式0666(任何人都可读写，不可执行) 创建一个名为name的文件，如果文件已存在会截断它(为空文件)
	4.打开文件： 让当前的程序， 和指定的文件之间建立一个连接
		os.Open(fileName)
		os.OpenFile(fileName, mode, perm)
	5. 关闭文件
		file.Close()
	6. 删除文件或目录：慎重
		os.Remove()  删除文件或空目录
		os.RemoveAll() 删除所有

	 */

	// 1. 路径
	//fileName1 := "F:\\goworkspace\\l_file\\test.txt" // 这样写的时候获取上级目录错误
	fileName1 := "F:/goworkspace/l_file/test.txt"
	fileName2 := "test.txt"

	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2)) //F:\goworkspace\l_file\test.txt <nil>

	fmt.Println("获取父目录：", path.Join(fileName1, ".."))

	// 2.创建目录
	//err := os.Mkdir("F:/goworkspace/l_file/bb", os.ModeAppend)
	//if err != nil {
	//	fmt.Println("err: ", err)
	//  return
	//}
	//fmt.Println("文件夹创建成功。。。 ")
	//
	//err2 := os.MkdirAll("F:/goworkspace/l_file/bb/cc/dd", os.ModeAppend)
	//if err2 != nil {
	//	fmt.Println("err: ", err2)
	//  return
	//}
	//fmt.Println("多级文件夹创建成功。。。。")

	// 3. 创建文件；Create采用模式0666(任何人都可读写，不可执行) 创建一个名为name的文件， 如果文件已经存在会截断它(为空文件)
	//file1, err := os.Create("F:/goworkspace/l_file/bb/test.txt")
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return
	//}
	//fmt.Println("文件创建成功...", file1)

	// 4. 打开文件：
	//file2, err := os.Open(fileName1)
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return
	//}
	//fmt.Println(file2)

	/**
	第一个参数： 文件名称
	第二个参数： 文件的打开方式
		// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
		O_RDONLY int = syscall.O_RDONLY // open the file read-only.
		O_WRONLY int = syscall.O_WRONLY // open the file write-only.
		O_RDWR   int = syscall.O_RDWR   // open the file read-write.
		// The remaining values may be or'ed in to control behavior.
		O_APPEND int = syscall.O_APPEND // append data to the file when writing.
		O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
		O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
		O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
		O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	第三个参数： 文件的权限， 文件不存在创建文件， 需要指定权限
	 */
	//file3, err := os.OpenFile(fileName1, os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return
	//}
	//fmt.Println(file3)
	//
	//// 5. 关闭文件
	//file3.Close()

	// 6. 删除文件或文件夹
	//rmFile:= "F:/goworkspace/l_file/cc/dd/a.txt"
	//err := os.Remove(rmFile) // 删除文件
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return
	//}
	//fmt.Println("删除文件成功。。。")

	//rmFile := "F:/goworkspace/l_file/cc/dd"
	//err := os.Remove(rmFile) // 删除空文件夹
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return
	//}
	//fmt.Println("删除文件夹成功。。。。")

	rmFile := "F:/goworkspace/l_file/cc"
	err := os.RemoveAll(rmFile)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("删除目录成功。。。。")

}