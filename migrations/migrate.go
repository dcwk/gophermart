package migrations

import (
	"context"
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed sql/*.sql
var embedMigrations embed.FS

func RunMigrations(dsn string) error {
	ctx := context.Background()
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.UpContext(ctx, db, "sql"); err != nil {
		return err
	}

	return nil
}
