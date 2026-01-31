package db

import "database/sql"

type Storage struct {
	UserRepository UserRepository
}

func NewStorage() *Storage {
	return &Storage{
		UserRepository: NewUserRepository(&sql.DB{}),
	}
}
