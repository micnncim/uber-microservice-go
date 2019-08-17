package repository

import (
	"context"

	"github.com/gocql/gocql"

	"github.com/micnncim/uber-microservice-go/service/user/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	Get(ctx context.Context, id string) (*entity.User, error)
	List(ctx context.Context) ([]*entity.User, error)
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	session *gocql.Session
}

func NewUserRepository(session *gocql.Session) UserRepository {
	return &userRepository{session: session}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	return nil
}

func (r *userRepository) Get(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}

func (r *userRepository) List(ctx context.Context) ([]*entity.User, error) {
	return nil, nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	return nil
}

var _ UserRepository = (*userRepository)(nil)
