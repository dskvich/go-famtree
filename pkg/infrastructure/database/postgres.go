package database

import (
	"database/sql"
	"embed"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/uptrace/bun/driver/pgdriver"
)

const dbName = "app"

//go:embed migrations/*.sql
var migrationsFS embed.FS

func NewPostgres(url, host string) (*sql.DB, error) {
	if url == "" {
		url = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbName, dbName, host, dbName)
	}
	log.Info().Msgf("PostgreSQL connection string: %s", url)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(url)))
	sqldb.SetMaxOpenConns(25)
	sqldb.SetMaxIdleConns(25)
	sqldb.SetConnMaxLifetime(5 * time.Minute)

	if err := runMigrations(sqldb); err != nil {
		return nil, fmt.Errorf("running migrationFS: %+v", err)
	}

	return sqldb, nil
}

func runMigrations(sqldb *sql.DB) error {
	source := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFS,
		Root:       "migrations",
	}
	if _, err := migrate.Exec(sqldb, "postgres", source, migrate.Up); err != nil {
		return err
	}
	return nil
}

func ClosePostgres(sqldb *sql.DB) {
	if err := sqldb.Close(); err != nil {
		log.Error().Err(err).Msg("closing PostgreSQL connection")
	}
}
