package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB *DataBaseConfig
}

type DataBaseConfig struct {
	Host     string
	User     string
	Password string
	DB       string
	Port     int
}

func LoadConfig() *Config {
	_ = godotenv.Load("infra/.env")
	config := &Config{
		DB: &DataBaseConfig{
			Host:     GetEnv("POSTGRES_HOST", "postgres"),
			User:     GetEnv("POSTGRES_USER", "postgres"),
			Password: GetEnv("POSTGRES_PASSWORD", "my_pass"),
			DB:       GetEnv("POSTGRES_DB", "my_db"),
			Port:     GetIntEnv("POSTGRES_PORT", 5432),
		},
	}
	return config
}

func GetEnv(value, defVal string) string {
	envVal := os.Getenv(value)
	if envVal == "" {
		return defVal
	}
	return envVal
}

func GetIntEnv(value string, defVal int) int {
	if envVal := os.Getenv(value); envVal != "" {
		if valInt, err := strconv.Atoi(value); err == nil {
			return valInt
		}
	}
	return defVal
}
