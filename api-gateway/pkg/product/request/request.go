package request

type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   int64  `json:"created_by"`
}

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
