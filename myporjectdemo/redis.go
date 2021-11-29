package main

import (
	"time"
	//"github.com/gomodule/redigo/redis"
	"github.com/gitstliu/go-redis-cluster"
)

func main() {

	//conn, _ := redis.Dial("tcp", ":6379")
	//defer conn.Close()
	//
	//// 直接执行, do方法，只有一步， 下面的send方法是三步， send是先放到缓冲区
	//// 这几个方法返回的都是字节码
	////reply, err := conn.Do("set", "c1", "hello")
	//
	//conn.Send("get", "c1")
	//conn.Send("set", "aaa", "ccc")
	//conn.Flush()
	//// 这个返回值只接受第一条执行的返回结果(中间send多条的时候)
	//reply, err := conn.Receive()
	////str, err := redis.String(conn.Receive())       // redis包中封装的把字节码转换为想要的类型
	//if err != nil {
	//	fmt.Println(">>> redis conn err:", err)
	//}
	//fmt.Println(reply)
	//
	//replay, err := redis.Values(conn.Receive())	// 返回的是interface类型,  使用values再使用scan可以赋值给想要的变量
	//var name string
	//redis.Scan(replay, &name)


	// 下面是操作redis集群
	cluster, _ :=redis.NewCluster(&redis.Options{
		StartNodes: []string{"ip:port","ip:port"},
		ConnTimeout: 50 * time.Millisecond,
		ReadTimeout: 50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		KeepAlive: 16,
		AliveTime: 60 * time.Millisecond,
	})
	defer cluster.Close()

	cluster.Do("set", "hello", "world")


}
