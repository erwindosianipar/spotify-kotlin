package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"kotlify/artist"
	"kotlify/model"
)

var _ artist.ArtistRepo = &ArtistRepoMysqlImpl{}

// ArtistRepoMysqlImpl is struct
type ArtistRepoMysqlImpl struct {
	db *sql.DB
}

// CreateArtistRepoMysqlImpl is repo
func CreateArtistRepoMysqlImpl(db *sql.DB) artist.ArtistRepo {
	return &ArtistRepoMysqlImpl{db}
}

// GetAllArtist is repo
func (m *ArtistRepoMysqlImpl) GetAllArtist() (*[]model.Artist, error) {
	query := "SELECT id, name, debut, category, image FROM artist"
	rows, err := m.db.Query(query)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[ArtistRepoMysqlImpl.GetAllArtist] Error when query get all artist: %w", err)
	}

	defer rows.Close()

	var sliceArtist []model.Artist

	for rows.Next() {
		var each = model.Artist{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Debut, &each.Category, &each.Image)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, fmt.Errorf("[ArtistRepoMysqlImpl.GetAllArtist] Error when scanning rows artist: %w", err)
		}

		sliceArtist = append(sliceArtist, each)
	}

	return &sliceArtist, nil
}

// GetArtistByID is aaa
func (m *ArtistRepoMysqlImpl) GetArtistByID(id int) (*model.Artist, error) {
	artist := model.Artist{}

	query := "SELECT id, name, debut, category, image FROM artist WHERE id = ?"
	err := m.db.QueryRow(query, id).Scan(&artist.ID, &artist.Name, &artist.Debut, &artist.Category, &artist.Image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[ArtistRepoMysqlImpl.GetAllArtist] Error when query get artist by id: %w", err)
	}

	return &artist, nil
}
