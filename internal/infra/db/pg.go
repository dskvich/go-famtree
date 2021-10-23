package db

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Pg struct {
	user    string
	pass    string
	host    string
	port    string
	name    string
	showSQL bool
	conn    *pg.DB
	ctx     context.Context
}

func NewPg(ctx context.Context) *Pg {
	db := new(Pg)
	db.ctx = ctx

	if err := db.configure(); err != nil {
		panic("Configuration failed: " + err.Error())
	}

	if err := db.migrate(); err != nil {
		panic("Migration failed: " + err.Error())
	}

	return db
}

func (db *Pg) configure() error {
	db.user = getEnv("DB_USER", "app")
	db.pass = getEnv("DB_PASS", "app")
	db.host = getEnv("DB_HOST", "localhost")
	db.port = getEnv("DB_PORT", "65432")
	db.name = getEnv("DB_NAME", "app")
	db.showSQL, _ = strconv.ParseBool(getEnv("DB_SHOW_SQL", "f"))

	db.conn = db.connect()

	if db.showSQL {
		db.conn.AddQueryHook(pgLogger{})
	}

	return nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (db *Pg) connect() *pg.DB {
	options := db.pgOptions()
	return pg.Connect(options)
}

func (db *Pg) pgOptions() *pg.Options {
	return &pg.Options{
		User:     db.user,
		Password: db.pass,
		Database: db.name,
		Addr:     fmt.Sprintf("%s:%s", db.host, db.port),

		MaxRetries:      1,
		MinRetryBackoff: -1,

		DialTimeout:  30 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,

		PoolSize:           10,
		MaxConnAge:         10 * time.Second,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        10 * time.Second,
		IdleCheckFrequency: 100 * time.Millisecond,
	}
}

func (db *Pg) migrate() error {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		url.QueryEscape(db.user),
		url.QueryEscape(db.pass),
		fmt.Sprintf("%s:%s", db.host, db.port),
		url.QueryEscape(db.name))

	m, err := migrate.New("file://internal/infra/db/migrations", connStr)
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
		fmt.Println("Unable to disconnect: connection is not defined")
	}
	if err := db.conn.Close(); err != nil {
		fmt.Println("Failed to close connection:", err.Error())
	}
}
