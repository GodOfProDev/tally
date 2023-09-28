package config

import "os"

type DBConfig struct {
	Uri string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Uri: os.Getenv("DB_URI"),
	}
}
