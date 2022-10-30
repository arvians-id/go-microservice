package routes

import (
	"github.com/arvians-id/go-microservice/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateProductRequest struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func UpdateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	var request UpdateProductRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.UpdateProduct(ctx, &pb.UpdateProductRequest{
		Id:          request.Id,
		Name:        request.Name,
		Description: request.Description,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
