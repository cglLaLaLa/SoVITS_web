package server

import "github.com/gin-gonic/gin"

func RunHttpServer(serviceName string, wrapper func(router *gin.Engine)) {
	//todo 后期配置化，不要写死
	addr := "127.0.0.1:14888"
	if addr == "" {
		panic("empty http address")
	}
	RunHttpServerOnAddr(addr, wrapper)
}

func RunHttpServerOnAddr(addr string, wrapper func(router *gin.Engine)) {
	apiRouter := gin.New()
	wrapper(apiRouter)
	apiRouter.Group("/api")
	apiRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	if err := apiRouter.Run(addr); err != nil {
		panic(err)
	}
}
