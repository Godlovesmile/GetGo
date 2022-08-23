package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建 server 接口

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}

	return server
}

func (server *Server) Handler(conn net.Conn) {
	// 当前链接处理业务
	fmt.Println("链接建立成功")
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
