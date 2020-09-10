package router

import (
	"net/http"

	"apiserver/handler"
	"apiserver/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 通过 g.Use() 来为每一个请求设置 Header
	// middleware
	// gin.Recovery()：在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，这时候为了不影响下一次请求的调用，需要通过 gin.Recovery()来恢复 API 服务器
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 router
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "")
	})
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", handler.HealthCheck)
		svcd.GET("/disk", handler.DiskCheck)
		svcd.GET("/cpu", handler.CPUCheck)
		svcd.GET("/ram", handler.RAMCheck)
	}
	return g
}
