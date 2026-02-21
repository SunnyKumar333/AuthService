package router

import (
	// controller "AuthService/controller"

	"AuthService/controller"
	"AuthService/middleware"

	chi "github.com/go-chi/chi/v5"
)

type Router interface {
	Register(chi.Router)
}

// var PingController = controller.NewPingController()

func SetupRouter(userRouter Router) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestLogger)
	router.Use(middleware.RateLimiterMiddelware)
	router.Get("/ping", controller.NewPingController().PingHandler)
	userRouter.Register(router)

	return router
}
