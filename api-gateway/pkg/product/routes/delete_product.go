package routes

import (
	"github.com/arvians-id/go-microservice/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	response, err := c.DeleteProduct(ctx, &pb.GetProductIdRequest{
		Id: id,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
