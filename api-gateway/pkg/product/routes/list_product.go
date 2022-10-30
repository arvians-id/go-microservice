package routes

import (
	"github.com/arvians-id/go-microservice/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"net/http"
)

func ListProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	response, err := c.ListProduct(ctx, new(empty.Empty))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
