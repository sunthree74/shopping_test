package usecase

import (
	"context"
	"errors"
	"github.com/sunthree74/shopping_test/interfaces"
	"github.com/sunthree74/shopping_test/model"
	"strconv"
)

var _ interfaces.UserUsecase = (*userUsecase)(nil)

type userUsecase struct {
	userRepo interfaces.UserRepository
}

func (u *userUsecase) Login(ctx context.Context, email string, password string) (model.User, error) {
	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	if user.Password != password {
		return model.User{}, errors.New("Email/Password Salah")
	}

	return user, nil
}

func (u *userUsecase) Create(ctx context.Context, user *model.User) error {
	result, resErr := u.userRepo.FindByEmail(ctx, user.Email)
	if resErr != nil {
		return resErr
	}
	if result.ID != 0 {
		return errors.New("Email suah terdaftar")
	}
	err := u.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Update(ctx context.Context, userID uint, newUserData model.User) error {
	panic("implement me")
}

func (u *userUsecase) Delete(ctx context.Context, userID uint) error {
	panic("implement me")
}

func (u *userUsecase) FindById(ctx context.Context, userID uint) (model.User, error) {
	user, err := u.userRepo.FindById(ctx, strconv.Itoa(int(userID)))
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userUsecase) FindByEmail(ctx context.Context, email string) (model.User, error) {
	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func InitializeUser(userRepo interfaces.UserRepository) *userUsecase {
	return &userUsecase{userRepo: userRepo}
}
