package user

import (
	"fmt"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/auth"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/config"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/user/pb"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/user/request"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

type ServiceClient struct {
	UserClient pb.UserServiceClient
}

func NewUserServiceClient(c *config.Config) pb.UserServiceClient {
	connection, err := grpc.Dial(c.UserSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewUserServiceClient(connection)
}

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	authMiddleware := auth.NewMiddlewareAuthConfig(authSvc)
	svc := &ServiceClient{
		UserClient: NewUserServiceClient(c),
	}

	routeGroup := r.Group("/user", authMiddleware.AuthRequired)
	routeGroup.GET("/", svc.ListUser)
	routeGroup.GET("/:id", svc.GetUser)
	routeGroup.POST("/", svc.CreateUser)
	routeGroup.PUT("/:id", svc.UpdateUser)
	routeGroup.DELETE("/:id", svc.DeleteUser)

	return svc
}

func (client *ServiceClient) ListUser(ctx *gin.Context) {
	response, err := client.UserClient.ListUser(ctx, new(empty.Empty))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}

func (client *ServiceClient) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	response, err := client.UserClient.GetUser(ctx, &pb.GetUserIdRequest{
		Id: int64(id),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}

func (client *ServiceClient) CreateUser(ctx *gin.Context) {
	var req request.CreateUserRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := client.UserClient.CreateUser(ctx, &pb.CreateUserRequest{
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

func (client *ServiceClient) UpdateUser(ctx *gin.Context) {
	var req request.UpdateUserRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	response, err := client.UserClient.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:       id,
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

func (client *ServiceClient) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	response, err := client.UserClient.DeleteUser(ctx, &pb.GetUserIdRequest{
		Id: id,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
