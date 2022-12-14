package main

import (
	"database/sql"
	"fmt"
	"github.com/arvians-id/go-microservice/adapter/pkg/user/pb"
	"github.com/arvians-id/go-microservice/config"
	"github.com/arvians-id/go-microservice/services/user-service/repository"
	"github.com/arvians-id/go-microservice/services/user-service/usecase"
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

func NewInitializedServer(configuration *config.Config) (pb.UserServiceServer, error) {
	db, err := NewInitializedDatabase(configuration)
	if err != nil {
		return nil, err
	}

	// Main App
	userRepository := repository.NewUserRepository()
	userService := usecase.NewUserService(userRepository, db)

	return userService, nil
}

func main() {
	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	conn, err := net.Listen("tcp", configuration.UserSvcUrl)
	if err != nil {
		log.Fatalln("Failed at listening", err)
	}

	fmt.Println("User service is running on port", configuration.UserSvcUrl)

	userService, err := NewInitializedServer(configuration)
	if err != nil {
		log.Fatalln("Failed at initializing server", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userService)

	if err := grpcServer.Serve(conn); err != nil {
		log.Fatalln("Failed at serving", err)
	}
}
