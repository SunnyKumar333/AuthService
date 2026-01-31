package router

import (
	controller "AuthService/controller"

	chi "github.com/go-chi/chi/v5"
)

type PingRouter struct {
	pingController controller.PingController
}

func NewPingRouter(pingController controller.PingController) Router {
	return &PingRouter{
		pingController: pingController,
	}
}

func (this *PingRouter) Register(router chi.Router) {
	router.Get("/ping", this.pingController.PingHandler)
}
