package main

import (
	"fmt"
	"strings"
)

func main() {
	/**
	strings包下的关于字符串的函数

	*/

	s1 := "helloworld"
	// 1. 是否包含指定的内容 -->
	fmt.Println(strings.Contains(s1, "l"))
	// 2. 是否包含chars中任意一个字符即可
	fmt.Println(strings.ContainsAny(s1, "abcd"))
	// 3. 统计substr在s中出现的次数
	fmt.Println(strings.Count(s1, "llod"))

	// 3. 以xxx前缀开头， 以xxx后缀结尾
	s2 := "2020年5月.txt"
	if strings.HasPrefix(s2, "2020") {
		fmt.Println("文件是2020年的")
	}
	if strings.HasSuffix(s2, "txt") {
		fmt.Println("文件是文本文档")
	}

	// 索引
	// helloworld
	fmt.Println(strings.Index(s1, "llod")) // 当没有是返回-1
	fmt.Println(strings.IndexAny(s1, "abcd"))
	fmt.Println(strings.LastIndex(s1, "l"))

	// 字符串的拼接
	ss1 := []string{"werw", "aaa", "ddd"}
	fmt.Println(strings.Join(ss1, ""))

	// 切割
	s4 := "123,12,4,34,32,42,5"
	ss2 := strings.Split(s4, ",")
	fmt.Println(ss2)

	for i := 0; i < len(ss2); i++ {
		fmt.Println(ss2[i])
	}

	// 重复，自己拼接自己count次
	s5 := strings.Repeat("hello", 5)
	fmt.Println(s5)

	// 替换, 参数n是替换几个，传入负值为全部替换
	s6 := strings.Replace(s1, "l", "3", 1)
	fmt.Println(s6)
	s7 := strings.ReplaceAll("helloworld", "l", "3")
	fmt.Println(s7)

	// 变大写， 变小写
	s8 := "hellO WorLd**123.."
	fmt.Println(strings.ToLower(s8))
	fmt.Println(strings.ToUpper(s8))

	/**
	截取字符串
	其他语言有substring(start, end) --> substr go中没有这个函数， 使用切片
	str[start:end] -->substr
	*/
	fmt.Println(s1)
	fmt.Println(s1[:5])
	fmt.Println(s1[5:])

}
