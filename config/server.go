package config

type ServerConfig struct {
	Port    int
	BaseApi string
}

func (s *ServerConfig) loadFromEnv() {
	s.Port = getEnvInt("SERVER_PORT", 8080)
	s.BaseApi = getEnv("BASE_API", "api/v1")
}
