package handler

import (
	"fmt"
	"kotlify/genre"
	"kotlify/library"
	"net/http"

	"github.com/gorilla/mux"
)

// GenreHandler for handling request to genre
type GenreHandler struct {
	GenreUsecase genre.GenreUsecase
}

// CreateGenreHandler for handling request to all genre route
func CreateGenreHandler(r *mux.Router, genreUsecase genre.GenreUsecase) {
	genreHandler := GenreHandler{genreUsecase}

	r.HandleFunc("/genre", genreHandler.getAllGenre).Methods(http.MethodGet)
}

func (m *GenreHandler) getAllGenre(resp http.ResponseWriter, req *http.Request) {
	genre, err := m.GenreUsecase.GetAllGenre()

	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong.")
		fmt.Println("[GenreHandler.getAllGenre] error when reading request body: ", err.Error())
		return
	}

	library.HandleSuccess(resp, genre)
}
