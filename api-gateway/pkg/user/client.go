package user

import (
	"fmt"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/config"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/user/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.UserServiceClient
}

func InitializeServiceClient(c *config.Config) pb.UserServiceClient {
	connection, err := grpc.Dial(c.UserSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewUserServiceClient(connection)
}
