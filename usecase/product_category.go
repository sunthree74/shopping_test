package usecase

import (
	"context"
	"github.com/sunthree74/shopping_test/interfaces"
	"github.com/sunthree74/shopping_test/model"
	"github.com/sunthree74/shopping_test/structs/response"
	"strconv"
)

var _ interfaces.ProductCategoryUsecase = (*productCategoryUsecase)(nil)

type productCategoryUsecase struct {
	productCategoryRepo interfaces.ProductCategoryRepository
}

func (c *productCategoryUsecase) GetList(ctx context.Context) ([]response.ProductCategory, error) {
	var cnv []response.ProductCategory
	cnv, err := c.productCategoryRepo.GetList(ctx)
	if err != nil {
		return []response.ProductCategory{}, err
	}

	return cnv, nil
}

func (c *productCategoryUsecase) Create(ctx context.Context, category *model.ProductCategory) (*model.ProductCategory, error) {
	category, err := c.productCategoryRepo.Create(ctx, category)
	if err != nil {
		return &model.ProductCategory{}, err
	}

	return category, nil
}

func (c *productCategoryUsecase) Update(ctx context.Context, ID uint, newCategoryData model.ProductCategory) error {
	panic("implement me")
}

func (c *productCategoryUsecase) Delete(ctx context.Context, ID uint) error {
	panic("implement me")
}

func (c *productCategoryUsecase) FindById(ctx context.Context, ID uint) (model.ProductCategory, error) {
	category, err := c.productCategoryRepo.FindById(ctx, strconv.Itoa(int(ID)))
	if err != nil {
		return model.ProductCategory{}, err
	}

	return category, nil
}

func InitializeProductCategory(productCategoryRepo interfaces.ProductCategoryRepository) *productCategoryUsecase {
	return &productCategoryUsecase{productCategoryRepo: productCategoryRepo}
}
