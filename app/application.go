package app

import (
	env "AuthService/config/env"
	controller "AuthService/controller"
	db "AuthService/db/repository"
	router "AuthService/router"
	service "AuthService/service"
	"fmt"
	http "net/http"
	"time"
)

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		Port: env.GetString("PORT", ":8080"),
	}
}

type Application struct {
	Config  *Config
	Storage *db.Storage
}

func NewApplication(config *Config, storage *db.Storage) *Application {
	return &Application{
		Config:  config,
		Storage: storage,
	}
}

func (this *Application) Run() error {
	userRepository := this.Storage.UserRepository
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRouter := router.NewUserRouter(userController)
	// Application run logic here
	server := &http.Server{
		Addr:         this.Config.Port,
		Handler:      router.SetupRouter(userRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on port", this.Config.Port)

	return server.ListenAndServe()
}
