package db

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

func Setup() (*bun.DB, error) {
	ctx := context.Background()

	// Open database connection (persistent file)
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:data/db.sqlite?cache=shared")
	if err != nil {
		return nil, err
	}

	// Create Bun database instance
	db := bun.NewDB(sqldb, sqlitedialect.New())

	// Add query debugging (optional)
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	// Create tables
	_, err = db.NewCreateTable().Model((*User)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		_ = db.Close()
		return nil, err
	}

	_, err = db.NewCreateTable().Model((*Note)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}
