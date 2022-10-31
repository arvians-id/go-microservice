package model

type Product struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   int64  `json:"created_by"`
	User        *User  `json:"user"`
}
