package service

import (
	"go.uber.org/fx"

	"github.com/kaluginivann/Dark_Kitchen/internal/service/users"
)

var Module = fx.Module("service",
	fx.Provide(users.New),
)
