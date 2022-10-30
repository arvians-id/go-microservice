package routes

import (
	"github.com/arvians-id/go-microservice/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, client pb.AuthServiceClient) {
	var request RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := client.Register(ctx, &pb.RegisterRequest{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
