package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	/**
	bufio: 高效io读写
		buffer缓存
		io：input/output
		可以通过修改偏移位置，来修改写入位置

	将io包下的Reader，Writer对象进行包装，带缓存的包装，提高读写效率
	 	func (b *Writer) Write(p []byte) (nn int, err error)
	    func (b *Writer) WriteByte(c byte) error
	    func (b *Writer) WriteRune(r rune) (size int, err error)
	    func (b *Writer) WriteString(s string) (int, error)

	*/

	fileName := "cc.txt"
	file1, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer file1.Close()
	w1 := bufio.NewWriter(file1)
	// 高效流，也可以设置偏移位置
	//file1.Seek(10, io.SeekStart)
	n, err := w1.WriteString("helloworld")
	fmt.Println(n)
	w1.Flush() //刷新数据到缓冲区

	//for i := 1; i <= 1000; i++ {
	//	w1.WriteString(fmt.Sprintf("%d:hello", i))
	//}
	//w1.Flush()

}
