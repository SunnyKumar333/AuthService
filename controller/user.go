package controller

import (
	service "AuthService/service"
	http "net/http"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userServive service.UserService) *UserController {
	return &UserController{}
}

func (this *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Responce from User Registeration!!"))
}
