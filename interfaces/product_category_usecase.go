package interfaces

import (
	"context"
	"github.com/sunthree74/shopping_test/model"
	"github.com/sunthree74/shopping_test/structs/response"
)

type ProductCategoryUsecase interface {
	Create(ctx context.Context, category *model.ProductCategory) (*model.ProductCategory, error)
	Update(ctx context.Context, ID uint, newCategoryData model.ProductCategory) error
	Delete(ctx context.Context, ID uint) error
	FindById(ctx context.Context, ID uint) (model.ProductCategory, error)
	GetList(ctx context.Context) ([]response.ProductCategory, error)
}
