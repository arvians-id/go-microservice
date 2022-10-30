package service

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/auth-service/internal/model"
	"github.com/arvians-id/go-microservice/auth-service/internal/pb"
	"github.com/arvians-id/go-microservice/auth-service/internal/repository"
	"github.com/arvians-id/go-microservice/auth-service/util"
	"net/http"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *sql.DB
	Jwt            util.JwtWrapper
}

func NewAuthService(authRepository repository.AuthRepository, db *sql.DB, jwt *util.JwtWrapper) pb.AuthServiceServer {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		DB:             db,
		Jwt:            *jwt,
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	responseLogin, err := s.AuthRepository.FindByEmail(ctx, tx, req.Email)
	if err != nil {
		return nil, err
	}

	isMatch := util.CheckPasswordHash(req.Password, responseLogin.Password)
	if !isMatch {
		return nil, err
	}

	token, err := s.Jwt.GenerateToken(responseLogin)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *AuthServiceImpl) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	var user model.User
	user.Name = req.Name
	user.Email = req.Email
	user.Password = util.HashPassword(req.Password)
	_, err = s.AuthRepository.Register(ctx, tx, &user)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *AuthServiceImpl) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer util.CommitOrRollback(tx)

	claims, err := s.Jwt.ValidateToken(req.Token)
	if err != nil {
		return nil, err
	}

	responseUser, err := s.AuthRepository.FindByEmail(ctx, tx, claims.Email)
	if err != nil {
		return nil, err
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: responseUser.Id,
	}, nil
}
