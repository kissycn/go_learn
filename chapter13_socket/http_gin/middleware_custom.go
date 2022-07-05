package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func LoggerHandFunc(c *gin.Context) {
	t := time.Now()

	// before
	c.Set("ReqID", "12345")

	c.Next()

	// after
	latency := time.Since(t)
	fmt.Println("req latency=", latency)
}

func main() {
	r := gin.Default()

	r.Use(LoggerHandFunc)

	r.GET("/ping", func(c *gin.Context) {
		rid := c.GetString("ReqID")
		fmt.Println(rid)
		c.JSON(http.StatusOK, "pong")
	})
	r.Run()
}
