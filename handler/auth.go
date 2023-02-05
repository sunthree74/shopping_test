package handler

import "github.com/sunthree74/shopping_test/interfaces"

type authHandler struct {
	userUsecase interfaces.UserUsecase
}

func HandleAuth(userUsecase interfaces.UserUsecase) *authHandler {
	return &authHandler{userUsecase: userUsecase}
}
