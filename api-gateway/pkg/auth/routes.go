package auth

import (
	"github.com/arvians-id/go-microservice/api-gateway/pkg/auth/routes"
	"github.com/arvians-id/go-microservice/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitializeServiceClient(c),
	}

	routeGroup := r.Group("/auth")
	routeGroup.POST("/login", svc.Login)
	routeGroup.POST("/register", svc.Register)

	return svc
}

func (s *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, s.Client)
}

func (s *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, s.Client)
}
