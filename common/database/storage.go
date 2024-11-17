package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sql.DB
}

func Open(connectionString string) (*Storage, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error opening PostgreSQL database: %v", err)
	}
	log.Printf("PostgreSQL Database connected successfully.")

	err = createDatabaseIfNotExist(db)
	if err != nil {
		return nil, fmt.Errorf("error creating PostgreSQL database: %v", err)
	}

	connectionStringWithDB := fmt.Sprintf("%s dbname=watch_tracker_db", connectionString)
	db, err = sql.Open("postgres", connectionStringWithDB)
	if err != nil {
		return nil, fmt.Errorf("error reconnecting to PostgreSQL database: %v", err)
	}

	err = createDatabase(db)
	if err != nil {
		return nil, fmt.Errorf("error creating PostgreSQL tables: %v", err)
	}

	return &Storage{DB: db}, nil
}

func createDatabaseIfNotExist(db *sql.DB) error {
	var dbName string

	err := db.QueryRow("SELECT 1 FROM pg_database WHERE datname = 'watch_tracker_db'").Scan(&dbName)

	if err != nil && err.Error() == "sql: no rows in result set" {
		_, err := db.Exec("CREATE DATABASE watch_tracker_db")
		if err != nil {
			return fmt.Errorf("failed to create database: %v", err)
		}
	}

	return nil
}

func createDatabase(db *sql.DB) error {
	createTables := `
	CREATE TABLE IF NOT EXISTS series (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		link TEXT NOT NULL,
		imdb FLOAT,
		start_year INT,
		end_year INT,
		poster_link TEXT,
		country TEXT,
		number_of_episode INT,
		episode_duration INT
	);
	`

	_, err := db.Exec(createTables)
	if err != nil {
		return fmt.Errorf("query execution error: %v", err)
	}

	return nil
}
