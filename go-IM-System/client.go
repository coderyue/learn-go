package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int // 当前client的模式
}

func NewClient(serverIP string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
		flag:       999,
	}

	// 链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return nil
	}

	client.conn = conn
	// 返回对象
	return client
}

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	// init函数再main函数之前执行
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器ip地址(默认值127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认8888)")
}

func main() {
	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	//client := NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println(">>>>>>> 连接服务器失败...")
		return
	}

	// 单独开启一个goroutine去处理server的回执消息
	go client.DealResponse()

	fmt.Println(">>>>>>> 连接服务器成功...")

	// 启动客户端业务
	client.Run()
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>>>请输入合法数字<<<<<<<")
		return false
	}
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}

		// 根据不同模式处理不同的业务
		switch client.flag {
		case 1:
			// 公聊模式
			//fmt.Println("公聊模式选择...")
			client.PublicChart()
		case 2:
			// 私聊模式
			//fmt.Println("私聊模式选择...")
			client.PrivateChart()
		case 3:
			// 更新用户名
			//fmt.Println("更新用户名选择...")
			client.UpdateName()
		}
	}
}

func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err:", err)
		return
	}
}

func (client *Client) PrivateChart() {

	var remoteName string
	var chatMsg string
	// 列出所有用户, 选择
	client.SelectUsers()
	fmt.Print(">>>>>>请输入聊天对象[用户名], exit退出:")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>请输入消息内容，exit退出:")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit"{
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write err:", err)
					break
				}
			}

			chatMsg = ""
			fmt.Println(">>>>请输入消息内容，exit退出:")
			fmt.Scanln(&chatMsg)
		}

		client.SelectUsers()
		fmt.Print(">>>>>>请输入聊天对象[用户名], exit退出:")
		fmt.Scanln(&remoteName)
	}
	
}

func (client *Client) PublicChart() {

	var chartMsg string
	fmt.Println("请输入聊天内容，exit退出")
	fmt.Scanln(&chartMsg)

	for {
		if chartMsg != "exit" {
			if len(chartMsg) != 0 {
				sendMsg := chartMsg + "\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write err:", err)
					break
				}
			}
		}

		chartMsg = ""
		fmt.Println("请输入聊天内容，exit退出")
		fmt.Scanln(&chartMsg)
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println(">>>>请输入用户名:")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}

	return true
}

// 处理server回应的消息， 直接显示到标准输出即可
func (client *Client) DealResponse() {
	// 一旦client.conn有数据，就直接copy到stdout标准输出上， 永久阻塞监听
	io.Copy(os.Stdout, client.conn)
	// 等同于下面的写法
	//for {
	//	buf := make([]byte, 1024)
	//	client.conn.Read(buf)
	//	fmt.Println(buf)
	//}
}