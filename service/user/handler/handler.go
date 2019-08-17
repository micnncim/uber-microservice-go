package handler

import (
	"github.com/micnncim/uber-microservice-go/service/user/controller"
)

type UserHandler interface {
}

type userHandler struct {
	userController controller.UserController
}

func NewUserHandler(userController controller.UserController) UserHandler {
	return &userHandler{
		userController: userController,
	}
}
