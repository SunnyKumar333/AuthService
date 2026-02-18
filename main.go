package main

import (
	app "AuthService/app"
	dbConfig "AuthService/config/db"
	env "AuthService/config/env"
	db "AuthService/db/repository"
)

func main() {
	env.Load()
	config := app.NewConfig()
	conn, err := dbConfig.SetupDB()
	if err != nil {
		panic(err)
	}
	storage := db.NewStorage(conn)
	app := app.NewApplication(config, storage)
	app.Run()
}
