package product

import (
	"fmt"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/config"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/product/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitializeServiceClient(c *config.Config) pb.ProductServiceClient {
	connection, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewProductServiceClient(connection)
}
