package controller

import (
	dto "AuthService/dto"
	service "AuthService/service"
	utils "AuthService/utils"
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

	loginUserPayload := &dto.LoginUserDTO{}
	// serializing body into json
	if jsonErr := utils.ReadJSONBody(r, loginUserPayload); jsonErr != nil {
		errorMessage := "Invalid Request Body"
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, jsonErr, errorMessage)
		return
	}

	//validating json body using validator package
	if validationErr := utils.Validator.Struct(loginUserPayload); validationErr != nil {
		errorMessage := "Validation Error"
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, validationErr, errorMessage)
		return
	}
	//attempt to login
	jwtToken, loginErr := this.userService.LoginUser(loginUserPayload)
	if loginErr != nil {
		errorMessage := "Invalid Credential"
		utils.WriteJSONErrorResponse(w, http.StatusUnauthorized, loginErr, errorMessage)
		return
	}

	responseMessage := "Login Successfull"

	utils.WriteJSONSuccessResponse(w, http.StatusOK, jwtToken, responseMessage)
}
