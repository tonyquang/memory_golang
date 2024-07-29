package router

import (
	"context"
	"net/http"

	"memory_golang/api/internal/api/rest/health"
	"memory_golang/api/pkg/httpserv"
)

// Router defines the routes & handlers of the app
type Router struct {
	ctx                  context.Context
	corsOrigins          []string
	isGQLIntrospectionOn bool
	healthRESTHandler    health.Handler
}

// Handler returns the Handler for use by the server
func (rtr Router) Handler() http.Handler {
	return httpserv.Handler(
		rtr.healthRESTHandler.CheckReadiness(),
	)
}
