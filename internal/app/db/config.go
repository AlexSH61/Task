package db

type Config struct {
	DataBaseUrl string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
