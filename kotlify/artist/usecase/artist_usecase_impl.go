package usecase

import (
	"kotlify/artist"
	"kotlify/model"
)

// ArtistUsecaseImpl is usecase
type ArtistUsecaseImpl struct {
	artistRepo artist.ArtistRepo
}

// CreateArtistUsecase is usecase implement
func CreateArtistUsecase(artistRepo artist.ArtistRepo) artist.ArtistUsecase {
	return &ArtistUsecaseImpl{artistRepo}
}

// GetAllArtist is aaa
func (m *ArtistUsecaseImpl) GetAllArtist() (*[]model.Artist, error) {
	return m.artistRepo.GetAllArtist()
}

// GetArtistByID is aaa
func (m *ArtistUsecaseImpl) GetArtistByID(id int) (*model.Artist, error) {
	return m.artistRepo.GetArtistByID(id)
}
