package config

type Config struct {
	JWT    *JWT
	Server *Server
}

func GetConfig() *Config {
	return &Config{
		JWT: &JWT{
			SecretKey: "secret-key",
		},
		Server: &Server{
			Port: ":8081",
		},
	}
}
