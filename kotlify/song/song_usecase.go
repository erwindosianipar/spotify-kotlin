package song

import "kotlify/model"

// SongUsecase is a interface
type SongUsecase interface {
	GetAllSong() (*[]model.Song, error)
	GetSongByArtistID(id int) (*[]model.Song, error)
	GetSongByGenreID(id int) (*[]model.Song, error)
}
