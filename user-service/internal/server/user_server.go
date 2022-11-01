package server

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/user-service/internal/pb"
	"github.com/arvians-id/go-microservice/user-service/internal/service"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

type UserServer struct {
	UserService service.UserService
	DB          *sql.DB
}

func NewUserServer(userService service.UserService) pb.UserServiceServer {
	return &UserServer{
		UserService: userService,
	}
}

func (server *UserServer) ListUser(ctx context.Context, empty *emptypb.Empty) (*pb.ListUserResponse, error) {
	users, err := server.UserService.ListUser(ctx)
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

func (server *UserServer) GetUser(ctx context.Context, request *pb.GetUserIdRequest) (*pb.GetUserResponse, error) {
	user, err := server.UserService.GetUser(ctx, request.Id)
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

func (server *UserServer) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	userResponse, err := server.UserService.CreateUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Status: http.StatusOK,
		Id:     userResponse.Id,
	}, nil
}

func (server *UserServer) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	_, err := server.UserService.UpdateUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{
		Status: http.StatusOK,
	}, nil
}

func (server *UserServer) DeleteUser(ctx context.Context, request *pb.GetUserIdRequest) (*pb.DeleteUserResponse, error) {
	err := server.UserService.DeleteUser(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{
		Status: http.StatusOK,
	}, nil
}
