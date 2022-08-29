package main

import "net"

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
	for {
		msg := <-user.C
		user.conn.Write([]byte(msg + "\n"))
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

// 用户发消息
func (user *User) UserMsg(msg string) {
	user.server.BroadCast(user, msg)
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
