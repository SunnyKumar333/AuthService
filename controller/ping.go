package controller

import (
	"net/http"
)

type PingController struct {
}

func NewPingController() *PingController {
	return &PingController{}
}

func (this *PingController) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}
