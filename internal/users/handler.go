package users

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kaluginivann/Dark_Kitchen/config"
	"github.com/kaluginivann/Dark_Kitchen/pkg/logger"
)

type UserHandler struct {
	UserService IService
	Config      *config.Config
	Logger      logger.Interface
}

func NewUserHandler(api huma.API, userService IService, config *config.Config, logger logger.Interface) {
	handler := &UserHandler{UserService: userService, Config: config, Logger: logger}
	huma.Register(
		api,
		huma.Operation{
			OperationID: "register-user",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/users/register", config.Server.BaseApi),
			Summary:     "Register Users",
		},
		handler.Register,
	)
}

func (u *UserHandler) Register(ctx context.Context, input *RegisterInput) (*RegisterOutput, error) {
	registerResp, err := u.UserService.Register(ctx, &input.Body)
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == UserAlreadyExistCode {
				u.Logger.Error("user already exists")
				return nil, huma.Error400BadRequest(UserAlreadyExist)
			}
		}
		u.Logger.Error("service error", "error", err)
		return nil, huma.Error500InternalServerError("Internal service error")
	}
	return &RegisterOutput{
		Body: RegisterResponse{
			Id:       registerResp.Id,
			Username: registerResp.Username,
			Email:    registerResp.Email,
		},
	}, nil
}
