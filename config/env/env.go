package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Some Error Occur while loading dotenv")
	}
}

func GetString(key string, fallback string) string {
	//this function take key and fallbact value as string and return there value from env and if key is not available then return fallback value
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}

func GetInt(key string, fallback int) int {
	//this function take key and fallbact value as int and return there value from env and if key is not available then return fallback value
	valueStr, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	var value int
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return fallback
	}
	return value
}

func GetBool(key string, fallback bool) bool {
	//this function take key and fallbact value as bool and return there value from env and if key is not available then return fallback value
	valueStr, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	var value bool
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return fallback
	}
	return value
}
