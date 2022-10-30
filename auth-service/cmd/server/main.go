package main

import (
	"database/sql"
	"fmt"
	"github.com/arvians-id/go-microservice/auth-service/internal/config"
	"github.com/arvians-id/go-microservice/auth-service/internal/pb"
	"github.com/arvians-id/go-microservice/auth-service/internal/repository"
	"github.com/arvians-id/go-microservice/auth-service/internal/service"
	"github.com/arvians-id/go-microservice/auth-service/util"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewInitializedDatabase(configuration config.Config) (*sql.DB, error) {
	db, err := config.NewPostgresSQL(configuration)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewInitializedServer(configuration config.Config) (pb.AuthServiceServer, error) {
	db, err := NewInitializedDatabase(configuration)
	if err != nil {
		return nil, err
	}

	// App
	jwtUtil := util.NewJwtWrapper(configuration)
	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, db, jwtUtil)

	return authService, nil
}

func main() {
	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	conn, err := net.Listen("tcp", configuration.Port)
	if err != nil {
		log.Fatalln("Failed at listening", err)
	}

	fmt.Println("Auth service is running on port", configuration.Port)

	authService, err := NewInitializedServer(configuration)
	if err != nil {
		log.Fatalln("Failed at initializing server", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authService)

	if err := grpcServer.Serve(conn); err != nil {
		log.Fatalln("Failed at serving", err)
	}
}
