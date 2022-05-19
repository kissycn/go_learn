package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/:id/:name", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindUri(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}
		fmt.Println("id=", c.Param("id"), "name=", c.Query("name"))
		c.JSON(http.StatusOK, user)
	})
	r.Run()
}
