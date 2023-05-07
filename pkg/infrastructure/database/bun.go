package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewBunDB(sqldb *sql.DB, showSQL bool) *bun.DB {
	db := bun.NewDB(sqldb, pgdialect.New())
	if showSQL {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}
	return db
}
