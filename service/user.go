package service

import (
	db "AuthService/db/repository"
)

type UserService interface {
	CreateUser() error
}

type userServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(userRepository db.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

func (this *userServiceImpl) CreateUser() error {
	return nil
}
