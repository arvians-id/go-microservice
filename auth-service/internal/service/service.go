package service

import (
	"context"
	"github.com/arvians-id/go-microservice/auth-service/internal/pb"
)

type AuthService interface {
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error)
}
