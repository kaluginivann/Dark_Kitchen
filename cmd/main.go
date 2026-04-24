package main

import (
	"github.com/kaluginivann/Dark_Kitchen/config"
	"github.com/kaluginivann/Dark_Kitchen/internal/server"
	"github.com/kaluginivann/Dark_Kitchen/pkg/logger"
)

func main() {
	config := config.LoadConfig()
	log := logger.New()

	server := server.New(config, log)

	server.Run()
}
