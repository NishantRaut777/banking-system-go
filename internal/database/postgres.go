package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect(databaseURL string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal("Failed to create DB pool:", err)
	}

	// Verify connection
	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Failed to ping DB:", err)
	}

	DB = pool
	log.Println("PostgreSQL connected successfully")
}

func Close() {
	if DB != nil {
		DB.Close()
		log.Println("PostgreSQL connection closed")
	}
}
