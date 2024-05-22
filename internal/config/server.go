package config

import (
	"flag"

	"github.com/caarlos0/env"
)

type ServerConf struct {
	RunAddress           string `env:"RUN_ADDRESS"`
	DatabaseUri          string `env:"DATABASE_URI"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func NewServerConf() (*ServerConf, error) {
	conf := &ServerConf{}

	flag.StringVar(&conf.RunAddress, "a", "localhost:8081", "server address")
	flag.StringVar(&conf.DatabaseUri, "d", "postgres://postgres:123456@localhost:5432/gophermart", "database dsn")
	flag.StringVar(&conf.AccrualSystemAddress, "r", "localhost:8080", "accrual system address")

	err := env.Parse(conf)
	if err != nil {
		return nil, err
	}

	return conf, err
}
