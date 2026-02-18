package controller

import (
	dto "AuthService/dto"
	service "AuthService/service"
	"fmt"
	http "net/http"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userServive service.UserService) *UserController {
	return &UserController{
		userService: userServive,
	}
}

func (this *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	userDTO := &dto.UserDTO{
		Username: "sunny",
		Email:    "sunny@gmail.com",
		Password: "Password@123",
	}
	this.userService.CreateUser(userDTO)
	w.Write([]byte("Responce from User Registeration!!"))
}

func (this *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	// this.userService.GetUserById()
}

func (this *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	result, err := this.userService.LoginUser()
	if err != nil {
		fmt.Println(err)
	}
	w.Write([]byte(result))
}
