package repository

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/model"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) ListUser(ctx context.Context, tx *sql.Tx) ([]*model.User, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (repository *UserRepositoryImpl) GetUser(ctx context.Context, tx *sql.Tx, id int64) (*model.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`
	row := tx.QueryRowContext(ctx, query, id)

	var user model.User
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) CreateUser(ctx context.Context, tx *sql.Tx, user *model.User) (*model.User, error) {
	var id int64
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := tx.QueryRowContext(ctx, query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return nil, err
	}

	user.Id = id
	return user, nil
}

func (repository *UserRepositoryImpl) UpdateUser(ctx context.Context, tx *sql.Tx, user *model.User) (*model.User, error) {
	query := `UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4`
	_, err := tx.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) DeleteUser(ctx context.Context, tx *sql.Tx, id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
