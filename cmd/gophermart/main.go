package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/dcwk/gophermart/internal/application"
	"github.com/dcwk/gophermart/internal/config"
	"github.com/dcwk/gophermart/migrations"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	conf, err := config.NewServerConf()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("pgx", conf.DatabaseDSN)
	if err != nil {
		panic(err)
	}
	err = migrations.RunMigrations(context.Background(), db)
	if err != nil {
		panic(err)
	}
	db.Close()

	application.Run(conf)
}
