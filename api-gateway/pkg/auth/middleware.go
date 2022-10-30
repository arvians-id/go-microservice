package auth

import (
	"github.com/arvians-id/go-microservice/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type MiddlewareAuthConfig struct {
	ServiceClient *ServiceClient
}

func NewMiddlewareAuthConfig(client *ServiceClient) *MiddlewareAuthConfig {
	return &MiddlewareAuthConfig{ServiceClient: client}
}

func (m *MiddlewareAuthConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")
	if authorization == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
		return
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
		return
	}

	res, err := m.ServiceClient.Client.Validate(ctx, &pb.ValidateRequest{
		Token: token[1],
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.Set("user_id", res.UserId)
	ctx.Next()
}
