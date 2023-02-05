package interfaces

import (
	"context"
	"github.com/sunthree74/shopping_test/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, uid string, user *model.User) error
	Delete(ctx context.Context, uid string) error
	FindById(ctx context.Context, uid string) (model.User, error)
	FindByEmail(ctx context.Context, email string) (model.User, error)
}
