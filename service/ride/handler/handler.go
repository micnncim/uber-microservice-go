package handler

import (
	"github.com/micnncim/uber-microservice-go/service/ride/controller"
)

type RideHandler interface {
}

type rideHandler struct {
	rideController controller.RideController
}

func NewRideHandler(rideController controller.RideController) RideHandler {
	return &rideHandler{
		rideController: rideController,
	}
}
