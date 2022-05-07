package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()
	fmt.Println("listen ok")
	var i int
	for {
		time.Sleep(time.Second * 10)
		if _, err := l.Accept(); err != nil {
			fmt.Println(err)
			break
		}
		i++
		fmt.Println("accept new conn:%d", i)
	}
}
