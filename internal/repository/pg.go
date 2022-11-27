package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/joffrua/go-famtree/config"

	log "github.com/sirupsen/logrus"

	"github.com/go-pg/pg/v10"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Pg struct {
	cfg  *config.Config
	url  string
	opts *pg.Options
	conn *pg.DB
}

const dbName = "app"

func NewPg(cfg *config.Config) *Pg {
	db := new(Pg)
	db.cfg = cfg

	if err := db.configure(); err != nil {
		log.Panicf("Configuration failed: %+v", err)
	}

	if err := db.migrate(); err != nil {
		log.Panicf("Migration failed: %+v", err)
	}

	return db
}

func (db *Pg) configure() error {
	db.url = db.cfg.PG.URL
	if db.url == "" {
		host := db.cfg.PG.Host
		if host == "" {
			host = "localhost:65432"
		}
		db.url = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbName, dbName, host, dbName)
	}
	log.Infof("pg connection string: %s", db.url)

	var err error
	db.opts, err = pg.ParseURL(db.url)
	if err != nil {
		return err
	}

	db.opts.MaxRetries = 1
	db.opts.MinRetryBackoff = -1

	db.opts.DialTimeout = 30 * time.Second
	db.opts.ReadTimeout = 10 * time.Second
	db.opts.WriteTimeout = 10 * time.Second

	db.opts.PoolSize = 10
	db.opts.MaxConnAge = 10 * time.Second
	db.opts.PoolTimeout = 30 * time.Second
	db.opts.IdleTimeout = 10 * time.Second
	db.opts.IdleCheckFrequency = 100 * time.Millisecond

	db.conn = pg.Connect(db.opts)

	if db.cfg.PG.ShowSQL {
		db.conn.AddQueryHook(pgLogger{})
	}

	return nil
}

func (db *Pg) migrate() error {
	m, err := migrate.New("file://internal/repository/migrations", db.url)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

type pgLogger struct{}

func (l pgLogger) BeforeQuery(ctx context.Context, _ *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (l pgLogger) AfterQuery(_ context.Context, q *pg.QueryEvent) error {
	b, _ := q.FormattedQuery()
	fmt.Println(string(b))
	return nil
}

func (db *Pg) GetConnection() *pg.DB {
	return db.conn
}

func (db *Pg) IsConnected() (bool, error) {
	if db.conn == nil {
		return false, fmt.Errorf("connection is not defined")
	}

	if _, err := db.conn.Exec("SELECT 1"); err != nil {
		return false, err
	}
	return true, nil
}

func (db *Pg) Disconnect() {
	if db.conn == nil {
		log.Error("Unable to disconnect: connection is not defined")
	}
	if err := db.conn.Close(); err != nil {
		log.Errorf("Failed to close connection: %+v", err)
	}
}
