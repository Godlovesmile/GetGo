package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("listen", err)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	// 解析url传递的参数，对于POST则解析响应包的主体（request body）
	r.ParseForm()
	// 注意: 如果没有调用ParseForm方法，下面无法获取表单的数据
	// 这些信息是输出到服务器端的打印信息
	fmt.Println(r.Form)
	fmt.Println("scheme", r.URL.Scheme)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	// 这个写入到w的是输出到客户端的
	fmt.Fprintf(w, "Hello world!")
}

// 防止多次递交表单, 使用 token; 这个值我们通过MD5(时间戳)来获取惟一值，然后我们把这个值存储到服务器端(session来控制)
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		toke := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, toke)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		fmt.Println("token: ", token)
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}

// 上传文件
func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")

		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
