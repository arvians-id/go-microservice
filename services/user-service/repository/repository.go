package repository

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/model"
)

type UserRepository interface {
	ListUser(ctx context.Context, tx *sql.Tx) ([]*model.User, error)
	GetUser(ctx context.Context, tx *sql.Tx, id int64) (*model.User, error)
	CreateUser(ctx context.Context, tx *sql.Tx, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, tx *sql.Tx, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, tx *sql.Tx, id int64) error
}
