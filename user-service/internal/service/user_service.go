package service

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/user-service/internal/model"
	"github.com/arvians-id/go-microservice/user-service/internal/pb"
	"github.com/arvians-id/go-microservice/user-service/internal/repository"
	"github.com/arvians-id/go-microservice/user-service/util"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (service *UserServiceImpl) ListUser(ctx context.Context) ([]*model.User, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	users, err := service.UserRepository.ListUser(ctx, tx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *UserServiceImpl) GetUser(ctx context.Context, id int64) (*model.User, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	user, err := service.UserRepository.GetUser(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserServiceImpl) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*model.User, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	var user model.User
	user.Name = request.Name
	user.Email = request.Email
	user.Password = util.HashPassword(request.Password)
	userResponse, err := service.UserRepository.CreateUser(ctx, tx, &user)
	if err != nil {
		return nil, err
	}

	return userResponse, nil
}

func (service *UserServiceImpl) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*model.User, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	_, err = service.UserRepository.GetUser(ctx, tx, request.Id)
	if err != nil {
		return nil, err
	}

	var user model.User
	user.Id = request.Id
	user.Name = request.Name
	user.Email = request.Email
	user.Password = util.HashPassword(request.Password)
	userResponse, err := service.UserRepository.UpdateUser(ctx, tx, &user)
	if err != nil {
		return nil, err
	}

	return userResponse, nil
}

func (service *UserServiceImpl) DeleteUser(ctx context.Context, id int64) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer util.CommitOrRollback(tx)

	_, err = service.UserRepository.GetUser(ctx, tx, id)
	if err != nil {
		return err
	}

	err = service.UserRepository.DeleteUser(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}
