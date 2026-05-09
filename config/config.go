package config

import (
	"os"
	"strconv"
)

type Config struct {
	Environment string
	Server      ServerConfig
	Postgres    Postgres
}

func NewConfig() (*Config, error) {
	var config Config

	config.Environment = getEnv("ENVIRONMENT", "DEV")

	config.Server.loadFromEnv()
	config.Postgres.loadFromEnv()

	return &config, nil
}

func getEnv(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		val = def
	}
	return val
}

func getEnvInt(key string, def int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return def
}
