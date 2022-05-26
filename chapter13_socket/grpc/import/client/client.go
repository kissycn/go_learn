package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"ielee.com/go_learn/chapter13_socket/grpc/import/service"
	"log"
)

func main() {
	// 1. 新建连接，端口是服务端开放的8002端口
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userService := service.NewUserServiceClient(conn)
	resp, err := userService.GerUser(context.Background(), &service.UserRequest{
		UserId: 1,
	})
	if err != nil {
		log.Fatal("调用gRPC方法错误: ", err)
	}

	fmt.Println("rpc call success resp:", resp.User)
}
