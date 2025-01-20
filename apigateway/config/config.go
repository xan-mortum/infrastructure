package config

type Config struct {
	Server    *Server
	Auth      *Auth
	Resources *Resources
}

func GetConfig() *Config {
	return &Config{
		Server: &Server{
			Port: ":8080",
		},
		Auth: &Auth{
			URL: "http://authapi:8081",
		},
		Resources: &Resources{
			URL: "http://resourceapi:8082",
		},
	}
}
