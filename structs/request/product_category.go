package request

type ProductCategoryForm struct {
	Name string `json:"name" binding:"required"`
}
