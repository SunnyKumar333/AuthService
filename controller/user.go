package controller

import (
	dto "AuthService/dto"
	service "AuthService/service"
	utils "AuthService/utils"
	"fmt"
	"net/http"
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

	loginUserPayload := r.Context().Value("payload").(dto.LoginUserDTO)

	fmt.Println("payload:", loginUserPayload)

	//attempt to login
	jwtToken, loginErr := this.userService.LoginUser(&loginUserPayload)
	if loginErr != nil {
		errorMessage := "Invalid Credential"
		utils.WriteJSONErrorResponse(w, http.StatusUnauthorized, loginErr, errorMessage)
		return
	}

	responseMessage := "Login Successfull"

	utils.WriteJSONSuccessResponse(w, http.StatusOK, jwtToken, responseMessage)
}
