package server

import (
	"github.com/kaluginivann/Dark_Kitchen/config"
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

func (s *Server) Run() {
	Database, err := db.New(s.Config)
	if err != nil {
		s.Logger.Error("Error from start db", "error", err)
	}
	defer Database.Db.Close()
}
