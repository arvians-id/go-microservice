package util

import (
	"github.com/arvians-id/go-microservice/auth-service/internal/config"
	"github.com/arvians-id/go-microservice/auth-service/internal/model"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaims struct {
	jwt.StandardClaims
	Id    int64
	Email string
	Name  string
}

func NewJwtWrapper(config *config.Config) *JwtWrapper {
	return &JwtWrapper{
		SecretKey:       config.JwtSecretKey,
		Issuer:          "auth-service",
		ExpirationHours: 24 * 365,
	}
}

func (j *JwtWrapper) GenerateToken(user *model.User) (string, error) {
	claims := &JwtClaims{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.SecretKey))
}

func (j *JwtWrapper) ValidateToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
