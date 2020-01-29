package usecase

import (
	"kotlify/genre"
	"kotlify/model"
)

// GenreUsecaseImpl is usecase
type GenreUsecaseImpl struct {
	genreRepo genre.GenreRepo
}

// CreateGenreUsecase is usecase implement
func CreateGenreUsecase(genreRepo genre.GenreRepo) genre.GenreUsecase {
	return &GenreUsecaseImpl{genreRepo}
}

// GetAllGenre is aaa
func (m *GenreUsecaseImpl) GetAllGenre() (*[]model.Genre, error) {
	return m.genreRepo.GetAllGenre()
}
