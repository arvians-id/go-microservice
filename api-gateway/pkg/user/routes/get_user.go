package routes

import (
	"github.com/arvians-id/go-microservice/api-gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(ctx *gin.Context, c pb.UserServiceClient) {
	id, err := strconv.Atoi(ctx.Param("id"))
	response, err := c.GetUser(ctx, &pb.GetUserIdRequest{
		Id: int64(id),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
