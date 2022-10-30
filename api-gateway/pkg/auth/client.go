package auth

import (
	"fmt"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/auth/pb"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitializeServiceClient(c *config.Config) pb.AuthServiceClient {
	connection, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(connection)
}
