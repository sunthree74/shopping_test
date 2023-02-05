package model

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	Name string `json:"name"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}
