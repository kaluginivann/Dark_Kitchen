package users

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kaluginivann/Dark_Kitchen/config"
	"github.com/kaluginivann/Dark_Kitchen/pkg/logger"
	"github.com/kaluginivann/Dark_Kitchen/pkg/req"
	"github.com/kaluginivann/Dark_Kitchen/pkg/res"
)

type UserHandler struct {
	UserService IService
	Config      *config.Config
	Logger      logger.Interface
}

func NewUserHandler(router *chi.Mux, userService IService, config *config.Config, logger logger.Interface) {
	handler := &UserHandler{UserService: userService, Config: config, Logger: logger}
	router.Route(fmt.Sprintf("%s/users", config.Server.BaseApi), func(r chi.Router) {
		r.Post("/register", handler.Register())
	})
}

func (u *UserHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.Decode[RegisterRequest](r.Body)
		if err != nil {
			res.JSON(w, "Invalid Body", http.StatusBadRequest, u.Logger)
			return
		}
		registerResp, err := u.UserService.Register(r.Context(), &body)
		if err != nil {
			if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
				if pgErr.Code == UserAlreadyExistCode {
					u.Logger.Error("user already exists")
					res.JSON(w, UserAlreadyExist, http.StatusBadRequest, u.Logger)
					return
				}
			}
			u.Logger.Error("service error", "error", err)
			res.JSON(w, "Internal service error", http.StatusInternalServerError, u.Logger)
			return
		}
		res.JSON(w, registerResp, http.StatusCreated, u.Logger)
	}
}
