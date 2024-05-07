package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, "hello world")
	})
	r.Run("0.0.0.0:8899") // 监听并在 0.0.0.0:8899 上启动服务
}

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build