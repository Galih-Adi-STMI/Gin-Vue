
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

// DB wraps the sqlx.DB connection
type DB struct {
	DB *sqlx.DB
}

// initDB initializes the database connection
func initDB() (*DB, error) {
	log.Println("Inisialisasi Database Postgres...")

	pgHost := os.Getenv("DB_HOST")
	pgPort := os.Getenv("DB_PORT")
	pgUser := os.Getenv("DB_USER")
	pgPassword := os.Getenv("DB_PASSWORD")
	pgDB := os.Getenv("DB_NAME")
	pgSSL := os.Getenv("DB_SSLMODE")

	pgConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pgHost, pgPort, pgUser, pgPassword, pgDB, pgSSL)
	db, err := sqlx.Open("postgres", pgConnString)
	if err != nil {
		return nil, fmt.Errorf("gagal menghubungkan ke database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("gagal ping database: %w", err)
	}
	return &DB{DB: db}, nil
}

// Close closes the database connection
func (d *DB) Close() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("gagal menutup koneksi database postgres: %w", err)
	}
	return nil
}
