package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip string
	Port int

	// 在线用户列表
	OnlineMap map[string]*User
	mapLock sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip: ip,
		Port: port,
		OnlineMap: make(map[string]*User),
		Message: make(chan string),
	}

	return server
}

// 启动服务器接口
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// close listen socket
	defer listener.Close()

	// 启动监听message的goroutines
	go this.ListenMessage()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			break
		}
		// do handler
		go this.Handler(conn)
	}

	// close listen socket
}

func (this *Server) Handler(conn net.Conn) {
	// 当前链接的业务
	//fmt.Println("链接建立成功")

	// 用户上线，将用户加入到onlineMap中
	user := NewUser(conn, this)

	//this.mapLock.Lock()
	//this.OnlineMap[user.Name] = user
	//this.mapLock.Unlock()

	// 广播当前用户上线消息
	//this.BroadCast(user, "上线了")
	user.Online()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if n == 0 {
			//this.BroadCast(user, "下线了")
			user.Offline()
			return
		}

		if err != nil && err != io.EOF {
			fmt.Println("Conn read err:", err)
			return
		}

		// 提取用户发送的消息, 去除'\n'
		msg := string(buf[:n-1])

		// 将得到的消息进行广播
		//this.BroadCast(user, msg)
		// 用户针对msg进行处理
		user.DoMessage(msg)

		// 用户的任意消息，代表当前用户是活跃的
		isLive <- true
	}()

	// 当前handler阻塞
	//select {}

	for {
		select {
		case <- isLive:
			// 当前用户是活跃的， 应该重置定时器
			// 不做任何事情， 为了激活select， 更新下面的定时器
		case <- time.After(time.Second * 10):
			// 已经超时
			// 将当前的user强制关闭
			user.SendMsg("你被踢了")

			// 销毁可用的资源
			close(user.C)

			// 关闭连接
			conn.Close()

			// 退出当前的handler
			return // runtime.Goexit()
		}
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

// 监听Message广播消息channel的goroutine， 一旦有消息， 就发送给全部的在线user
func (this *Server) ListenMessage() {
	for {
		msg := <- this.Message

		// 将msg发送给全部在线user
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}
