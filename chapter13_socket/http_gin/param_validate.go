package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Employee struct {
	UserName string `json:"userName" binding:"required"`
	NickName string `json:"nickName" binding:"required"`
	Age      int    `json:"age" binding:"required,gte=1,lte=120"`
	Email    string `json:"email" binding:"required,email"`
}

func main() {
	r := gin.Default()
	r.POST("/reg", func(c *gin.Context) {
		var e Employee
		err := c.ShouldBindJSON(&e)
		if nil != err {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusOK, "success")
	})
	r.Run()
}
