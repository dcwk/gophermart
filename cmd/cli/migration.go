package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/dcwk/gophermart/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

func main() {
	conf, err := config.NewServerConf()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("pgx", conf.DatabaseUri)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pwd = "./../../migrations"
	if err := goose.Up(db, pwd); err != nil {
		panic(fmt.Sprintf("Can't migrations up: %s", err))
	}
}
