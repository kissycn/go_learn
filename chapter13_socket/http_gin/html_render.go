package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Index Page"})
	})

	r.GET("/posts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts.tmpl", gin.H{"title": "Posts Page"})
	})

	r.Run()
}
