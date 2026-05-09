package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/kaluginivann/Dark_Kitchen/config"
)

type Server struct {
	httpServer *http.Server
	logger     *zap.Logger
}

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(30 * time.Second))
	return r
}

func New(cfg *config.Config, logger *zap.Logger, router *chi.Mux) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
			Handler: router,
		},
		logger: logger,
	}
}

func manageServerLifecycle(lc fx.Lifecycle, s *Server) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			s.logger.Info("Starting HTTP server", zap.String("addr", s.httpServer.Addr))
			go func() {
				if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					s.logger.Error("HTTP server error", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.logger.Info("Shutting down HTTP server")
			return s.httpServer.Shutdown(ctx)
		},
	})
}
