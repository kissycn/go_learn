package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name     string `form:"name"`
	NickName string `form:"nickName"`
	Address  string `form:"addr"`
}

func main() {
	r := gin.Default()
	r.GET("/query", func(c *gin.Context) {
		var person Person
		c.ShouldBind(&person)
		fmt.Println(person)
		c.JSON(http.StatusOK, person)
	})

	r.GET("/query1", func(c *gin.Context) {
		var p Person
		p.Name = c.Query("name")
		p.NickName = c.Query("nickName")
		c.JSON(http.StatusOK, p)
	})
	r.Run()
}
