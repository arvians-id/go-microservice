package service

import (
	"context"
	"github.com/arvians-id/go-microservice/user-service/internal/model"
	"github.com/arvians-id/go-microservice/user-service/internal/pb"
)

type UserService interface {
	ListUser(ctx context.Context) ([]*model.User, error)
	GetUser(ctx context.Context, id int64) (*model.User, error)
	CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*model.User, error)
	UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*model.User, error)
	DeleteUser(ctx context.Context, id int64) error
}
