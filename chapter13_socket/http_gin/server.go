package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"sync"
	"time"
)

type Product struct {
	Username    string    `json:"username" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Price       int       `json:"price" binding:"gte=0"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type productHandler struct {
	sync.Mutex
	products map[string]Product
}

func newProductHandler() *productHandler {
	return &productHandler{
		products: make(map[string]Product),
	}
}

func (p *productHandler) Create(c *gin.Context) {
	p.Mutex.Lock()
	p.Mutex.Unlock()

	// 1. parser param
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "param error"})
		return
	}

	// 2. validate param
	if _, ok := p.products[product.Name]; ok {
		c.JSON(http.StatusBadRequest, gin.H{"err": "already exists"})
		return
	}

	product.CreatedAt = time.Now()
	p.products[product.Name] = product

	c.JSON(http.StatusOK, product)
}

func (p *productHandler) Get(c *gin.Context) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	product, ok := p.products[c.Param("name")]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"err": "product not fount by:" + c.Param("name")})
		return
	}

	c.JSON(http.StatusOK, product)
}

func router() http.Handler {
	router := gin.Default()
	handler := newProductHandler()

	v1 := router.Group("/v1")
	{
		productV1 := v1.Group("/product").Use(gin.BasicAuth(gin.Accounts{"foo": "bar"}))
		productV1.GET(":name", handler.Get)
		productV1.POST("", handler.Create)
	}

	return router
}

func main() {
	var eg errgroup.Group

	insecureServer := http.Server{
		Addr:         ":8080",
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	secureServer := http.Server{
		Addr:         ":8443",
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	eg.Go(func() error {
		err := insecureServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	eg.Go(func() error {
		err := secureServer.ListenAndServeTLS("./server.pem", "./server.key")
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}
