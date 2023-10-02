package payloads

import "github.com/Hann0/music-api-clone/internal/models"

type AlbumDTO struct {
	ID           int         `json:"id"`
	Title        string      `json:"title"`
	UPC          string      `json:"upc"`
	Cover        string      `json:"cover"`
	Type         string      `json:"type"`
	Artist       *ArtistDTO  `json:"artist"`
	NumberTracks int         `json:"nb_tracks"`
	Tracks       []*TrackDTO `json:"tracks"`
}

func NewAlbumDTO(album *models.Album) *AlbumDTO {
	return &AlbumDTO{
		ID:    album.ID,
		Title: album.Title,
		Cover: album.Cover,
		UPC:   album.UPC,
		Type:  "album",
	}
}

func (a *AlbumDTO) SetArtist(artist *ArtistDTO) {
	a.Artist = artist
}

func (a *AlbumDTO) AddTrack(track *TrackDTO) {
	a.Tracks = append(a.Tracks, track)
}

func (a *AlbumDTO) SetNumberTracks(number int) {
	a.NumberTracks = number
}
