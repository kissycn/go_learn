package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/encrypt", func(c *gin.Context) {
		c.String(200, "pong")
	})

	autotls.Run(r, "example1.com", "example2.com")
}
