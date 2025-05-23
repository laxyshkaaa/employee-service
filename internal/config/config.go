package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Env    string `env:"ENV" envDefault:"local"`
	Server struct {
		Port string `env:"PORT"`
		Host string `env:"HOST"`
	}
	Postgres struct {
		DSN string `env:"POSTGRES_DSN"`
	}
}

func MustLoad() *Config {
	var config Config
	if err := cleanenv.ReadEnv(&config); err != nil {
		panic("failed to load env vars: " + err.Error())
	}
	return &config
}
