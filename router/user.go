package router

import (
	controller "AuthService/controller"
	"AuthService/middleware"

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
	router.Post("/register", this.userController.RegisterUser)
	router.With(middleware.LoginRequestValidator).Post("/login", this.userController.LoginUser)
}
