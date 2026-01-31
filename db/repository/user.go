package db

import (
	"database/sql"
)

type UserRepository interface {
	Create() error
}

type UserSqlRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserSqlRepository{
		db: db,
	}
}

func (this *UserSqlRepository) Create() error {
	return nil
}
