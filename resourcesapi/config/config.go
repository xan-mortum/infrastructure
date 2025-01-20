package config

type Config struct {
	Server *Server
}

func GetConfig() *Config {
	return &Config{
		Server: &Server{
			Port: ":8082",
		},
	}
}
