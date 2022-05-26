package service

import "context"

type UserService struct {
}

func (UserService) GerUser(context.Context, *UserRequest) (*UserResponse, error) {
	return &UserResponse{User: &User{
		Name: "jerry",
		Addr: &Address{
			Addr:    "chengdu",
			Zipcode: "10086",
		},
	}}, nil
}
