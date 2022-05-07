package chapter13_socket

import (
	"fmt"
	"net"
	"testing"
)

func TestSocketServer(t *testing.T) {
	s, err := net.Listen("tcp", "8080")
	if nil != err {
		fmt.Println(err)
		return
	}

	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}

		go HandConn(c)
	}
}

func TestSocketClient(t *testing.T) {

}

func HandConn(c net.Conn) {
	defer c.Close()
}
