package handler

import (
	"fmt"
	"kotlify/artist"
	"kotlify/library"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ArtistHandler for handling request to artist
type ArtistHandler struct {
	ArtistUsecase artist.ArtistUsecase
}

// CreateArtistHandler for handling request to all artist route
func CreateArtistHandler(r *mux.Router, artistUsecase artist.ArtistUsecase) {
	artistHandler := ArtistHandler{artistUsecase}

	r.HandleFunc("/artist", artistHandler.getAllArtist).Methods(http.MethodGet)
	r.HandleFunc("/artist/{id}", artistHandler.getArtistByID).Methods(http.MethodGet)
}

func (m *ArtistHandler) getAllArtist(resp http.ResponseWriter, req *http.Request) {
	artist, err := m.ArtistUsecase.GetAllArtist()

	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong.")
		fmt.Println("[ArtistHandler.getAllArtist] error when reading request body: ", err.Error())
		return
	}

	library.HandleSuccess(resp, artist)
}

func (m *ArtistHandler) getArtistByID(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strID := muxVar["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		library.HandleError(resp, "Oops, Parameter ID Must be a Number.")
		return
	}

	artist, err := m.ArtistUsecase.GetArtistByID(id)
	if err != nil {
		library.HandleError(resp, err.Error())
		return
	}

	library.HandleSuccess(resp, artist)
}
