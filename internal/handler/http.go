package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hann0/music-api-clone/internal/ports"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	AlbumRepository  ports.AlbumRepository
	ArtistRepository ports.ArtistRepository
}

func NewServer(albumRepository ports.AlbumRepository, artistRepository ports.ArtistRepository) *Server {
	return &Server{
		AlbumRepository:  albumRepository,
		ArtistRepository: artistRepository,
	}
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) GetAlbum(w http.ResponseWriter, r *http.Request) {
	albumID := chi.URLParam(r, "id")

	if albumID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(albumID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	album, err := s.AlbumRepository.ReadAlbum(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(album)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func (s *Server) GetAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := s.AlbumRepository.ReadAlbums()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(albums) <= 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(albums)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func (s *Server) GetArtist(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) GetTrack(w http.ResponseWriter, r *http.Request) {

}
