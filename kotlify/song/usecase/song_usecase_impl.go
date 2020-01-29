package usecase

import (
	"kotlify/model"
	"kotlify/song"
)

// SongUsecaseImpl is usecase
type SongUsecaseImpl struct {
	songRepo song.SongRepo
}

// CreateSongUsecase is usecase implement
func CreateSongUsecase(songRepo song.SongRepo) song.SongUsecase {
	return &SongUsecaseImpl{songRepo}
}

// GetAllSong is aaa
func (m *SongUsecaseImpl) GetAllSong() (*[]model.Song, error) {
	return m.songRepo.GetAllSong()
}

// GetSongByArtistID is aaa
func (m *SongUsecaseImpl) GetSongByArtistID(id int) (*[]model.Song, error) {
	return m.songRepo.GetSongByArtistID(id)
}

// GetSongByGenreID is aaa
func (m *SongUsecaseImpl) GetSongByGenreID(id int) (*[]model.Song, error) {
	return m.songRepo.GetSongByGenreID(id)
}
