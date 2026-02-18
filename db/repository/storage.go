package db

import "database/sql"

type Storage struct {
	UserRepository UserRepository
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		UserRepository: NewUserRepository(db),
	}
}
