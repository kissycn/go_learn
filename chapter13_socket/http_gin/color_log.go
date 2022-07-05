package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// gin.DisableConsoleColor()
	// by default
	//gin.ForceConsoleColor()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/color", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
	r.Run()
}
