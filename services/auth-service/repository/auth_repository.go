package repository

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/model"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user *model.User) (*model.User, error) {
	var id int64
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	row := tx.QueryRowContext(ctx, query, user.Name, user.Email, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	user.Id = id

	return user, nil
}

func (repository *AuthRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE email = $1`
	row := tx.QueryRowContext(ctx, query, email)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
