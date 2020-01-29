package song

import "kotlify/model"

// SongRepo is a interface
type SongRepo interface {
	GetAllSong() (*[]model.Song, error)
	GetSongByArtistID(id int) (*[]model.Song, error)
	GetSongByGenreID(id int) (*[]model.Song, error)
}
