package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    /*
       遍历文件夹

    */

    dirName := "G:/goworkspace/learn-go"

    listFiles(dirName, 0)
}

func listFiles(dirName string, level int) {
    // level 用来记录当前递归的层次，生成带有层次感的空格
    fileInfos, err := ioutil.ReadDir(dirName)
    s := "|--"
    for i := 0; i < level; i++ {
        s = "|  " + s
    }
    if err != nil {
        log.Fatal(err)
    }
    for _, f := range fileInfos {
        fileName := dirName + "/" + f.Name()
        fmt.Println(s + fileName)
        if f.IsDir() {
            listFiles(fileName, level + 1)
        }
    }
}
