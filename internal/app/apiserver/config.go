package apiserver

import "github.com/AlexSH61/firstRestAPi/internal/app/db"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Task     *db.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Task:     db.NewConfig(),
	}
}
