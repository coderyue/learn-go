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

	将io包下的Reader，Writer对象进行包装，带缓存的包装，提高读写效率

		ReadBytes()
		ReadString()
		ReadLine()
	 */

	//fileName := "E:/我的壁纸/timg.jpg"
	//file, err := os.Open(fileName)
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return
	//}
	//bf := bufio.NewReader(file)
	// 1. Read() 高效读取
	//bs := make([]byte, 1024, 1024)
	//n, err := bf.Read(bs)
	//fmt.Println(n)
	//fmt.Println(bs[:n])

	// 2. ReadLine() 不建议使用
	//data, flag, err := bf.ReadLine()
	//fmt.Println(string(data))
	//fmt.Println(flag)
	//fmt.Println(err)

	// 3. ReadString()
	//s1, err := bf.ReadString('\n')
	//fmt.Println(err)
	//fmt.Println(s1)

	//for {
	//	s1, err := bf.ReadString('\n')
	//	if err == io.EOF {
	//		fmt.Println("读取完毕....")
	//		break
	//	}
	//	fmt.Println(s1)
	//}

	// ReadBytes()
	//data, err := bf.ReadBytes('\n')
	//fmt.Println(err)
	//fmt.Println(data)

	// Scanner  输入有空格，那么空格后面的数据读不到
	//s2 := ""
	//fmt.Scanln(&s2)
	//fmt.Println(s2)

	b2 := bufio.NewReader(os.Stdin)
	s2, _ := b2.ReadString('\n')
	fmt.Println(s2)


}
