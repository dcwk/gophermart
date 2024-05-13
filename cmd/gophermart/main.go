package main

import (
	"log"

	"github.com/dcwk/gophermart/internal/config"
	"github.com/dcwk/gophermart/internal/server"
)

func main() {
	conf, err := config.NewServerConf()
	if err != nil {
		log.Fatal(err)
	}

	server.Run(conf)
}
