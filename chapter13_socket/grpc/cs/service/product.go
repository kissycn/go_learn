package service

import (
	"context"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type ProductService struct {
}

func (ProductService) GetProductStock(context.Context, *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		ProdStock:     9527,
	}, nil
}
