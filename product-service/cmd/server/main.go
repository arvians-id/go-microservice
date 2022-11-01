package main

import (
	"database/sql"
	"fmt"
	"github.com/arvians-id/go-microservice/product-service/internal/client"
	"github.com/arvians-id/go-microservice/product-service/internal/config"
	"github.com/arvians-id/go-microservice/product-service/internal/pb"
	"github.com/arvians-id/go-microservice/product-service/internal/repository"
	"github.com/arvians-id/go-microservice/product-service/internal/server"
	"github.com/arvians-id/go-microservice/product-service/internal/service"
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

func NewInitializedServer(configuration *config.Config) (pb.ProductServiceServer, error) {
	db, err := NewInitializedDatabase(configuration)
	if err != nil {
		return nil, err
	}

	// Another service
	userService := client.InitializeUserServiceClient(configuration)

	// Main App
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, userService, db)
	productServer := server.NewProductServer(productService)

	return productServer, nil
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

	fmt.Println("Product service is running on port", configuration.Port)

	productService, err := NewInitializedServer(configuration)
	if err != nil {
		log.Fatalln("Failed at initializing server", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, productService)

	if err := grpcServer.Serve(conn); err != nil {
		log.Fatalln("Failed at serving", err)
	}
}
