package main

import (
	"log"

	"github.com/dcwk/gophermart/internal/config"
)

func main() {
	conf, err := config.NewServerConf()
	if err != nil {
		log.Fatal(err)
	}

	handlers.Run(conf)
}
