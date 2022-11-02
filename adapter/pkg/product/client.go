package product

import (
	"fmt"
	"github.com/arvians-id/go-microservice/adapter/pkg/auth"
	"github.com/arvians-id/go-microservice/adapter/pkg/product/pb"
	"github.com/arvians-id/go-microservice/adapter/pkg/product/request"
	"github.com/arvians-id/go-microservice/config"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"io"
	"net/http"
	"strconv"
)

type ServiceClient struct {
	ProductClient pb.ProductServiceClient
}

func NewProductServiceClient(c *config.Config) pb.ProductServiceClient {
	connection, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewProductServiceClient(connection)
}

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	authMiddleware := auth.NewMiddlewareAuthConfig(authSvc)
	svc := &ServiceClient{
		ProductClient: NewProductServiceClient(c),
	}

	routeGroup := r.Group("/product", authMiddleware.AuthRequired)
	routeGroup.GET("/", svc.ListProduct)
	routeGroup.GET("/:id", svc.GetProduct)
	routeGroup.POST("/", svc.CreateProduct)
	routeGroup.PUT("/:id", svc.UpdateProduct)
	routeGroup.DELETE("/:id", svc.DeleteProduct)

	return svc
}

func (client *ServiceClient) ListProduct(ctx *gin.Context) {
	response, err := client.ProductClient.ListProduct(ctx, &pb.ListProductRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var products []*pb.Product
	for {
		product, err := response.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		products = append(products, product)
	}

	ctx.JSON(http.StatusOK, &products)
}

func (client *ServiceClient) GetProduct(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := client.ProductClient.GetProduct(ctx, &pb.GetProductIdRequest{
		Id: id,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}

func (client *ServiceClient) CreateProduct(ctx *gin.Context) {
	var req request.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := client.ProductClient.CreateProduct(ctx, &pb.CreateProductRequest{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}

func (client *ServiceClient) UpdateProduct(ctx *gin.Context) {
	var req request.UpdateProductRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := client.ProductClient.UpdateProduct(ctx, &pb.UpdateProductRequest{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}

func (client *ServiceClient) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	response, err := client.ProductClient.DeleteProduct(ctx, &pb.GetProductIdRequest{
		Id: id,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
