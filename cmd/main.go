package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/kaluginivann/Dark_Kitchen/config"
	"github.com/kaluginivann/Dark_Kitchen/internal/logger"
	"github.com/kaluginivann/Dark_Kitchen/internal/repository"
)

func main() {
	build().Run()
}

func build() *fx.App {
	return fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: log.WithOptions(zap.IncreaseLevel(zapcore.ErrorLevel)),
			}
		}),
		config.Module,
		logger.Module,
		repository.Module,
		//server.Module,
	)
}
