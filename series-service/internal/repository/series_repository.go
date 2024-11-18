package repository

import (
	"database/sql"
	"fmt"

	"github.com/capybara120404/common/database"
	"github.com/capybara120404/common/models"
)

var (
	prepareError = "failed to prepare SQL statement: %w"
)

type SeriesRepository struct {
	storage *database.Storage
}

func NewSeriesRepository(storage *database.Storage) *SeriesRepository {
	return &SeriesRepository{
		storage: storage,
	}
}

func (repository *SeriesRepository) AddSeries(series models.Series) error {
	stmt, err := repository.storage.DB.Prepare(`INSERT INTO series (title, link, imdb, start_year, end_year, poster_link, country, number_of_episode, episode_duration)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	var endYear interface{}
	if series.EndYear != nil {
		endYear = *series.EndYear
	} else {
		endYear = nil
	}

	_, err = stmt.Exec(series.Title, series.Link, series.IMDB, series.StartYear, endYear,
		series.PosterLink, series.Country, series.NumberOfEpisode, series.EpisodeDuration)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}

	return nil
}

func (repository *SeriesRepository) DeleteSeriesByID(id uint) error {
	stmt, err := repository.storage.DB.Prepare("DELETE FROM series WHERE id = $1")
	if err != nil {
		return fmt.Errorf(prepareError, err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("failed to execute delete statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not determine rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no series found with id %d", id)
	}

	return nil
}

func (repository *SeriesRepository) GetAllSeries() ([]models.Series, error) {
	stmt, err := repository.storage.DB.Prepare("SELECT * FROM series")
	if err != nil {
		return nil, fmt.Errorf(prepareError, err)
	}
	defer stmt.Close()

	var series []models.Series
	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("error querying series from the database: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var s models.Series
		err := rows.Scan(&s.ID,
			&s.Title,
			&s.Link,
			&s.IMDB,
			&s.StartYear,
			&s.EndYear,
			&s.PosterLink,
			&s.Country,
			&s.NumberOfEpisode,
			&s.EpisodeDuration)
		if err != nil {
			return nil, fmt.Errorf("error scanning series data: %w", err)
		}
		series = append(series, s)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over series rows: %w", err)
	}

	return series, nil
}

func (repository *SeriesRepository) GetSeriesById(id uint) (*models.Series, error) {
	stmt, err := repository.storage.DB.Prepare("SELECT * FROM series WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf(prepareError, err)
	}
	defer stmt.Close()

	var series models.Series
	err = stmt.QueryRow(id).Scan(&series.ID,
		&series.Title,
		&series.Link,
		&series.IMDB,
		&series.StartYear,
		&series.EndYear,
		&series.PosterLink,
		&series.Country,
		&series.NumberOfEpisode,
		&series.EpisodeDuration)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no series found with id %d", id)
		}
		return nil, fmt.Errorf("failed to execute query and scan result: %w", err)
	}

	return &series, nil
}
