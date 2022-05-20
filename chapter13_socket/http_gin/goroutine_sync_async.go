package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.JSON(http.StatusOK, "ok")
	})

	r.GET("/async", func(c *gin.Context) {
		cp := c.Copy()

		go func() {
			time.Sleep(5 * time.Second)
			log.Print(cp.Request.URL.Path, " down!")
		}()

		c.JSON(http.StatusOK, "down!")
	})

	r.Run()
}
