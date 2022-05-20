package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		var a1 []int
		a1[100] = 1
		c.JSON(200, "pong")
	})
	r.Run()
}
