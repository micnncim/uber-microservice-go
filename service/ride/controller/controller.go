package controller

import (
	"context"

	"go.uber.org/zap"

	"github.com/micnncim/uber-microservice-go/service/ride/entity"
	"github.com/micnncim/uber-microservice-go/service/ride/repository"
)

type RideController interface {
	Create(ctx context.Context, ride *entity.Ride) error
	Get(ctx context.Context, id string) (*entity.Ride, error)
	Delete(ctx context.Context, id string) error
}

type rideController struct {
	rideRepository repository.RideRepository
	logger         *zap.Logger
}

type Option func(*rideController)

func WithLogger(l *zap.Logger) Option {
	return func(c *rideController) {
		c.logger = l
	}
}

func NewRideController(r repository.RideRepository, opts ...Option) RideController {
	c := &rideController{
		rideRepository: r,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *rideController) Create(ctx context.Context, ride *entity.Ride) error {
	return c.rideRepository.Create(ctx, ride)
}

func (c *rideController) Get(ctx context.Context, id string) (*entity.Ride, error) {
	return c.rideRepository.Get(ctx, id)
}

func (c *rideController) Delete(ctx context.Context, id string) error {
	return c.rideRepository.Delete(ctx, id)
}

var _ RideController = (*rideController)(nil)
