package jsonmodule

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
)

// Server info
type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

// Serverslice info
// struct tag 重新定义 json 数据的 key, 首字母小写
type Serverslice struct {
	Servers []Server `json:"servers"`
}

// HandleJSONData info
func HandleJSONData() {
	// 1. 已知 json 结构, 先定义结构体
	var s Serverslice
	str := `{"servers": [{"serverName": "Beijing", "serverIP": "127.0.0.1"}, {"serverName": "Anhui", "serverIP": "127.0.0.1"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	s.Servers = append(s.Servers, Server{ServerName: "Chongqing_VPN", ServerIP: "127.0.0.1"})
	fmt.Println(s)

	// 输出 JSON 数据串
	c, error := json.Marshal(s)
	if error != nil {
		fmt.Println("json err:", error)
	}
	fmt.Println(string(c))

	// 2. 未知 json 数据结构, 官方推荐, 通过 interface, 断言访问
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)

	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	// 访问数据, 通过断言访问
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	// 3. 第三方工具包, go-simplejson 处理 json 数据
	js, err := simplejson.NewJson([]byte(`{
			"test": {
				"array": [1, "2", 3],
				"int": 10,
				"float": 5.150,
				"bignum": 9223372036854775807,
				"string": "simplejson",
				"bool": true
			}
		}`))
	if err != nil {
		panic(err)
	}
	fmt.Println(js)

	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()
	fmt.Println(arr, i, ms)
}
