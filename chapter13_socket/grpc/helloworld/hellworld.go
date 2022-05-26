package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"ielee.com/go_learn/chapter13_socket/grpc/helloworld/service"
)

func main() {
	user := &service.User{
		Username: "mszlu",
	}
	//转换为protobuf
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser.String())
}
