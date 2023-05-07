package database

import (
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/sushkevichd/go-famtree/config"

	"github.com/rs/zerolog/log"

	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Pg struct {
	cfg   *config.PG
	url   string
	sqldb *sql.DB
}

const dbName = "app"

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
		return nil, fmt.Errorf("running migrations: %+v", err)
	}

	return sqldb, nil
}

func runMigrations(sqldb *sql.DB) error {
	migrationsPath, err := findMigrationsPath(".")
	if err != nil {
		return err
	}
	log.Info().Msgf("migrations path: %+v", migrationsPath)

	driver, err := postgres.WithInstance(sqldb, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsPath),
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func findMigrationsPath(startPath string) (string, error) {
	var migrationsPath string
	err := filepath.WalkDir(startPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && d.Name() == "migrations" {
			migrationsPath = strings.ReplaceAll(path, "\\", "/")
			return fs.SkipDir
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if migrationsPath == "" {
		return "", errors.New("migrations path not found")
	}

	return migrationsPath, nil
}

func ClosePostgres(sqldb *sql.DB) {
	if err := sqldb.Close(); err != nil {
		log.Error().Err(err).Msg("closing PostgreSQL connection")
	}
}
