package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DB     Mongo
	Server Server
}

type Mongo struct {
	URI      string `default:"mongodb://localhost:27017"`
	Username string `default:"admin"`
	Password string `default:"g0langn1nja"`
	Database string `default:"crud_logs"`
}

type Server struct {
	Port int `default:"9000"`
}

func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("db", &cfg.DB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("server", &cfg.Server); err != nil {
		return nil, err
	}

	return cfg, nil
}
