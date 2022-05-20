package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	// 默认路由输出格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	r.GET("/world", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	r.Run()
}
