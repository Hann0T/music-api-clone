package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Hann0/music-api-clone/internal/ports"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	AlbumRepository  ports.AlbumRepository
	ArtistRepository ports.ArtistRepository
	TrackRepository  ports.TrackRepository
}

func NewServer(albumRepository ports.AlbumRepository, artistRepository ports.ArtistRepository, trackRepository ports.TrackRepository) *Server {
	return &Server{
		AlbumRepository:  albumRepository,
		ArtistRepository: artistRepository,
		TrackRepository:  trackRepository,
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
		fmt.Println("Error trying to marshal the model, error: ", err)
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
		fmt.Println("Error trying to marshal the model, error: ", err)
		return
	}

	w.Write(resp)
}

func (s *Server) GetArtist(w http.ResponseWriter, r *http.Request) {
	artistID := chi.URLParam(r, "id")

	if artistID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(artistID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	album, err := s.ArtistRepository.ReadArtist(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(album)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error trying to marshal the model, error: ", err)
		return
	}

	w.Write(resp)
}

func (s *Server) GetTrack(w http.ResponseWriter, r *http.Request) {
	trackID := chi.URLParam(r, "id")

	if len(trackID) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(trackID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	track, err := s.TrackRepository.ReadTrack(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("Error trying to get the model, error: ", err)
		return
	}

	data, err := json.Marshal(track)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error trying to marshal the model, error: ", err)
		return
	}

	w.Write(data)
}
