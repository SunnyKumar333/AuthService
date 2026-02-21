package service

import (
	db "AuthService/db/repository"
	dto "AuthService/dto"
	modals "AuthService/models"
	"AuthService/utils"
	hash "AuthService/utils"
	"fmt"

	env "AuthService/config/env"

	jwt "github.com/golang-jwt/jwt/v5"

	"errors"
)

type UserService interface {
	CreateUser(*dto.UserDTO) (*modals.User, error)
	LoginUser(*dto.LoginUserDTO) (string, error)
	GetUserById(string) (*dto.UserResponseDTO, error)
}

type userServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(userRepository db.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

func (this *userServiceImpl) CreateUser(userDTO *dto.UserDTO) (*modals.User, error) {
	hashedPassword, err := hash.HashPassword(userDTO.Password)
	if err != nil {
		return nil, err
	}
	userDTO.Password = hashedPassword
	user, err := this.userRepository.Create(userDTO)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (this *userServiceImpl) LoginUser(loginUserPayload *dto.LoginUserDTO) (string, error) {
	password := loginUserPayload.Password
	email := loginUserPayload.Email

	// 1. get user from database with given email
	user, err := this.userRepository.GetByEmail(email)

	if err != nil {
		return "", err
	}
	fmt.Println(user)

	// 2. verify password and hased_password
	isPasswordValid := utils.ValidatePassword(password, user.Password)

	if !isPasswordValid {
		return "", errors.New("Password dosen't Match")
	}

	// 3. create jwt token
	payload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	fmt.Println(env.GetString("JWT_SECRET", "my_secret"))
	token, err := tokenObj.SignedString([]byte(env.GetString("JWT_SECRET", "my_secret")))
	if err != nil {
		fmt.Println("Error wile generating JWt Token:", err)
		return "", nil
	}
	fmt.Println(token)

	return token, nil
}

func (this *userServiceImpl) GetUserById(userId string) (*dto.UserResponseDTO, error) {
	fmt.Println("Finding user with user_id:", userId)
	return this.userRepository.GetById(userId)
}
