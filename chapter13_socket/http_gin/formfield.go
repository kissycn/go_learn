package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StructA struct {
	FieldA string `form:"field_a"`
	FieldB string `form:"field_b"`
}

func main() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		var a StructA
		c.Bind(&a)
		c.JSON(http.StatusOK, a)
	})
	r.Run()
}
