package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	artistHandler "kotlify/artist/handler"
	artistRepo "kotlify/artist/repo"
	artistUsecase "kotlify/artist/usecase"

	genreHandler "kotlify/genre/handler"
	genreRepo "kotlify/genre/repo"
	genreUsecase "kotlify/genre/usecase"

	songHandler "kotlify/song/handler"
	songRepo "kotlify/song/repo"
	songUsecase "kotlify/song/usecase"
)

func main() {
	port := "8080"

	connect := "root:erwindo123@tcp(127.0.0.1:3306)/kotlify"

	db, err := sql.Open("mysql", connect)
	if err != nil {
		log.Fatal("Error When Connect to DB " + connect + " : " + err.Error())
	}

	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

	artistRepo := artistRepo.CreateArtistRepoMysqlImpl(db)
	artistUsecase := artistUsecase.CreateArtistUsecase(artistRepo)
	artistHandler.CreateArtistHandler(router, artistUsecase)

	genreRepo := genreRepo.CreateGenreRepoMysqlImpl(db)
	genreUsecase := genreUsecase.CreateGenreUsecase(genreRepo)
	genreHandler.CreateGenreHandler(router, genreUsecase)

	songRepo := songRepo.CreateSongRepoMysqlImpl(db)
	songUsecase := songUsecase.CreateSongUsecase(songRepo)
	songHandler.CreateSongHandler(router, songUsecase)

	fmt.Println("Starting Web Server at port : " + port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
