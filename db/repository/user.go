package db

import (
	dto "AuthService/dto"
	models "AuthService/models"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository interface {
	Create(*dto.UserDTO) (*dto.UserResponseDTO, error)
	GetById(string) (*dto.UserResponseDTO, error)
	GetByEmail(string) (*models.User, error)
	// DeleteById(int) (*models.User, error)
}

type UserSqlRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserSqlRepository{
		db: db,
	}
}

func (this *UserSqlRepository) Create(userDTO *dto.UserDTO) (*dto.UserResponseDTO, error) {
	query := "INSERT INTO users (username,email,password_hashed) VALUES (?,?,?)"
	result, err := this.db.Exec(query, userDTO.Username, userDTO.Email, userDTO.Password)

	if err != nil {
		fmt.Println("Error while creating user:", err)
		return nil, err
	}

	rowAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("Row affected error:", rowErr)
		return nil, rowErr
	}

	if rowAffected == 0 {
		fmt.Println("User not created!")
		return nil, nil
	}

	fmt.Println("User Created Successfully!")

	user := &dto.UserResponseDTO{}

	query = "SELECT id,username,email FROM users WHERE email=?"

	row := this.db.QueryRow(query, userDTO.Email)
	err = row.Scan(&user.Id, &user.Username, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No User fount with email:", user.Email)
			return nil, errors.New("No Record Found with this email")
		} else {
			fmt.Println("Error:", err)
			return nil, err
		}

	}
	// fmt.Println("Fount user with email:", email)
	return user, nil

}

func (this *UserSqlRepository) GetById(userId string) (*dto.UserResponseDTO, error) {
	query := "SELECT id,username,email FROM users WHERE id=?"
	row := this.db.QueryRow(query, userId)

	user := &dto.UserResponseDTO{}

	err := row.Scan(&user.Id, &user.Username, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No User found with given id", err)
			return nil, errors.New("No Record Found with this id")
		} else {
			fmt.Println("Error while scanning rows:", err)
			return nil, err
		}

	}

	fmt.Println("User fetched successfully:", user)
	return user, nil

}

func (this *UserSqlRepository) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id,username,email,password_hashed FROM users WHERE email= ? "
	row := this.db.QueryRow(query, email)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No User fount with email:", email)
			return nil, errors.New("No Record Found with this email")
		} else {
			fmt.Println("Error:", err)
			return nil, err
		}

	}
	fmt.Println("Fount user with email:", email)
	return user, nil

}
