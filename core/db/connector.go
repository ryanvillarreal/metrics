package db

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	_ "modernc.org/sqlite"

	"github.com/ryanvillarreal/metrics/core/db/scans_db"
	"github.com/tomsteele/go-nmap"
)

//go:embed schema.sql
var ddl string

// priv run() for starting the db in the bg using ctx
func run() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite", "core/web/data/collection.db")
	if err != nil {
		return err
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	// scans_db is our exported package using sqlc generate
	queries := scans_db.New(db)

	// tests
	tmp, err := queries.Count(ctx)
	if err != nil {
		return err
	}
	log.Println(tmp)

	return err
}

// Start used to init the db into mem as early as possible
func Start() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// all encomapassing
func StoreMetrics(nmapRun *nmap.NmapRun) error {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	return nil
}
