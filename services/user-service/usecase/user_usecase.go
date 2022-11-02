package usecase

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/adapter/pkg/user/pb"
	"github.com/arvians-id/go-microservice/model"
	"github.com/arvians-id/go-microservice/services/user-service/repository"
	util2 "github.com/arvians-id/go-microservice/util"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

type UserService struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB) pb.UserServiceServer {
	return &UserService{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (u UserService) ListUser(ctx context.Context, empty *emptypb.Empty) (*pb.ListUserResponse, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util2.CommitOrRollback(tx)

	users, err := u.UserRepository.ListUser(ctx, tx)
	if err != nil {
		return nil, err
	}

	var userResponse []*pb.User
	for _, user := range users {
		userResponse = append(userResponse, &pb.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return &pb.ListUserResponse{
		Data: userResponse,
	}, nil
}

func (u UserService) GetUser(ctx context.Context, request *pb.GetUserIdRequest) (*pb.GetUserResponse, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util2.CommitOrRollback(tx)

	user, err := u.UserRepository.GetUser(ctx, tx, request.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Data: &pb.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (u UserService) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util2.CommitOrRollback(tx)

	var user model.User
	user.Name = request.Name
	user.Email = request.Email
	user.Password = util2.HashPassword(request.Password)
	userResponse, err := u.UserRepository.CreateUser(ctx, tx, &user)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Status: http.StatusOK,
		Id:     userResponse.Id,
	}, nil
}

func (u UserService) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util2.CommitOrRollback(tx)

	_, err = u.UserRepository.GetUser(ctx, tx, request.Id)
	if err != nil {
		return nil, err
	}

	var user model.User
	user.Id = request.Id
	user.Name = request.Name
	user.Email = request.Email
	user.Password = util2.HashPassword(request.Password)
	_, err = u.UserRepository.UpdateUser(ctx, tx, &user)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{
		Status: http.StatusOK,
	}, nil
}

func (u UserService) DeleteUser(ctx context.Context, request *pb.GetUserIdRequest) (*pb.DeleteUserResponse, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util2.CommitOrRollback(tx)

	_, err = u.UserRepository.GetUser(ctx, tx, request.Id)
	if err != nil {
		return nil, err
	}

	err = u.UserRepository.DeleteUser(ctx, tx, request.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{
		Status: http.StatusOK,
	}, nil
}
