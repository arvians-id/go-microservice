package product

import (
	"github.com/arvians-id/go-microservice/api-gateway/pkg/auth"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/config"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/product/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	authMiddleware := auth.NewMiddlewareAuthConfig(authSvc)
	svc := &ServiceClient{
		Client: InitializeServiceClient(c),
	}

	routeGroup := r.Group("/product", authMiddleware.AuthRequired)
	routeGroup.GET("/", svc.ListProduct)
	routeGroup.GET("/:id", svc.GetProduct)
	routeGroup.POST("/", svc.CreateProduct)
	routeGroup.PUT("/:id", svc.UpdateProduct)
	routeGroup.DELETE("/:id", svc.DeleteProduct)

	return svc
}

func (s *ServiceClient) ListProduct(ctx *gin.Context) {
	routes.ListProduct(ctx, s.Client)
}

func (s *ServiceClient) GetProduct(ctx *gin.Context) {
	routes.GetProduct(ctx, s.Client)
}

func (s *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, s.Client)
}

func (s *ServiceClient) UpdateProduct(ctx *gin.Context) {
	routes.UpdateProduct(ctx, s.Client)
}

func (s *ServiceClient) DeleteProduct(ctx *gin.Context) {
	routes.DeleteProduct(ctx, s.Client)
}
