package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hann0/music-api-clone/internal/payloads"
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
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(albumID)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusInternalServerError)
		return
	}

	album, err := s.AlbumRepository.ReadAlbum(id)

	if err != nil {
		http.Error(w, "Not album found", http.StatusNotFound)
		return
	}

	artist, err := s.ArtistRepository.ReadArtist(album.ArtistID)

	if err != nil {
		http.Error(w, "Not artist found", http.StatusNotFound)
		return
	}

	tracks, err := s.TrackRepository.ReadTracksByAlbum(album)

	if err != nil {
		http.Error(w, "Not tracks found", http.StatusNotFound)
		return
	}

	albumDTO := payloads.NewAlbumDTO(album)
	albumDTO.SetArtist(payloads.NewArtistDTO(artist))
	albumDTO.SetNumberTracks(len(tracks))

	for _, track := range tracks {
		albumDTO.AddTrack(payloads.NewTrackDTO(track))
	}

	resp, err := json.Marshal(albumDTO)

	if err != nil {
		http.Error(w, "Error marshalling data", http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func (s *Server) GetAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := s.AlbumRepository.ReadAlbums()

	if err != nil {
		http.Error(w, "Error fetching the albums", http.StatusNotFound)
		return
	}

	if len(albums) <= 0 {
		http.Error(w, "Not albums found", http.StatusNotFound)
		return
	}

	var albumsDTO []*payloads.AlbumDTO

	for _, album := range albums {
		albumsDTO = append(albumsDTO, payloads.NewAlbumDTO(album))
	}

	resp, err := json.Marshal(albumsDTO)

	if err != nil {
		http.Error(w, "Error marshalling data", http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func (s *Server) GetArtist(w http.ResponseWriter, r *http.Request) {
	artistID := chi.URLParam(r, "id")

	if artistID == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(artistID)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusInternalServerError)
		return
	}

	artist, err := s.ArtistRepository.ReadArtist(id)

	if err != nil {
		http.Error(w, "Not artist found", http.StatusNotFound)
		return
	}

	albums, err := s.AlbumRepository.ReadAlbumsByArtist(artist)
	if err != nil {
		http.Error(w, "Not albums found", http.StatusNotFound)
		return
	}

	artistDTO := payloads.NewArtistDTO(artist)
	artistDTO.SetNumberAlbums(len(albums))

	resp, err := json.Marshal(artistDTO)

	if err != nil {
		http.Error(w, "Error marshalling data", http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func (s *Server) GetTrack(w http.ResponseWriter, r *http.Request) {
	trackId := chi.URLParam(r, "id")

	if trackId == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(trackId)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusInternalServerError)
		return
	}

	track, err := s.TrackRepository.ReadTrack(id)

	if err != nil {
		http.Error(w, "No track found", http.StatusNotFound)
		return
	}

	album, err := s.AlbumRepository.ReadAlbum(track.AlbumID)

	if err != nil {
		http.Error(w, "No album found", http.StatusNotFound)
		return
	}

	artist, err := s.ArtistRepository.ReadArtist(album.ArtistID)

	if err != nil {
		http.Error(w, "No artist found", http.StatusNotFound)
		return
	}

	trackDTO := payloads.NewTrackDTO(track)
	trackDTO.SetAlbum(payloads.NewAlbumDTO(album))
	trackDTO.SetArtist(payloads.NewArtistDTO(artist))

	data, err := json.Marshal(trackDTO)

	if err != nil {
		http.Error(w, "Error marshalling data", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
