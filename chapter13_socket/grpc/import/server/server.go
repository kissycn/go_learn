package main

import (
	"google.golang.org/grpc"
	"ielee.com/go_learn/chapter13_socket/grpc/import/service"
	"log"
	"net"
)

func main() {
	var userService service.UserServiceServer = new(service.UserService)
	server := grpc.NewServer()

	service.RegisterUserServiceServer(server, userService)
	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	_ = server.Serve(listener)
}
