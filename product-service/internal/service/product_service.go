package service

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/product-service/internal/model"
	"github.com/arvians-id/go-microservice/product-service/internal/pb"
	"github.com/arvians-id/go-microservice/product-service/internal/repository"
	"github.com/arvians-id/go-microservice/product-service/util"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	UserService       pb.UserServiceClient
	DB                *sql.DB
}

func NewProductService(productRepository repository.ProductRepository, userService pb.UserServiceClient, db *sql.DB) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		UserService:       userService,
		DB:                db,
	}
}

func (service *ProductServiceImpl) ListProduct(ctx context.Context) ([]*model.Product, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	products, err := service.ProductRepository.ListProduct(ctx, tx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (service *ProductServiceImpl) GetProduct(ctx context.Context, id int64) (*model.Product, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	product, err := service.ProductRepository.GetProduct(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductServiceImpl) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*model.Product, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	_, err = service.UserService.GetUser(ctx, &pb.GetUserIdRequest{
		Id: req.CreatedBy,
	})
	if err != nil {
		return nil, err
	}

	product, err := service.ProductRepository.CreateProduct(ctx, tx, &model.Product{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
	})
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductServiceImpl) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*model.Product, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	_, err = service.ProductRepository.GetProduct(ctx, tx, req.Id)
	if err != nil {
		return nil, err
	}

	product, err := service.ProductRepository.UpdateProduct(ctx, tx, &model.Product{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductServiceImpl) DeleteProduct(ctx context.Context, id int64) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer util.CommitOrRollback(tx)

	_, err = service.ProductRepository.GetProduct(ctx, tx, id)
	if err != nil {
		return err
	}

	err = service.ProductRepository.DeleteProduct(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}
