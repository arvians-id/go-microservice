package service

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/product-service/internal/model"
	"github.com/arvians-id/go-microservice/product-service/internal/pb"
	"github.com/arvians-id/go-microservice/product-service/internal/repository"
	"github.com/arvians-id/go-microservice/product-service/util"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
	UserService       pb.UserServiceClient
	DB                *sql.DB
}

func NewProductService(productRepository repository.ProductRepository, userService pb.UserServiceClient, db *sql.DB) pb.ProductServiceServer {
	return &ProductService{
		ProductRepository: productRepository,
		UserService:       userService,
		DB:                db,
	}
}

func (p ProductService) ListProduct(ctx context.Context, empty *emptypb.Empty) (*pb.ListProductResponse, error) {
	tx, err := p.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	products, err := p.ProductRepository.ListProduct(ctx, tx)
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

func (p ProductService) GetProduct(ctx context.Context, req *pb.GetProductIdRequest) (*pb.GetProductResponse, error) {
	tx, err := p.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	product, err := p.ProductRepository.GetProduct(ctx, tx, req.Id)
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

func (p ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	tx, err := p.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	_, err = p.UserService.GetUser(ctx, &pb.GetUserIdRequest{
		Id: req.CreatedBy,
	})
	if err != nil {
		return nil, err
	}

	product, err := p.ProductRepository.CreateProduct(ctx, tx, &model.Product{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Status: http.StatusOK,
		Id:     product.Id,
	}, nil
}

func (p ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	tx, err := p.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	_, err = p.ProductRepository.GetProduct(ctx, tx, req.Id)
	if err != nil {
		return nil, err
	}

	_, err = p.ProductRepository.UpdateProduct(ctx, tx, &model.Product{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProductResponse{
		Status: http.StatusOK,
	}, nil
}

func (p ProductService) DeleteProduct(ctx context.Context, req *pb.GetProductIdRequest) (*pb.DeleteProductResponse, error) {
	tx, err := p.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	_, err = p.ProductRepository.GetProduct(ctx, tx, req.Id)
	if err != nil {
		return nil, err
	}

	err = p.ProductRepository.DeleteProduct(ctx, tx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteProductResponse{
		Status: http.StatusOK,
	}, nil
}
