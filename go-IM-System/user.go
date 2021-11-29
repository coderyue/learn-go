package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// 监听当前user channel的方法，  一旦有消息， 就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}

// 用户上线功能
func (this *User) Online() {
	// 用户上线， 将用户加入到onlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.server.BroadCast(this, "上线了")
}

// 用户下线功能
func (this *User) Offline() {
	// 用户下线， 将用户从onlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 	广播用户下线消息
	this.server.BroadCast(this, "下线了")
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息功能
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前在线用户都有哪些
		for _, user := range this.server.OnlineMap {
			sendMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(sendMsg)
		}
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式 rename|张三
		// 判断张三是否存在
		_, ok := this.server.OnlineMap[msg[7:]]
		if ok {
			this.SendMsg("当前用户名被使用\n")
		} else {
			this.server.mapLock.Lock()
			//this.Name = msg[7:]  // 这样处理会使里面多个key指向同一个用户, 所以删除再添加
			newName := strings.Split(msg, "|")[1]
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.SendMsg("更新名字成功" + this.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		//消息格式： to|张三|消息内容
		// 1. 获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("消息格式不正确, 请使用 \"to|张三|消息内容\"格式。\n")
			return
		}
		// 根据用户名得到对方user对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("该用户不在线\n")
			return
		}
		// 2. 获取消息内容
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息内容， 请重发!")
			return
		}
		remoteUser.SendMsg(this.Name + "对您说:" + content)
	} else {
		this.server.BroadCast(this, msg)
	}

}
