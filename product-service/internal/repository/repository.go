package repository

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/product-service/internal/model"
)

type ProductRepository interface {
	ListProduct(ctx context.Context, tx *sql.Tx) ([]*model.Product, error)
	GetProduct(ctx context.Context, tx *sql.Tx, id int64) (*model.Product, error)
	CreateProduct(ctx context.Context, tx *sql.Tx, product *model.Product) (*model.Product, error)
	UpdateProduct(ctx context.Context, tx *sql.Tx, product *model.Product) (*model.Product, error)
	DeleteProduct(ctx context.Context, tx *sql.Tx, id int64) error
}
