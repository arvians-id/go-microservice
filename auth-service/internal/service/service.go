package service

import (
	"context"
	"github.com/arvians-id/go-microservice/auth-service/internal/pb"
)

type AuthService interface {
	Login(ctx context.Context, req *pb.LoginRequest) (string, error)
	Register(ctx context.Context, req *pb.RegisterRequest) error
	Validate(ctx context.Context, req *pb.ValidateRequest) (int64, error)
}
