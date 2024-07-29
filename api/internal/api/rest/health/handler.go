package health

import (
	"memory_golang/api/internal/controller/system"
	"memory_golang/api/pkg/httpserv"
	"net/http"
)

// Handler is the web handler for this pkg
type Handler struct {
	systemCtrl system.Controller
}

// New instantiates a new Handler and returns it
func New(systemCtrl system.Controller) Handler {
	return Handler{systemCtrl: systemCtrl}
}

func (h Handler) CallLeakChannel() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		go h.systemCtrl.MonitorChannel()

		return nil
	})
}

func (h Handler) CallLeakGoRoutine() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		go h.systemCtrl.MonitorGoroutine()

		return nil
	})
}

func (h Handler) CallLeakMap() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		go h.systemCtrl.MonitorMap()

		return nil
	})
}
