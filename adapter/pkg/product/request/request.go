package request

type CreateProductRequest struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	CreatedBy   int64  `form:"created_by"`
}

type UpdateProductRequest struct {
	Name        string `form:"name"`
	Description string `form:"description"`
}
