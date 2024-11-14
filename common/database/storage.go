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

	err = createDatabase(db)
	if err != nil {
		return nil, fmt.Errorf("error creating PostgreSQL database: %v", err)
	}

	return &Storage{DB: db}, nil
}

func (c *Storage) Close() error {
	if err := c.DB.Close(); err != nil {
		return fmt.Errorf("error closing database: %v", err)
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
		poster TEXT,
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
