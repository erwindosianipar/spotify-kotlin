package genre

import "kotlify/model"

// GenreRepo is a interface
type GenreRepo interface {
	GetAllGenre() (*[]model.Genre, error)
}
