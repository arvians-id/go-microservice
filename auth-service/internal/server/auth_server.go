package server

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/auth-service/internal/pb"
	"github.com/arvians-id/go-microservice/auth-service/internal/service"
	"github.com/arvians-id/go-microservice/auth-service/util"
	"net/http"
)

type AuthServer struct {
	AuthService service.AuthService
	DB          *sql.DB
	Jwt         util.JwtWrapper
}

func NewAuthServer(authService service.AuthService, jwt *util.JwtWrapper) pb.AuthServiceServer {
	return &AuthServer{
		AuthService: authService,
		Jwt:         *jwt,
	}
}

func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := server.AuthService.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (server *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	err := server.AuthService.Register(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Status: http.StatusOK,
	}, nil
}

func (server *AuthServer) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	idUser, err := server.AuthService.Validate(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: idUser,
	}, nil
}
