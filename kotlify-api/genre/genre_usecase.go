package genre

import "kotlify/model"

// GenreUsecase is a interface
type GenreUsecase interface {
	GetAllGenre() (*[]model.Genre, error)
}
