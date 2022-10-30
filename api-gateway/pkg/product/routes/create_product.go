package routes

import (
	"github.com/arvians-id/go-microservice/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   int64  `json:"created_by"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	var req CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.CreateProduct(ctx, &pb.CreateProductRequest{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
