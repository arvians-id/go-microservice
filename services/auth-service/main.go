package main

import (
	"database/sql"
	"fmt"
	"github.com/arvians-id/go-microservice/adapter/pkg/auth/pb"
	"github.com/arvians-id/go-microservice/config"
	"github.com/arvians-id/go-microservice/services/auth-service/repository"
	"github.com/arvians-id/go-microservice/services/auth-service/usecase"
	"github.com/arvians-id/go-microservice/util"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewInitializedDatabase(configuration *config.Config) (*sql.DB, error) {
	db, err := config.NewPostgresSQL(configuration)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewInitializedServer(configuration *config.Config) (pb.AuthServiceServer, error) {
	db, err := NewInitializedDatabase(configuration)
	if err != nil {
		return nil, err
	}

	// App
	jwtUtil := util.NewJwtWrapper(configuration)
	authRepository := repository.NewAuthRepository()
	authService := usecase.NewAuthService(authRepository, db, jwtUtil)

	return authService, nil
}

func main() {
	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	conn, err := net.Listen("tcp", configuration.AuthSvcUrl)
	if err != nil {
		log.Fatalln("Failed at listening", err)
	}

	fmt.Println("Auth service is running on port", configuration.AuthSvcUrl)

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
