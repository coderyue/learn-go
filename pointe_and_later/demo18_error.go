package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err) // open /test.txt: The system cannot find the file specified.
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("1.op:", ins.Op)
			fmt.Println("2.path", ins.Path)
			fmt.Println("3.err:", ins.Err)
		}
		return
	}
	fmt.Println(f.Name(), "文件打开成功..")
}
