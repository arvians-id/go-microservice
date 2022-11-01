package service

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-microservice/auth-service/internal/model"
	"github.com/arvians-id/go-microservice/auth-service/internal/pb"
	"github.com/arvians-id/go-microservice/auth-service/internal/repository"
	"github.com/arvians-id/go-microservice/auth-service/util"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *sql.DB
	Jwt            util.JwtWrapper
}

func NewAuthService(authRepository repository.AuthRepository, db *sql.DB, jwt *util.JwtWrapper) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		DB:             db,
		Jwt:            *jwt,
	}
}

func (service *AuthServiceImpl) Login(ctx context.Context, req *pb.LoginRequest) (string, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return "", err
	}
	defer util.CommitOrRollback(tx)

	responseLogin, err := service.AuthRepository.FindByEmail(ctx, tx, req.Email)
	if err != nil {
		return "", err
	}

	isMatch := util.CheckPasswordHash(req.Password, responseLogin.Password)
	if !isMatch {
		return "", err
	}

	token, err := service.Jwt.GenerateToken(responseLogin)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *AuthServiceImpl) Register(ctx context.Context, req *pb.RegisterRequest) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer util.CommitOrRollback(tx)

	var user model.User
	user.Name = req.Name
	user.Email = req.Email
	user.Password = util.HashPassword(req.Password)
	_, err = service.AuthRepository.Register(ctx, tx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (service *AuthServiceImpl) Validate(ctx context.Context, req *pb.ValidateRequest) (int64, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer util.CommitOrRollback(tx)

	claims, err := service.Jwt.ValidateToken(req.Token)
	if err != nil {
		return 0, err
	}

	responseUser, err := service.AuthRepository.FindByEmail(ctx, tx, claims.Email)
	if err != nil {
		return 0, err
	}

	return responseUser.Id, nil
}
