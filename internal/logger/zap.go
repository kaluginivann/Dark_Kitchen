package logger

import (
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/kaluginivann/Dark_Kitchen/config"
)

func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error

	switch cfg.Environment {
	case "DEV":
		logger, err = zap.NewDevelopment()
	case "PROD":
		logger, err = zap.NewProduction()
	default:
		return nil, errors.New("unknown environment")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	logger.Info("Selected environment", zap.String("env", cfg.Environment))
	return logger, nil
}
