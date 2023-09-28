package config

import (
	"os"
	"strconv"
)

type ServerConfig struct {
	Host string
	Port int
}

func NewServerConfig() (*ServerConfig, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	return &ServerConfig{
		Host: os.Getenv("HOST"),
		Port: port,
	}, nil
}
