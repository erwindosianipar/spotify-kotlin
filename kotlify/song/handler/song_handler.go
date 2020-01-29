package handler

import (
	"fmt"
	"kotlify/library"
	"kotlify/song"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// SongHandler for handling request to song
type SongHandler struct {
	SongUsecase song.SongUsecase
}

// CreateSongHandler for handling request to all song route
func CreateSongHandler(r *mux.Router, songUsecase song.SongUsecase) {
	songHandler := SongHandler{songUsecase}

	r.HandleFunc("/song", songHandler.getAllSong).Methods(http.MethodGet)
	r.HandleFunc("/song/artist/{id}", songHandler.getSongByArtistID).Methods(http.MethodGet)
	r.HandleFunc("/song/genre/{id}", songHandler.getSongByGenreID).Methods(http.MethodGet)
}

func (m *SongHandler) getAllSong(resp http.ResponseWriter, req *http.Request) {
	song, err := m.SongUsecase.GetAllSong()

	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong.")
		fmt.Println("[SongHandler.getAllSong] error when reading request body: ", err.Error())
		return
	}

	library.HandleSuccess(resp, song)
}

func (m *SongHandler) getSongByArtistID(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strID := muxVar["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		library.HandleError(resp, "Oops, Parameter ID Must be a Number.")
		return
	}

	song, err := m.SongUsecase.GetSongByArtistID(id)

	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong.")
		fmt.Println("[SongHandler.getSongByArtistID] error when reading request body: ", err.Error())
		return
	}

	library.HandleSuccess(resp, song)
}

func (m *SongHandler) getSongByGenreID(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strID := muxVar["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		library.HandleError(resp, "Oops, Parameter ID Must be a Number.")
		return
	}

	song, err := m.SongUsecase.GetSongByGenreID(id)

	if err != nil {
		library.HandleError(resp, "Oops, Something Went Wrong.")
		fmt.Println("[SongHandler.getSongByGenreID] error when reading request body: ", err.Error())
		return
	}

	library.HandleSuccess(resp, song)
}
