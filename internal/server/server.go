package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kaluginivann/Dark_Kitchen/config"
	"github.com/kaluginivann/Dark_Kitchen/internal/users"
	"github.com/kaluginivann/Dark_Kitchen/pkg/db"
	"github.com/kaluginivann/Dark_Kitchen/pkg/logger"
)

type Interface interface {
	Run()
}

type Server struct {
	Config *config.Config
	Logger logger.Interface
}

func New(config *config.Config, log logger.Interface) *Server {
	return &Server{
		Config: config,
		Logger: log,
	}
}

func (s *Server) buildDependencies() (*chi.Mux, *db.Database, error) {
	Database, err := db.New(s.Config)
	if err != nil {
		s.Logger.Error("Error from start db", "error", err)
		return nil, nil, err
	}

	// Repository
	UserRepository := users.NewRepository(Database, s.Logger)

	// Service
	UserService := users.NewService(UserRepository, s.Logger)

	// Main router
	mainRouter := chi.NewRouter()
	mainRouter.Use(middleware.RequestID)
	mainRouter.Use(middleware.RealIP)
	mainRouter.Use(middleware.Recoverer)
	mainRouter.Use(middleware.Logger)
	mainRouter.Use(middleware.Timeout(30 * time.Second))

	// Handlers
	users.NewUserHandler(mainRouter, UserService, s.Config, s.Logger)

	return mainRouter, Database, nil
}

func (s *Server) Run() {
	mainRouter, dataBase, err := s.buildDependencies()
	if err != nil {
		s.Logger.Error("Server is not set", "error", err)
		return
	}
	defer dataBase.Db.Close()

	s.Logger.Info("Server start listen", "port", s.Config.Server.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.Config.Server.Port), mainRouter); err != nil {
		s.Logger.Error("Failed to start server", "error", err)
		return
	}

}
