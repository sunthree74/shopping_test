package repository

import (
	"context"
	"fmt"
	"github.com/sunthree74/shopping_test/interfaces"
	"github.com/sunthree74/shopping_test/model"
	"github.com/sunthree74/shopping_test/structs/response"
	"gorm.io/gorm"
	"net/http"
)

var _ interfaces.ProductCategoryRepository = (*productCategoryRepository)(nil)

type productCategoryRepository struct {
	db         *gorm.DB
	httpClient *http.Client
}

func (c *productCategoryRepository) GetList(ctx context.Context) ([]response.ProductCategory, error) {
	var categories []response.ProductCategory
	err := c.db.WithContext(ctx).Model(model.ProductCategory{}).Scan(&categories).
		Error
	if err != nil {
		return []response.ProductCategory{}, fmt.Errorf("get product category query error: %w", err)
	}

	if len(categories) < 1 {
		return []response.ProductCategory{}, gorm.ErrRecordNotFound
	}

	return categories, nil
}

func (c *productCategoryRepository) Create(ctx context.Context, category *model.ProductCategory) (*model.ProductCategory, error) {
	if err := c.db.WithContext(ctx).Create(&category).Error; err != nil {
		return &model.ProductCategory{}, fmt.Errorf("product category create query execution error: %w", err)
	}

	return category, nil
}

func (c *productCategoryRepository) Update(ctx context.Context, id string, category *model.ProductCategory) error {
	if err := c.db.WithContext(ctx).Model(&model.ProductCategory{}).Where("id = ?", id).Updates(category).Error; err != nil {
		return fmt.Errorf("product category update query execution error: %w", err)
	}

	return nil
}

func (c *productCategoryRepository) Delete(ctx context.Context, id string) error {
	if err := c.db.WithContext(ctx).Delete(&model.ProductCategory{}, id).Error; err != nil {
		return fmt.Errorf("product category delete query execution error: %w", err)
	}

	return nil
}

func (c productCategoryRepository) FindById(ctx context.Context, id string) (model.ProductCategory, error) {
	var data model.ProductCategory
	if err := c.db.WithContext(ctx).Where("id = ?", id).Find(&data).Error; err != nil {
		return model.ProductCategory{}, fmt.Errorf("product category get by id query execution error: %w", err)
	}

	return data, nil
}

func InitializeProductCategory(db *gorm.DB, httpClient *http.Client) *productCategoryRepository {
	return &productCategoryRepository{db: db, httpClient: httpClient}
}
