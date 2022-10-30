package routes

import (
	"github.com/arvians-id/go-microservice/api-gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"net/http"
)

func ListUser(ctx *gin.Context, c pb.UserServiceClient) {
	response, err := c.ListUser(ctx, new(empty.Empty))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
