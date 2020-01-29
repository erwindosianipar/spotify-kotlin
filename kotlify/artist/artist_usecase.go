package artist

import "kotlify/model"

// ArtistUsecase is a interface
type ArtistUsecase interface {
	GetAllArtist() (*[]model.Artist, error)
	GetArtistByID(id int) (*model.Artist, error)
}
