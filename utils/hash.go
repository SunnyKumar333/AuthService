package utils

import (
	"fmt"

	bcrypt "golang.org/x/crypto/bcrypt"
)

func HashPassword(plainPassword string) (string, error) {
	fmt.Println("Generating Hash Password.......")
	hassedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hassedPassword), nil
}

func ValidatePassword(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
