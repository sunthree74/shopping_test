package interfaces

import (
	"context"
	"github.com/sunthree74/shopping_test/model"
	"github.com/sunthree74/shopping_test/structs/response"
)

type ProductCategoryRepository interface {
	Create(ctx context.Context, category *model.ProductCategory) (*model.ProductCategory, error)
	Update(ctx context.Context, id string, category *model.ProductCategory) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (model.ProductCategory, error)
	GetList(ctx context.Context) ([]response.ProductCategory, error)
}
