package repository

import (
	"context"
	"database/sql"

	"github.com/micnncim/uber-microservice-go/service/ride/entity"
)

type RideRepository interface {
	Create(ctx context.Context, ride *entity.Ride) error
	Get(ctx context.Context, id string) (*entity.Ride, error)
	Delete(ctx context.Context, id string) error
}

type rideRepository struct {
	conn *sql.Conn
}

func NewRideRepository(conn *sql.Conn) RideRepository {
	return &rideRepository{conn: conn}
}

func (r *rideRepository) Create(ctx context.Context, ride *entity.Ride) error {
	return nil
}

func (r *rideRepository) Get(ctx context.Context, id string) (*entity.Ride, error) {
	return nil, nil
}

func (r *rideRepository) Update(ctx context.Context) error {
	return nil
}

func (r *rideRepository) Delete(ctx context.Context, id string) error {
	return nil
}

var _ RideRepository = (*rideRepository)(nil)
