package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			p.ClientIP,
			p.TimeStamp.Format(time.RFC1123),
			p.Method,
			p.Path,
			p.Request.Proto,
			p.StatusCode,
			p.Latency,
			p.Request.UserAgent(),
			p.ErrorMessage,
		)
	}))

	r.GET("/log", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
	r.Run(":8080")
}
