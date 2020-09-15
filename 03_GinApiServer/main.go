package main

import (
	"apiserver/config"
	"apiserver/router"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

func main() {
	// 把用户传递的命令行参数解析为对应变量的值
	pflag.Parse()

	// init config
	// 增加了 config.Init(*cfg) 调用，用来初始化配置
	if err := config.Init(*cfg); err != nil {
		// panic内置函数停止当前goroutine的正常执行
		panic(err)
	}

	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...,
	)

	// Ping the server to make sure the router is working
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()
	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	// log.Printf(http.ListenAndServe(":8080", g).Error())
	g.Run(":8080")
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err != nil {
			return err
		}
		if resp.StatusCode == 200 {
			return nil
		}
		log.Print("waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}
