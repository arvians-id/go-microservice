package model

import "github.com/arvians-id/go-microservice/adapter/pkg/product/pb"

type Product struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   int64  `json:"created_by"`
	User        *User  `json:"user"`
}

func (p *Product) ToProtoBuffer() *pb.Product {
	return &pb.Product{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		CreatedBy:   p.CreatedBy,
		User: &pb.UserService{
			Id:    p.User.Id,
			Name:  p.User.Name,
			Email: p.User.Email,
		},
	}
}
