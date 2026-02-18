package db

import (
	dto "AuthService/dto"
	models "AuthService/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Create(*dto.UserDTO) (*models.User, error)
	GetById(int) (*models.User, error)
	GetByEmail(string) (*models.User, error)
	// GetAll() ([]*models.User, error)
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

func (this *UserSqlRepository) Create(userDTO *dto.UserDTO) (*models.User, error) {
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

	return &models.User{}, nil

}

func (this *UserSqlRepository) GetById(userId int) (*models.User, error) {
	query := "SELECT * FROM users WHERE id=?"
	row := this.db.QueryRow(query, 2)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No User found with given id", err)
			return nil, err
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
			return nil, err
		} else {
			fmt.Println("Error:", err)
			return nil, err
		}

	}
	fmt.Println("Fount user with email:", email)
	return user, nil

}
