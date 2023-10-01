package main

import (
	"log"
	"net/http"

	"github.com/Hann0/music-api-clone/internal/adapters/repositories"
	"github.com/Hann0/music-api-clone/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	server := handler.NewServer(
		repositories.NewInMemoryAlbumRepository(),
		repositories.NewInMemoryArtistRepository(),
		repositories.NewInMemoryTrackRepository(),
	)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// auth
	r.Post("/login", server.Login)
	r.Post("/register", server.Register)
	r.Post("/logout", server.Logout)

	// api
	r.Get("/album/{id}", server.GetAlbum)
	r.Get("/albums", server.GetAlbums)
	r.Get("/artist/{id}", server.GetArtist)
	r.Get("/track/{id}", server.GetTrack)

	log.Fatal(http.ListenAndServe(":8080", r))
}
