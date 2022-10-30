package routes

import (
	"github.com/arvians-id/go-microservice/api-gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateUserRequest struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UpdateUser(ctx *gin.Context, c pb.UserServiceClient) {
	var req UpdateUserRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:       req.Id,
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
