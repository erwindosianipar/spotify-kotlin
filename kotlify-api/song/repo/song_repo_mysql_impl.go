package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"kotlify/model"
	"kotlify/song"
)

var _ song.SongRepo = &SongRepoMysqlImpl{}

// SongRepoMysqlImpl is struct
type SongRepoMysqlImpl struct {
	db *sql.DB
}

// CreateSongRepoMysqlImpl is repo
func CreateSongRepoMysqlImpl(db *sql.DB) song.SongRepo {
	return &SongRepoMysqlImpl{db}
}

// GetAllSong is a repo
func (m *SongRepoMysqlImpl) GetAllSong() (*[]model.Song, error) {
	query := "SELECT id, genre_id, artist_id, name FROM song"
	rows, err := m.db.Query(query)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[SongRepoMysqlImpl.GetAllSong] Error when query get all song: %w", err)
	}

	defer rows.Close()

	var sliceSong []model.Song

	for rows.Next() {
		var each = model.Song{}
		var err = rows.Scan(&each.ID, &each.GenreID, &each.ArtistID, &each.Name)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, fmt.Errorf("[SongRepoMysqlImpl.GetAllSong] Error when scanning rows song: %w", err)
		}

		sliceSong = append(sliceSong, each)
	}

	return &sliceSong, nil
}

// GetSongByArtistID is repo impl
func (m *SongRepoMysqlImpl) GetSongByArtistID(id int) (*[]model.Song, error) {
	query := "SELECT id, genre_id, artist_id, name FROM song WHERE artist_id = ?"
	rows, err := m.db.Query(query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[SongRepoMysqlImpl.GetSongByArtistID] Error when query get all song: %w", err)
	}

	defer rows.Close()

	var sliceSong []model.Song

	for rows.Next() {
		var each = model.Song{}
		var err = rows.Scan(&each.ID, &each.GenreID, &each.ArtistID, &each.Name)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, fmt.Errorf("[SongRepoMysqlImpl.GetSongByArtistID] Error when scanning rows song: %w", err)
		}

		sliceSong = append(sliceSong, each)
	}

	return &sliceSong, nil
}

// GetSongByGenreID is repo impl
func (m *SongRepoMysqlImpl) GetSongByGenreID(id int) (*[]model.Song, error) {
	query := "SELECT id, genre_id, artist_id, name FROM song WHERE genre_id = ?"
	rows, err := m.db.Query(query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[SongRepoMysqlImpl.GetSongByGenreID] Error when query get all song: %w", err)
	}

	defer rows.Close()

	var sliceSong []model.Song

	for rows.Next() {
		var each = model.Song{}
		var err = rows.Scan(&each.ID, &each.GenreID, &each.ArtistID, &each.Name)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, fmt.Errorf("[SongRepoMysqlImpl.GetSongByGenreID] Error when scanning rows song: %w", err)
		}

		sliceSong = append(sliceSong, each)
	}

	return &sliceSong, nil
}
