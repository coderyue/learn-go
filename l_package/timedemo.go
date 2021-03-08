package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/**
	time包：

	 */

	// 1.获取的当前时间
	t1 := time.Now()
	fmt.Printf("%T\n", t1)
	fmt.Println(t1) //2021-03-03 15:22:17.7119777 +0800 CST m=+0.004242401

	// 2.获取指定的时间
	t2 := time.Date(2008, 7, 15, 16, 23, 12, 0, time.Local)
	fmt.Println(t2)

	// 3. time-->string之间的转换
	/**
	t1.Format("格式模板")-->string
		模板日期必须是固定：06-1-2-3-4-5
	 */
	s1 := t1.Format("2006年01月02日 15:04:05")
	fmt.Println(s1)

	s2 := t1.Format("2006/01/02")
	fmt.Println(s2)

	// string ->time
	/*
	time.Parse("模板", str) -> time, err
	 */
	s3 := "1999年10月10日" // string
	// 模板格式和字符串对应不上， 就会报错
	t3, err := time.Parse("2006年01月02日", s3)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(t3)
	fmt.Printf("%T\n", t3)

	fmt.Println(t1.String())

	// 根据当前的时间， 获取指定内容
	year, month, day := t1.Date() // 年，月，日
	fmt.Println(year, month, day)

	hour, min, sec := t1.Clock()
	fmt.Println(hour, min, sec)

	year2 := t1.Year()
	fmt.Println("年：", year2)
	fmt.Println(t1.YearDay()) // 今年一共过了多少天

	fmt.Println("月：", t1.Month())
	fmt.Println("日：", t1.Day())
	fmt.Println("时:", t1.Hour())
	fmt.Println("分：", t1.Minute())
	fmt.Println("秒：", t1.Second())
	fmt.Println("纳秒：", t1.Nanosecond())
	fmt.Println("星期: ", t1.Weekday())

	// 5. 时间戳， 指定的日期，距离1970年1月1日0点0时0分0秒的时间差值： 秒， 纳秒

	t4 := time.Date(1970, 1, 1, 1, 0, 0, 0 , time.UTC)
	timeStamp1 := t4.Unix() //秒的差值, 秒的时间戳
	fmt.Println(timeStamp1)
	fmt.Println(t1.Unix())

	timeStamp3 := t4.UnixNano()
	fmt.Println(timeStamp3) // 3600 000 000 000
	fmt.Println(t1.UnixNano())

	// 6. 时间间隔
	t5 := t1.Add(time.Minute)
	fmt.Println(t1)
	fmt.Println(t5)

	t6 := t1.Add(time.Hour * 24)
	fmt.Println(t6)

	t7 := t1.AddDate(1, 0, 0)
	fmt.Println(t7)

	// 时间间隔
	d1 := t5.Sub(t1)
	fmt.Printf("%T, %v\n", d1, d1)

	// 7.睡眠
	time.Sleep(3 * time.Second) //睡眠3秒
	fmt.Println("睡眠了3秒...")

	// 睡眠[1-10]的随机数
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10) + 1
	fmt.Printf("%T, %v\n", randNum, randNum)
	// 相同类型的才能计算 把randNum变成Duration在计算
	time.Sleep(time.Duration(randNum) * time.Second)

	fmt.Println("睡醒了")

}
