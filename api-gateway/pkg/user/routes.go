package user

import (
	"github.com/arvians-id/go-microservice/api-gateway/pkg/auth"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/config"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	authMiddleware := auth.NewMiddlewareAuthConfig(authSvc)
	svc := &ServiceClient{
		Client: InitializeServiceClient(c),
	}

	routeGroup := r.Group("/user", authMiddleware.AuthRequired)
	routeGroup.GET("/", svc.ListUser)
	routeGroup.GET("/:id", svc.GetUser)
	routeGroup.POST("/", svc.CreateUser)
	routeGroup.PUT("/:id", svc.UpdateUser)
	routeGroup.DELETE("/:id", svc.DeleteUser)

	return svc
}

func (s *ServiceClient) ListUser(ctx *gin.Context) {
	routes.ListUser(ctx, s.Client)
}

func (s *ServiceClient) GetUser(ctx *gin.Context) {
	routes.GetUser(ctx, s.Client)
}

func (s *ServiceClient) CreateUser(ctx *gin.Context) {
	routes.CreateUser(ctx, s.Client)
}

func (s *ServiceClient) UpdateUser(ctx *gin.Context) {
	routes.UpdateUser(ctx, s.Client)
}

func (s *ServiceClient) DeleteUser(ctx *gin.Context) {
	routes.DeleteUser(ctx, s.Client)
}
