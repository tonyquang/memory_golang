package router

import (
	"context"

	"memory_golang/api/internal/api/rest/health"
	"memory_golang/api/internal/controller/system"
	"memory_golang/api/internal/controller/user"
)

// New creates and returns a new Router instance
func New(
	ctx context.Context,
	corsOrigin []string,
	isGQLIntrospectionOn bool,
	systemCtrl system.Controller,
	userCtrl user.Controller,
) Router {
	return Router{
		ctx:                  ctx,
		corsOrigins:          corsOrigin,
		isGQLIntrospectionOn: isGQLIntrospectionOn,
		healthRESTHandler:    health.New(systemCtrl, userCtrl),
	}
}
