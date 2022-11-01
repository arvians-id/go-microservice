package server

import (
	"context"
	"github.com/arvians-id/go-microservice/product-service/internal/pb"
	"github.com/arvians-id/go-microservice/product-service/internal/service"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

type ProductServer struct {
	ProductService service.ProductService
}

func NewProductServer(productService service.ProductService) pb.ProductServiceServer {
	return &ProductServer{
		ProductService: productService,
	}
}

func (server *ProductServer) ListProduct(ctx context.Context, empty *emptypb.Empty) (*pb.ListProductResponse, error) {
	products, err := server.ProductService.ListProduct(ctx)
	if err != nil {
		return nil, err
	}

	var productResponse []*pb.Product
	for _, product := range products {
		productResponse = append(productResponse, &pb.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			CreatedBy:   product.CreatedBy,
			User: &pb.UserService{
				Id:    product.User.Id,
				Name:  product.User.Name,
				Email: product.User.Email,
			},
		})
	}

	return &pb.ListProductResponse{
		Data: productResponse,
	}, nil
}

func (server *ProductServer) GetProduct(ctx context.Context, req *pb.GetProductIdRequest) (*pb.GetProductResponse, error) {
	product, err := server.ProductService.GetProduct(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetProductResponse{
		Data: &pb.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			CreatedBy:   product.CreatedBy,
			User: &pb.UserService{
				Id:    product.User.Id,
				Name:  product.User.Name,
				Email: product.User.Email,
			},
		},
	}, nil
}

func (server *ProductServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := server.ProductService.CreateProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Status: http.StatusOK,
		Id:     product.Id,
	}, nil
}

func (server *ProductServer) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	_, err := server.ProductService.UpdateProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProductResponse{
		Status: http.StatusOK,
	}, nil
}

func (server *ProductServer) DeleteProduct(ctx context.Context, req *pb.GetProductIdRequest) (*pb.DeleteProductResponse, error) {
	err := server.ProductService.DeleteProduct(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteProductResponse{
		Status: http.StatusOK,
	}, nil
}
