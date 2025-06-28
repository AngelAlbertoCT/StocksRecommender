package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	url := os.Getenv("DB_URL")

	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatalf("❌ Error parsing DB config: %v", err)
	}

	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("❌ Failed to connect to CockroachDB: %v", err)
	}

	fmt.Println("✅ Connected to CockroachDB")

	_, err = DB.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS stocks (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		ticker STRING,
		company STRING,
		brokerage STRING,
		action STRING,
		rating_from STRING,
		rating_to STRING,
		target_from STRING,
		target_to STRING,
		time TIMESTAMPTZ,
		UNIQUE (ticker, company, brokerage, time)
	)
`)

	if err != nil {
		log.Fatalf("❌ Error creating table: %v", err)
	}

	fmt.Println("✅ Table checked/created")
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
