package main

import (
	app "AuthService/app"
	env "AuthService/config/env"
	db "AuthService/db/repository"
)

func main() {
	env.Load()
	config := app.NewConfig()
	storage := db.NewStorage()
	app := app.NewApplication(config, storage)
	app.Run()
}
