package repository

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/product-service/internal/model"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) ListProduct(ctx context.Context, tx *sql.Tx) ([]*model.Product, error) {
	query := `SELECT p.*, u.id, u.name, u.email FROM products p LEFT JOIN users u ON p.created_by = u.id`
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var products []*model.Product
	for rows.Next() {
		product := model.Product{
			User: &model.User{},
		}
		rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.CreatedBy,
			&product.User.Id,
			&product.User.Name,
			&product.User.Email,
		)

		products = append(products, &product)
	}

	return products, nil
}

func (repository *ProductRepositoryImpl) GetProduct(ctx context.Context, tx *sql.Tx, id int64) (*model.Product, error) {
	query := `SELECT p.*, u.id, u.name, u.email FROM products p LEFT JOIN users u ON p.created_by = u.id WHERE p.id = $1`
	row := tx.QueryRowContext(ctx, query, id)

	product := model.Product{
		User: &model.User{},
	}
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.CreatedBy,
		&product.User.Id,
		&product.User.Name,
		&product.User.Email,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repository *ProductRepositoryImpl) CreateProduct(ctx context.Context, tx *sql.Tx, product *model.Product) (*model.Product, error) {
	var id int64
	query := `INSERT INTO products (name, description, created_by) VALUES ($1, $2, $3) RETURNING id`
	row := tx.QueryRowContext(ctx, query, product.Name, product.Description, product.CreatedBy)
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	product.Id = id

	return product, nil
}

func (repository *ProductRepositoryImpl) UpdateProduct(ctx context.Context, tx *sql.Tx, product *model.Product) (*model.Product, error) {
	query := `UPDATE products SET name = $1, description = $2 WHERE id = $3`
	_, err := tx.ExecContext(ctx, query, product.Name, product.Description, product.Id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) DeleteProduct(ctx context.Context, tx *sql.Tx, id int64) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
