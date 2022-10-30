package repository

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/auth-service/internal/model"
)

type AuthRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user *model.User) (*model.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*model.User, error)
}
