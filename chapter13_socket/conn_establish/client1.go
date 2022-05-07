package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	log.Println("begin dial...")
	c, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()
	fmt.Printf("dial ok")
}
