package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
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
	user := NewUser(conn, server)

	// 用户上线, 存入 onlineMap
	user.UserOnline()

	// 监听用户是否发送消息
	isLive := make(chan bool)

	// 接受客户端用户发送消息
	go func() {
		buf := make([]byte, 4096)

		for {
			n, err := conn.Read(buf)

			fmt.Println("n = ", n)

			if n == 0 {
				user.UserOffline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("conn Read err: ", err)
				return
			}

			// 接受用户发送消息
			msg := string(buf[:n-1])
			// 广播消息
			user.UserMsg(msg)

			isLive <- true
		}
	}()

	// 当前 handler 阻塞
	for {
		select {
		case <-isLive:
			// 重置下面定时器
		case <-time.After(time.Second * 4000):
			// 用户超时, 进行强制剔除
			user.UserMsg("ti chu")
			user.UserOffline()
			close(user.C)
			// conn.Close()

			return
		}
	}
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
