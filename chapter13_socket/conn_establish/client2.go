package main

import (
	"fmt"
	"net"
)

func main() {
	var conns []net.Conn
	for i := 0; i < 1000; i++ {
		conn := establishConn(i)
		if nil != conn {
			conns = append(conns, conn)
		}
	}
}

func establishConn(i int) net.Conn {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(i, ":connect to server")
	return conn
}
