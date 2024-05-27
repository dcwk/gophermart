package main

import (
	"log"

	"github.com/dcwk/gophermart/internal/application"
	"github.com/dcwk/gophermart/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	conf, err := config.NewServerConf()
	if err != nil {
		log.Fatal(err)
	}

	application.Run(conf)
}
