package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main() {
	s1 := http.Server{
		Addr:    ":8080",
		Handler: router01(),
	}

	s2 := http.Server{
		Addr:    ":8090",
		Handler: router02(),
	}

	group := errgroup.Group{}
	group.Go(func() error {
		return s1.ListenAndServe()
	})
	group.Go(func() error {
		return s2.ListenAndServe()
	})

	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}

func router01() http.Handler {
	e := gin.Default()
	e.GET("/s1", func(c *gin.Context) {
		c.JSON(http.StatusOK, "s1")
	})
	return e
}

func router02() http.Handler {
	e := gin.Default()
	e.GET("/s2", func(c *gin.Context) {
		c.JSON(http.StatusOK, "s2")
	})
	return e
}
