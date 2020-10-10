package main

import (
	"fmt"
	"net/http"
	"sessionmod/session"
	"text/template"
)

/*
session的基本原理是由服务器为每个会话维护一份信息数据，客户端和服务端依靠一个全局唯一的标识来访问这份数据，以达到交互的目的;
	1. 生成全局唯一标识符(sessionid)
	2. 开辟数据存储空间
	3. 将session的全局唯一标识符发送客户端

*/
var globalSessions *session.Manager
var err error

func init() {
	fmt.Println("--- init ---")
	globalSessions, err = session.NewManager("memory", "gosessionid", 3600)

	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("--- main ---")
	http.HandleFunc("/login", login)
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()

	if r.Method == "GET" {
		t, err := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
		if err != nil {
			panic(err)
		}
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}
