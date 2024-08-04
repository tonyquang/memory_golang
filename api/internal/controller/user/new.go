package user

import (
	"context"

	"memory_golang/api/internal/model"
	"memory_golang/api/internal/repository"
)

type Controller interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry) Controller {
	return impl{repo: repo}
}

type impl struct {
	repo repository.Registry
}
