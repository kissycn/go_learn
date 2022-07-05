package main

import (
	"fmt"
	"google.golang.org/grpc"
	"ielee.com/go_learn/chapter13_socket/grpc/cs/service"
	"log"
	"net"
)

func main() {
	var p service.ProdServiceServer = new(service.ProductService)

	server := grpc.NewServer()
	service.RegisterProdServiceServer(server, p)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	_ = server.Serve(listener)

	fmt.Println("")
}
