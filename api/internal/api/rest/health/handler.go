package health

import (
	"encoding/json"
	"net/http"
	"strconv"

	"memory_golang/api/internal/controller/system"
	"memory_golang/api/internal/controller/user"
	"memory_golang/api/internal/model"
	"memory_golang/api/pkg/httpserv"
)

// Handler is the web handler for this pkg
type Handler struct {
	systemCtrl system.Controller
	userCtrl   user.Controller
}

// New instantiates a new Handler and returns it
func New(systemCtrl system.Controller, userCtrl user.Controller) Handler {
	return Handler{systemCtrl: systemCtrl, userCtrl: userCtrl}
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

type userReqResp struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (h Handler) CreateUser() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		decoder := json.NewDecoder(r.Body)
		var userReq userReqResp
		if err := decoder.Decode(&userReq); err != nil {
			return httpserv.Error{
				Status: http.StatusBadRequest,
				Code:   "err_create_user",
				Desc:   "invalid create user request",
			}
		}

		ctx := r.Context()
		userModel, err := h.userCtrl.CreateUser(ctx, model.User{Email: userReq.Email})
		if err != nil {
			return err
		}

		httpserv.RespondJSON(ctx, w, userReqResp{
			ID:    strconv.Itoa(userModel.ID),
			Email: userModel.Email,
		})

		return nil
	})
}
