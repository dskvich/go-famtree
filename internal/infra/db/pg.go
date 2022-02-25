package db

import (
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"time"

	"github.com/joffrua/go-famtree/config"

	log "github.com/sirupsen/logrus"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Pg struct {
	cfg   *config.Config
	url   string
	sqldb *sql.DB
	db    *bun.DB
}

const dbName = "app"

func NewPg(cfg *config.Config) *Pg {
	pg := new(Pg)
	pg.cfg = cfg

	if err := pg.configure(); err != nil {
		log.Panicf("Configuration failed: %+v", err)
	}

	if err := pg.migrate(); err != nil {
		log.Panicf("Migration failed: %+v", err)
	}

	return pg
}

func (pg *Pg) configure() error {
	pg.url = pg.cfg.PG.URL
	if pg.url == "" {
		pg.url = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbName, dbName, pg.cfg.PG.Host, dbName)
	}
	log.Infof("pg connection string: %s", pg.url)

	pg.sqldb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(pg.url)))
	pg.sqldb.SetMaxOpenConns(25)
	pg.sqldb.SetMaxIdleConns(25)
	pg.sqldb.SetConnMaxLifetime(5 * time.Minute)

	pg.db = bun.NewDB(pg.sqldb, pgdialect.New())

	if pg.cfg.PG.ShowSQL {
		pg.db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return nil
}

func (pg *Pg) migrate() error {
	folder := ""
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && d.Name() == "migrations" {
			path = strings.ReplaceAll(path, "\\", "/")
			folder = path
		}
		return nil
	})
	if err != nil {
		return err
	}

	log.Infof("migrations folder is: %+v", folder)

	m, err := migrate.New(fmt.Sprintf("file://%s", folder), pg.url)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func (pg *Pg) GetConnection() *bun.DB {
	return pg.db
}

func (pg *Pg) Disconnect() {
	if pg.db == nil {
		log.Error("Unable to disconnect: bun.DB is not defined")
	}

	if pg.sqldb == nil {
		log.Error("Unable to disconnect: sql.DB is not defined")
	}

	if err := pg.db.Close(); err != nil {
		log.Errorf("Failed to close bun.DB: %+v", err)
	}

	if err := pg.sqldb.Close(); err != nil {
		log.Errorf("Failed to close sql.DB: %+v", err)
	}
}
