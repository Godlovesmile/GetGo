package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户 map 表
	OnlineMap map[string]*User
	// 锁
	mapLock sync.RWMutex

	// 消息广播 channel
	Message chan string
}

// 创建 server 接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// Message 消息写入
func (server *Server) BroadCast(user *User, msg string) {
	sendMessage := "[" + user.Name + "]" + ":" + msg

	server.Message <- sendMessage
}

// 监听并广播消息
func (server *Server) ListenMessager() {
	for {
		msg := <-server.Message

		// 消息广播所有在线 user
		server.mapLock.Lock()
		for _, user := range server.OnlineMap {
			user.C <- msg
		}
		server.mapLock.Unlock()
	}
}

func (server *Server) Handler(conn net.Conn) {
	// 当前链接处理业务
	// fmt.Println("链接建立成功")
	user := NewUser(conn)

	// 用户上线, 存入 onlineMap
	server.mapLock.Lock()
	server.OnlineMap[user.Name] = user
	server.mapLock.Unlock()

	// 有人上线, message 消息增加
	server.BroadCast(user, "已上线")
}

// 启动服务接口
func (server *Server) Start() {
	// 1. socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))

	if err != nil {
		fmt.Println("listen err: ", err)
		return
	}

	// 4. close listen
	defer listener.Close()

	// 启动 Message 监听 gorutine
	go server.ListenMessager()

	for {
		// 2. accept
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("listen accept err: ", err)
			continue
		}

		// 3. do hander
		go server.Handler(conn)
	}
}
