package client

import (
	"context"
	"fmt"
	"github.com/arvians-id/go-microservice/product-service/internal/config"
	"github.com/arvians-id/go-microservice/product-service/internal/pb"
	"google.golang.org/grpc"
)

type UserServiceClient struct {
	Client pb.UserServiceClient
}

func InitializeUserServiceClient(c *config.Config) pb.UserServiceClient {
	connection, err := grpc.Dial(c.UserSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewUserServiceClient(connection)
}

func (u *UserServiceClient) GetUser(id int64) (*pb.GetUserResponse, error) {
	return u.Client.GetUser(context.Background(), &pb.GetUserIdRequest{
		Id: id,
	})
}
