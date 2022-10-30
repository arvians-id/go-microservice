package routes

import (
	"github.com/arvians-id/go-microservice/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, client pb.AuthServiceClient) {
	var request LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := client.Login(ctx, &pb.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
