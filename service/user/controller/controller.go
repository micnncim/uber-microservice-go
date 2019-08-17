package controller

import (
	"context"

	"github.com/micnncim/uber-microservice-go/service/user/entity"
	"github.com/micnncim/uber-microservice-go/service/user/repository"
)

type UserController interface {
	Create(ctx context.Context, user *entity.User) error
	Get(ctx context.Context, id string) (*entity.User, error)
	List(ctx context.Context) ([]*entity.User, error)
	Delete(ctx context.Context, id string) error
}

type userController struct {
	r repository.UserRepository
}

func NewUserController(r repository.UserRepository) UserController {
	return &userController{r: r}
}

func (c *userController) Create(ctx context.Context, user *entity.User) error {
	return c.r.Create(ctx, user)
}

func (c *userController) Get(ctx context.Context, id string) (*entity.User, error) {
	return c.r.Get(ctx, id)
}

func (c *userController) List(ctx context.Context) ([]*entity.User, error) {
	return c.r.List(ctx)
}

func (c *userController) Delete(ctx context.Context, id string) error {
	return c.r.Delete(ctx, id)
}

var _ UserController = (*userController)(nil)
