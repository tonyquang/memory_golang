package router

import (
	"context"
	"net/http"

	"memory_golang/api/internal/api/rest/health"
	"memory_golang/api/pkg/httpserv"

	"github.com/go-chi/chi/v5"
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
		rtr.routes,
	)
}

func (rtr Router) routes(r chi.Router) {
	r.Group(rtr.public)
}

func (rtr Router) public(r chi.Router) {
	prefix := "/api/public"

	// v1
	r.Group(func(r chi.Router) {
		prefix = prefix + "/v1"
		r.Group(func(r chi.Router) {
			r.Get(prefix+"/ping", func(writer http.ResponseWriter, request *http.Request) {
				httpserv.RespondJSON(request.Context(), writer, map[string]string{
					"message": "pong",
				})
			})
			r.Get(prefix+"/channel", rtr.healthRESTHandler.CallLeakChannel())
			r.Get(prefix+"/goroutine", rtr.healthRESTHandler.CallLeakGoRoutine())
			r.Get(prefix+"/map", rtr.healthRESTHandler.CallLeakMap())
			r.Post(prefix+"/users", rtr.healthRESTHandler.CreateUser())
		})
	})
}
