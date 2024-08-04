package user

import (
	"context"

	"memory_golang/api/internal/model"
	"memory_golang/api/pkg/db/pg"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	Insert(ctx context.Context, user model.User) (model.User, error)
}

// New returns an implementation instance satisfying Repository
func New(dbConn pg.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn pg.ContextExecutor
}
