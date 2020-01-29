package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"kotlify/genre"
	"kotlify/model"
)

var _ genre.GenreRepo = &GenreRepoMysqlImpl{}

// GenreRepoMysqlImpl is struct
type GenreRepoMysqlImpl struct {
	db *sql.DB
}

// CreateGenreRepoMysqlImpl is repo
func CreateGenreRepoMysqlImpl(db *sql.DB) genre.GenreRepo {
	return &GenreRepoMysqlImpl{db}
}

// GetAllGenre is a repo
func (m *GenreRepoMysqlImpl) GetAllGenre() (*[]model.Genre, error) {
	query := "SELECT id, name FROM genre"
	rows, err := m.db.Query(query)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[GenreRepoMysqlImpl.GetAllGenre] Error when query get all genre: %w", err)
	}

	defer rows.Close()

	var sliceGenre []model.Genre

	for rows.Next() {
		var each = model.Genre{}
		var err = rows.Scan(&each.ID, &each.Name)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, fmt.Errorf("[GenreRepoMysqlImpl.GetAllGenre] Error when scanning rows genre: %w", err)
		}

		sliceGenre = append(sliceGenre, each)
	}

	return &sliceGenre, nil
}
