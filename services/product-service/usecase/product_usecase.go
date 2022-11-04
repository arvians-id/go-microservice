package usecase

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/adapter/pkg/product/pb"
	userpb "github.com/arvians-id/go-microservice/adapter/pkg/user/pb"
	"github.com/arvians-id/go-microservice/model"
	"github.com/arvians-id/go-microservice/services/product-service/repository"
	"github.com/arvians-id/go-microservice/util"
	"github.com/golang/protobuf/ptypes/empty"
	"net/http"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
	UserService       userpb.UserServiceClient
	StorageS3         util.StorageS3
	DB                *sql.DB
}

func NewProductService(productRepository repository.ProductRepository, userService userpb.UserServiceClient, storageS3 *util.StorageS3, db *sql.DB) pb.ProductServiceServer {
	return &ProductService{
		ProductRepository: productRepository,
		UserService:       userService,
		StorageS3:         *storageS3,
		DB:                db,
	}
}

func (p ProductService) ListProduct(empty *empty.Empty, stream pb.ProductService_ListProductServer) error {
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}
	defer util.CommitOrRollback(tx)

	products, err := p.ProductRepository.ListProduct(context.Background(), tx)
	if err != nil {
		return err
	}

	for _, product := range products {
		err := stream.Send(product.ToProtoBuffer())
		if err != nil {
			return err
		}
	}

	return nil
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

	_, err = p.UserService.GetUser(ctx, &userpb.GetUserIdRequest{
		Id: req.CreatedBy,
	})
	if err != nil {
		return nil, err
	}

	product, err := p.ProductRepository.CreateProduct(ctx, tx, &model.Product{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
		Image:       req.Image,
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

	product, err := p.ProductRepository.GetProduct(ctx, tx, req.Id)
	if err != nil {
		return nil, err
	}

	if req.Image != "" {
		_ = p.StorageS3.DeleteFromAWS(product.Image)
		product.Image = req.Image
	}

	_, err = p.ProductRepository.UpdateProduct(ctx, tx, &model.Product{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Image:       product.Image,
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
