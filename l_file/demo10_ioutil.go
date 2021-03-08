package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	/**
	ioutil包
		ReadFile()
		WriteFile()
		ReadDir()

	 */

	//读取文件中的所有数据
	//fileName := "cc.txt"
	//data, err := ioutil.ReadFile(fileName)
	//fmt.Println(err)
	//fmt.Println(string(data))

	// 写出数据
	//fileName := "bb.txt"
	//s1 := "从前有座山，山上有座庙"
	//// 该方法重新写的时候，会擦除原有的数据
	//err := ioutil.WriteFile(fileName, []byte(s1), os.ModePerm)
	//fmt.Println(err)

	// ReadAll()
	//s2 := "a man`s success is relay on his own effort"
	//r1 := strings.NewReader(s2)
	//data, err := ioutil.ReadAll(r1)
	//fmt.Println(err)
	//fmt.Println(data)
	//fmt.Println(string(data))

	// ReadDir() // 读取一个目录下的子内容，包括文件和文件夹， 但是只能读取一层
	//dirName := "F:/"
	//fileInfos, err := ioutil.ReadDir(dirName)
	//fmt.Println(err)
	//fmt.Println(len(fileInfos))
	//
	//for i := 0; i < len(fileInfos); i++ {
	//	fmt.Println(fileInfos[i].Name())
	//}

	// 创建临时文件或者文件夹
	dir, err := ioutil.TempDir("F:/goworkspace/l_file/", "Test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(dir)
	fmt.Println(dir)

	file1, err := ioutil.TempFile(dir, "temp")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(file1.Name())
	fmt.Println(file1.Name())
	time.Sleep(time.Second * 1)


}
