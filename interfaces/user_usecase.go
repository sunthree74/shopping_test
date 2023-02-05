package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sunthree74/shopping_test/model"
)

type UserUsecase interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, userID uint, newUserData model.User) error
	Delete(ctx context.Context, userID uint) error
	FindById(ctx context.Context, userID uint) (model.User, error)
	FindByEmail(ctx context.Context, email string) (model.User, error)
	Login(ctx context.Context, email string, password string) (model.User, error)
	FindByJWT(ctx *gin.Context) (model.User, error)
}
