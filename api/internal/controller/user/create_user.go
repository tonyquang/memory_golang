package user

import (
	"context"

	"memory_golang/api/internal/model"
)

func (i impl) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	return i.repo.User().Insert(ctx, user)
}
