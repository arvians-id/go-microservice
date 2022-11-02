package auth

import (
	"fmt"
	"github.com/arvians-id/go-microservice/adapter/pkg/auth/pb"
	"github.com/arvians-id/go-microservice/adapter/pkg/auth/request"
	"github.com/arvians-id/go-microservice/config"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
)

type ServiceClient struct {
	AuthClient pb.AuthServiceClient
}

func NewAuthServiceClient(c *config.Config) pb.AuthServiceClient {
	connection, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(connection)
}

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		AuthClient: NewAuthServiceClient(c),
	}

	routeGroup := r.Group("/auth")
	routeGroup.POST("/login", svc.Login)
	routeGroup.POST("/register", svc.Register)

	return svc
}

func (client *ServiceClient) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := client.AuthClient.Login(ctx, &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}

func (client *ServiceClient) Register(ctx *gin.Context) {
	var req request.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := client.AuthClient.Register(ctx, &pb.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
