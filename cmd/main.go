package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Hann0/music-api-clone/internal/adapters/repositories"
	"github.com/Hann0/music-api-clone/internal/handler"
	"github.com/Hann0/music-api-clone/migrations"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	connectionStr := "user=username dbname=deezer_go password=password sslmode=disable" // get data from env
	db, err := sqlx.Connect("postgres", connectionStr)

	if err != nil {
		log.Fatalln(err)
	}

	err = migrations.Run(connectionStr)

	if err != nil {
		log.Fatalln(err)
	}

	err = migrations.PopulateDB(db)
	if err != nil {
		log.Fatalln(err)
	}

	server := handler.NewServer(
		repositories.NewAlbumPostgresRepository(db),
		repositories.NewArtistPostgresRepository(db),
		repositories.NewTrackPostgresRepository(db),
	)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(handler.JsonHeader)

	// auth
	r.Post("/login", server.Login)
	r.Post("/register", server.Register)
	r.Post("/logout", server.Logout)

	// api
	r.Get("/album/{id}", server.GetAlbum)
	r.Get("/albums", server.GetAlbums)
	r.Get("/artist/{id}", server.GetArtist)
	r.Get("/track/{id}", server.GetTrack)

	fmt.Println("Server running in port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
