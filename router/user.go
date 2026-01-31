package router

import (
	controller "AuthService/controller"

	chi "github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controller.UserController
}

func NewUserRouter(userController *controller.UserController) Router {
	return &UserRouter{
		userController: userController,
	}
}

func (this *UserRouter) Register(router chi.Router) {
	router.Get("/register", this.userController.RegisterUser)
}
