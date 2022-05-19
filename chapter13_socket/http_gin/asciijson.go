package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		data := map[string]string{
			"lang": "Go语言",
			"tag":  "<br>",
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run()
}
