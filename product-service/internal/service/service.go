package service

import (
	"context"
	"github.com/arvians-id/go-microservice/product-service/internal/model"
	"github.com/arvians-id/go-microservice/product-service/internal/pb"
)

type ProductService interface {
	ListProduct(ctx context.Context) ([]*model.Product, error)
	GetProduct(ctx context.Context, id int64) (*model.Product, error)
	CreateProduct(ctx context.Context, request *pb.CreateProductRequest) (*model.Product, error)
	UpdateProduct(ctx context.Context, request *pb.UpdateProductRequest) (*model.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
}
