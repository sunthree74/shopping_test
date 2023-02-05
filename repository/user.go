package repository

import (
	"context"
	"fmt"
	"github.com/sunthree74/shopping_test/interfaces"
	"github.com/sunthree74/shopping_test/model"
	"gorm.io/gorm"
	"net/http"
)

var _ interfaces.UserRepository = (*userRepository)(nil)

type userRepository struct {
	db         *gorm.DB
	httpClient *http.Client
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var data model.User
	if err := u.db.WithContext(ctx).Where("email = ?", email).Find(&data).Error; err != nil {
		return model.User{}, fmt.Errorf("user get by email query execution error: %w", err)
	}

	return data, nil
}

func (u *userRepository) Create(ctx context.Context, user *model.User) error {
	if err := u.db.WithContext(ctx).Save(user).Error; err != nil {
		return fmt.Errorf("user create query execution error: %w", err)
	}

	return nil
}

func (u *userRepository) Update(ctx context.Context, uid string, user *model.User) error {
	if err := u.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", uid).Updates(user).Error; err != nil {
		return fmt.Errorf("user update query execution error: %w", err)
	}

	return nil
}

func (u *userRepository) Delete(ctx context.Context, uid string) error {
	if err := u.db.WithContext(ctx).Delete(&model.User{}, uid).Error; err != nil {
		return fmt.Errorf("user delete query execution error: %w", err)
	}

	return nil
}

func (u userRepository) FindById(ctx context.Context, uid string) (model.User, error) {
	var data model.User
	if err := u.db.WithContext(ctx).Where("id = ?", uid).First(&data).Error; err != nil {
		return model.User{}, fmt.Errorf("user get by id query execution error: %w", err)
	}

	return data, nil
}

func InitializeUser(db *gorm.DB, httpClient *http.Client) *userRepository {
	return &userRepository{db: db, httpClient: httpClient}
}
