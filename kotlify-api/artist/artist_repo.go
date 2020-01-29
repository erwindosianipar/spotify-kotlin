package artist

import "kotlify/model"

// ArtistRepo is a interface
type ArtistRepo interface {
	GetAllArtist() (*[]model.Artist, error)
	GetArtistByID(id int) (*model.Artist, error)
}
