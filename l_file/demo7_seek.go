package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	/**
	断点续传
		文件传递： 文件复制
	 */

	srcFile := "E:/我的壁纸/timg.jpg"
	destFile := srcFile[strings.LastIndex(srcFile, "/") + 1:]
	fmt.Println(destFile)
	HttpCopy(srcFile, destFile)
	//tempFile := destFile + ".temp.jpg"
	//fmt.Println(tempFile)
	//
	//file1, err := os.Open(srcFile)
	//HandleError(err)
	//file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	//HandleError(err)
	//file3, err := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	//HandleError(err)
	//
	//defer file1.Close()
	//defer file2.Close()
	//defer file3.Close()
	//
	//// 先读取临时文件中的数据， 再seek
	//file3.Seek(0, io.SeekStart)
	//bs := make([]byte, 100, 100)
	//n1, err := file3.Read(bs)
	////HandleError(err)
	//countStr := string(bs[:n1])
	//count, err := strconv.ParseInt(countStr, 10, 64)
	////HandleError(err)
	//fmt.Println(count)
	//
	//// 根据count设置偏移量, 设置读， 写的位置
	//file1.Seek(count, io.SeekStart)
	//file2.Seek(count, io.SeekStart)
	//data := make([]byte, 1024, 1024)
	//n2 := -1 // 读取的数据量
	//n3 := -1 // 写出的数据量
	//total := int(count) // 读取的总量
	//
	//// 复制文件
	//for {
	//	n2, err = file1.Read(data)
	//	if err == io.EOF || n2 == 0 {
	//		fmt.Println("文件复制完毕", total)
	//		file3.Close()
	//		// 这里不关闭， 无法删除临时文件
	//		os.Remove(tempFile)
	//		break
	//	}
	//	n3, err = file2.Write(data[:n2])
	//	total += n3
	//	// 将复制的总量存储到临时文件中
	//	file3.Seek(0, io.SeekStart)
	//	file3.WriteString(strconv.Itoa(total))
	//	fmt.Printf("total: %d\n", total)
	//
	//	// 假装断电
	//	//if total > 8000 {
	//	//	panic("假装断电了。。。")
	//	//}
	//}

}

// 讲的是有另一个文件，来存储总共传递来了多少数据，每次都是先读取读到了那里，然后从那里继续复制
// 这样万一存储的值和实际传输的值，对不上的时候，最后下载下来数据也不对
func HttpCopy(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	fmt.Println(file1)
	HandleError(err)
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	fmt.Println(file2)
	HandleError(err)
	defer file1.Close()
	defer file2.Close()

	// 先获取已经传输了多少, 没有目标文件的时候会创建，这里不用判断目标文件是否存在
	bsTrans, err := ioutil.ReadFile(destFile) // 读取所有数据
	HandleError(err)
	total := len(bsTrans)
	file1.Seek(int64(len(bsTrans)), io.SeekStart)
	file2.Seek(int64(len(bsTrans)), io.SeekStart)
	data := make([]byte, 1024, 1024)
	for {
		n, err := file1.Read(data)
		if err == io.EOF || n == 0 {
			fmt.Println("文件复制完毕。。。")
			break
		}
		total += n
		//if total > 8000 {
		//	panic("假装断电。。。。")
		//}
		file2.Write(data[:n])
	}
	return 0, nil
}

func HandleError(err error)  {
	if err != nil {
		fmt.Println("err: ", err)
	}
}