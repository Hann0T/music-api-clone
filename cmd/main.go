package main

import (
	"log"
	"net/http"

	"github.com/Hann0/music-api-clone/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// auth
	r.Post("/login", handler.Login)
	r.Post("/register", handler.Login)
	r.Post("/logout", handler.Login)

	// api
	r.Get("/album/{id}", handler.GetAlbum)
	r.Get("/artist/{id}", handler.GetArtist)
	r.Get("/track/{id}", handler.GetTrack)

    log.Fatal(http.ListenAndServe(":8080", r))
}
