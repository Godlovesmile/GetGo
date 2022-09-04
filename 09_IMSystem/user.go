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
	// 当前用户关联的 server
	server *Server
}

// 监听当前 user channel, 有消息, 直接推送给对应客户端
func (user *User) ListenMessage() {
	// for {
	// 	msg := <-user.C
	// 	_, err := user.conn.Write([]byte(msg + "\n"))

	// 	if err != nil {
	// 		fmt.Println("=== test err ===: ", err)
	// 	}
	// }

	for msg := range user.C {
		_, err := user.conn.Write([]byte(msg + "\n"))

		if err != nil {
			panic(err)
		}
	}

	err := user.conn.Close()

	if err != nil {
		panic(err)
	}
}

// 用户上线
func (user *User) UserOnline() {
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()

	user.UserMsg("shang xian")
}

// 用户下线
func (user *User) UserOffline() {
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()

	user.UserMsg("xia xian")
}

// 消息发送写入处理
func (user *User) SendMsg(msg string) {
	user.conn.Write([]byte(msg))
}

// 用户发消息
func (user *User) UserMsg(msg string) {
	// 特殊指令, 查询当前所有用户
	if msg == "who" {
		user.server.mapLock.Lock()
		for _, itemUser := range user.server.OnlineMap {
			onlineMsg := "[" + itemUser.Name + "]" + ":" + "online...\n"
			user.SendMsg(onlineMsg)
		}
		user.server.mapLock.Unlock()
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式: to|张三|msg
		// 1. 获取对方的用户名
		remoteName := strings.Split(msg, "|")[1]

		if remoteName == "" {
			user.SendMsg("ge shi err \n")
			return
		}

		// 2. 根据用户名, 得到对方 User 对象
		remoteUser, ok := user.server.OnlineMap[remoteName]

		if !ok {
			user.SendMsg("user is no \n")
			return
		}

		// 3. 发送消息
		content := strings.Split(msg, "|")[2]

		if content == "" {
			user.SendMsg("no msg, again send")
			return
		}

		remoteUser.SendMsg(user.Name + "say:" + content)
	} else {
		user.server.BroadCast(user, msg)
	}
}

// 创建一个用户 api
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 启动监听当前 user channel 消息的 goroutine
	go user.ListenMessage()

	return user
}
