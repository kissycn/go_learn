package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.Writer(f)
	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	r.Run()
}
